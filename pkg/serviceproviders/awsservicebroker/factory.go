/**
 * Copyright 2020 Appvia Ltd <info@appvia.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package awsservicebroker

import (
	"bytes"
	"fmt"

	"github.com/appvia/kore/pkg/utils"

	"github.com/appvia/kore/pkg/utils/configuration"

	corev1 "github.com/appvia/kore/pkg/apis/core/v1"
	servicesv1 "github.com/appvia/kore/pkg/apis/services/v1"
	"github.com/appvia/kore/pkg/kore"
	"github.com/appvia/kore/pkg/serviceproviders/openservicebroker"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	osb "github.com/kubernetes-sigs/go-open-service-broker-client/v2"
)

func init() {
	kore.RegisterServiceProviderFactory(ProviderFactory{})
}

const (
	S3BucketTagImported  = "kore.appvia.io/initialized"
	ComponentIAMRole     = "IAM Role"
	ComponentDynamoDB    = "DynamoDB Table"
	ComponentS3Bucket    = "S3 Bucket"
	ComponentHelmRelease = "Helm Release"
)

type ProviderFactory struct{}

func (d ProviderFactory) Type() string {
	return "aws-servicebroker"
}

// JSONSchemas returns all JSON schema versions for the provider's configuration
func (d ProviderFactory) JSONSchemas() map[string]string {
	return map[string]string{
		"https://appvia.io/kore/schemas/serviceprovider/aws-servicebroker/v1.json": providerSchemaV1,
	}
}

func (d ProviderFactory) Create(ctx kore.Context, serviceProvider *servicesv1.ServiceProvider) (_ kore.ServiceProvider, _ error) {
	var config = DefaultProviderConfiguration()

	// Migrate deprecated field values
	if serviceProvider.Spec.Configuration != nil {
		serviceProvider.Spec.Configuration.Raw = bytes.ReplaceAll(serviceProvider.Spec.Configuration.Raw, []byte("aws_access_key_id"), []byte("awsAccessKeyID"))
		serviceProvider.Spec.Configuration.Raw = bytes.ReplaceAll(serviceProvider.Spec.Configuration.Raw, []byte("aws_secret_access_key"), []byte("awsSecretAccessKey"))
	}

	if _, err := configuration.ParseObjectConfiguration(ctx, ctx.Client(), serviceProvider, config); err != nil {
		return nil, fmt.Errorf("failed to process aws-servicebroker configuration: %w", err)
	}

	config.PlatformMapping = map[string]string{
		"*": "AWS",
	}

	config.PlanConfigurationOverrides = map[string][]openservicebroker.PlanConfigurationOverride{
		"*": {
			{
				Name:     "target_account_id",
				Value:    "{{ .cluster.status.providerData.awsAccountID }}",
				Const:    true,
				Required: utils.BoolPtr(true),
			},
			{
				Name:     "target_role_name",
				Value:    config.AWSIAMRoleName,
				Const:    true,
				Required: utils.BoolPtr(true),
			},
			{
				Name:     "VpcId",
				Value:    "{{ .cluster.status.providerData.vpc.vpcID }}",
				Const:    true,
				Required: utils.BoolPtr(true),
			},
		},
	}

	namespaceName := "kore-serviceprovider-" + serviceProvider.Name

	clientSecret, err := getServiceAccountToken(ctx, ctx.Client(), namespaceName, serviceProvider.Name+"-aws-servicebroker-client")
	if err != nil {
		return nil, err
	}

	certsSecret, err := getSecret(ctx, ctx.Client(), namespaceName, serviceProvider.Name+"-aws-servicebroker-cert")
	if err != nil {
		return nil, err
	}

	clientConfiguration := osb.DefaultClientConfiguration()
	osbConfig := openservicebroker.ProviderConfiguration{
		ClientConfiguration:  *clientConfiguration,
		CatalogConfiguration: config.CatalogConfiguration,
	}

	osbConfig.URL = fmt.Sprintf("https://%s-aws-servicebroker.kore-serviceprovider-%s.svc", serviceProvider.Name, serviceProvider.Name)
	osbConfig.AuthConfig = &osb.AuthConfig{
		BearerConfig: &osb.BearerConfig{Token: string(clientSecret.Data["token"])},
	}
	osbConfig.CAData = certsSecret.Data["ca.crt"]

	osbClient, err := osb.NewClient(&osbConfig.ClientConfiguration)
	if err != nil {
		return nil, err
	}

	return openservicebroker.NewProvider(serviceProvider.Name, osbConfig, osbClient), nil
}

func (d ProviderFactory) SetUp(ctx kore.Context, serviceProvider *servicesv1.ServiceProvider) (complete bool, _ error) {
	var config = DefaultProviderConfiguration()

	if _, err := configuration.ParseObjectConfiguration(ctx, ctx.Client(), serviceProvider, config); err != nil {
		return false, fmt.Errorf("failed to process aws-servicebroker configuration: %w", err)
	}

	cfg := aws.NewConfig().
		WithCredentials(credentials.NewStaticCredentials(config.AWSAccessKeyID, config.AWSSecretAccessKey, "")).
		WithRegion(config.Region)

	sess := session.Must(session.NewSession(cfg))

	if err := d.ensureIAMRole(sess, config); err != nil {
		serviceProvider.Status.Components.SetCondition(corev1.Component{
			Name:    ComponentIAMRole,
			Status:  corev1.ErrorStatus,
			Message: "Failed to create IAM role",
			Detail:  err.Error(),
		})
		return false, err
	}

	serviceProvider.Status.Components.SetCondition(corev1.Component{Name: ComponentIAMRole, Status: corev1.SuccessStatus})

	if err := d.ensureDynamoDBTable(sess, config); err != nil {
		serviceProvider.Status.Components.SetCondition(corev1.Component{
			Name:    ComponentDynamoDB,
			Status:  corev1.ErrorStatus,
			Message: "Failed to create or update DynamoDB table",
			Detail:  err.Error(),
		})
		return false, err
	}

	serviceProvider.Status.Components.SetCondition(corev1.Component{Name: ComponentDynamoDB, Status: corev1.SuccessStatus})

	if config.S3BucketName != "awsservicebroker" {
		if err := d.ensureS3Bucket(sess, config); err != nil {
			serviceProvider.Status.Components.SetCondition(corev1.Component{
				Name:    ComponentS3Bucket,
				Status:  corev1.ErrorStatus,
				Message: "Failed to create or update S3 bucket",
				Detail:  err.Error(),
			})
			return false, err
		}
	}

	serviceProvider.Status.Components.SetCondition(corev1.Component{Name: ComponentS3Bucket, Status: corev1.SuccessStatus})

	complete, err := d.ensureHelmRelease(ctx, config, serviceProvider.Name, config.AWSAccessKeyID, config.AWSSecretAccessKey)
	if err != nil {
		serviceProvider.Status.Components.SetCondition(corev1.Component{
			Name:    ComponentHelmRelease,
			Status:  corev1.ErrorStatus,
			Message: "Failed to deploy the aws-servicebroker Helm chart",
			Detail:  err.Error(),
		})
		return false, err
	}

	if !complete {
		serviceProvider.Status.Components.SetCondition(corev1.Component{Name: ComponentHelmRelease, Status: corev1.PendingStatus})
		return false, nil
	}

	serviceProvider.Status.Components.SetCondition(corev1.Component{Name: ComponentHelmRelease, Status: corev1.SuccessStatus})

	return true, nil
}

func (d ProviderFactory) TearDown(ctx kore.Context, serviceProvider *servicesv1.ServiceProvider) (complete bool, _ error) {
	var config = DefaultProviderConfiguration()

	if _, err := configuration.ParseObjectConfiguration(ctx, ctx.Client(), serviceProvider, config); err != nil {
		return false, fmt.Errorf("failed to process aws-servicebroker configuration: %w", err)
	}

	done, err := d.deleteHelmRelease(ctx, serviceProvider.Name)
	if err != nil {
		serviceProvider.Status.Components.SetCondition(corev1.Component{
			Name:    ComponentHelmRelease,
			Status:  corev1.ErrorStatus,
			Message: "Failed to delete Helm deployment",
			Detail:  err.Error(),
		})
		return false, err
	}
	if !done {
		return false, nil
	}

	serviceProvider.Status.Components.SetCondition(corev1.Component{Name: ComponentHelmRelease, Status: corev1.DeletedStatus})

	cfg := aws.NewConfig().
		WithCredentials(credentials.NewStaticCredentials(config.AWSAccessKeyID, config.AWSSecretAccessKey, "")).
		WithRegion(config.Region)

	sess := session.Must(session.NewSession(cfg))

	if err := d.deleteDynamoDBTable(sess, config); err != nil {
		serviceProvider.Status.Components.SetCondition(corev1.Component{
			Name:    ComponentDynamoDB,
			Status:  corev1.ErrorStatus,
			Message: "Failed to delete DynamoDB table",
			Detail:  err.Error(),
		})
		return false, err
	}

	if err := d.deleteIAMRole(sess, config); err != nil {
		serviceProvider.Status.Components.SetCondition(corev1.Component{
			Name:    ComponentIAMRole,
			Status:  corev1.ErrorStatus,
			Message: "Failed to delete IAM role",
			Detail:  err.Error(),
		})
		return false, err
	}

	serviceProvider.Status.Components.SetCondition(corev1.Component{Name: ComponentDynamoDB, Status: corev1.DeletedStatus})

	return true, nil
}

func (d ProviderFactory) DefaultProviders() []servicesv1.ServiceProvider {
	return nil
}
