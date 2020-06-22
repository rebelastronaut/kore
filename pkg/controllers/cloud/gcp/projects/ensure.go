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

package projects

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	configv1 "github.com/appvia/kore/pkg/apis/config/v1"
	corev1 "github.com/appvia/kore/pkg/apis/core/v1"
	gcp "github.com/appvia/kore/pkg/apis/gcp/v1alpha1"
	"github.com/appvia/kore/pkg/utils"
	gcputils "github.com/appvia/kore/pkg/utils/cloud/gcp"
	"github.com/appvia/kore/pkg/utils/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client"

	log "github.com/sirupsen/logrus"
	cloudbilling "google.golang.org/api/cloudbilling/v1"
	cloudresourcemanager "google.golang.org/api/cloudresourcemanager/v1"
	iam "google.golang.org/api/iam/v1"
	"google.golang.org/api/option"
	servicemanagement "google.golang.org/api/servicemanagement/v1"
	serviceusage "google.golang.org/api/serviceusage/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

const (
	// ServiceAccountKeyMax is the max number of service account keys to keep
	ServiceAccountKeyMax = 2
	// ServiceAccountDeadline is the time period we should rotate credentials
	ServiceAccountDeadline = 30 * (24 * time.Hour)
)

// EnsurePermitted is responsible for checking the project has access to the credentials
func (t ctrl) EnsurePermitted(ctx context.Context, project *gcp.Project) error {
	// @step: we check if the gcp organization has been allocated to us
	permitted, err := t.Teams().Team(project.Namespace).Allocations().IsPermitted(ctx, project.Spec.Organization)
	if err != nil {
		return err
	}
	if !permitted {
		project.Status.Conditions.SetCondition(corev1.Component{
			Name:    "provision",
			Message: "GCP Organization has not been allocated to you",
			Status:  corev1.FailureStatus,
		})

		return errors.New("gcp organization has not been allocated to team")
	}

	return nil
}

// EnsureProjectUnclaimed is responsible for making sure the project is unclaimed
func (t ctrl) EnsureProjectUnclaimed(ctx context.Context, project *gcp.Project) error {
	logger := log.WithFields(log.Fields{
		"name":      project.Name,
		"namespace": project.Namespace,
	})

	// @step: check if the project project has already been claimed else where
	projected, err := t.IsProjected(ctx, project)
	if err != nil {
		logger.WithError(err).Error("trying to check if the project is already projected")

		project.Status.Status = corev1.FailureStatus
		project.Status.Conditions.SetCondition(corev1.Component{
			Name:    "provision",
			Message: "Unable to fulfil request, project name has already been projected in the organization",
			Status:  corev1.FailureStatus,
		})

		return errors.New("failed to check if project is already projected")
	}
	if projected {
		logger.Warn("attempting to project gcp project which has already been provisioned")

		project.Status.Status = corev1.FailureStatus
		project.Status.Conditions.SetCondition(corev1.Component{
			Name:    "provision",
			Message: "Project has already been projected by another team in kore",
			Status:  corev1.FailureStatus,
		})

		return errors.New("gcp project name already provisioned")
	}

	return nil
}

// EnsureOrganization is responsible for checking and retrieving the gcp org
func (t ctrl) EnsureOrganization(ctx context.Context, project *gcp.Project) (*gcp.Organization, error) {
	org := &gcp.Organization{}

	key := types.NamespacedName{
		Namespace: project.Spec.Organization.Namespace,
		Name:      project.Spec.Organization.Name,
	}

	if err := t.mgr.GetClient().Get(ctx, key, org); err != nil {
		project.Status.Conditions.SetCondition(corev1.Component{
			Name:    "provision",
			Detail:  err.Error(),
			Message: "Attempting to retrieve the GCP Organization resources from API",
			Status:  corev1.FailureStatus,
		})

		return nil, err
	}

	// @step: check if the admin project exists and if successful
	if org.Status.Status != corev1.SuccessStatus {
		project.Status.Conditions.SetCondition(corev1.Component{
			Name:    "provision",
			Detail:  "resource is in failing state",
			Message: "GCP Admin Project is in a failing state, cannot provision projects",
			Status:  corev1.FailureStatus,
		})

		return nil, errors.New("admin project still provisioning or failed")
	}

	return org, nil
}

// EnsureOrganizationCredentials is responsible for retrieving the credentials
func (t ctrl) EnsureOrganizationCredentials(ctx context.Context, org *gcp.Organization, project *gcp.Project) (*configv1.Secret, error) {
	// @TODO we probably shouldn't rely on the parent name here
	secret := &configv1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      org.Spec.CredentialsRef.Name,
			Namespace: org.Spec.CredentialsRef.Namespace,
		},
	}

	found, err := kubernetes.GetIfExists(ctx, t.mgr.GetClient(), secret)
	if err != nil {
		project.Status.Conditions.SetCondition(corev1.Component{
			Name:    "provision",
			Detail:  err.Error(),
			Message: "Attempting to retrieve the GCP Organization credentials",
			Status:  corev1.FailureStatus,
		})

		return nil, err
	}
	if !found {
		project.Status.Conditions.SetCondition(corev1.Component{
			Name:    "provision",
			Detail:  "credentials not found",
			Message: "GCP Organization credentials either not provisioned or failed",
			Status:  corev1.FailureStatus,
		})

		return nil, errors.New("credentials not found")
	}

	// @step: check the credentials have been verified
	if secret.Status.Verified == nil || !*secret.Status.Verified {
		project.Status.Conditions.SetCondition(corev1.Component{
			Name:    "provision",
			Detail:  "credentials failed verification",
			Message: "GCP Organization credentials have failed verification",
			Status:  corev1.FailureStatus,
		})

		return nil, errors.New("organizational credentials not verified")
	}

	// @step: decode the secret for them
	if err := secret.Decode(); err != nil {
		project.Status.Conditions.SetCondition(corev1.Component{
			Name:    "provision",
			Detail:  err.Error(),
			Message: "GCP Organization credentials service account key invalid",
			Status:  corev1.FailureStatus,
		})

		return nil, errors.New("organizational credentials service account invalid")
	}

	return secret, nil
}

// EnsureProject is responsible for ensuring the project is there
func (t ctrl) EnsureProject(ctx context.Context,
	credentials *configv1.Secret,
	org *gcp.Organization,
	project *gcp.Project) error {

	logger := log.WithFields(log.Fields{
		"name":      project.Name,
		"namespace": project.Namespace,
		"project":   project.Spec.ProjectName,
	})
	stage := "Provision GCP Project"
	success := "GCP Project successfully provisioned"

	// @step: create the client
	client, err := cloudresourcemanager.NewService(ctx, option.WithCredentialsJSON([]byte(credentials.Spec.Data["key"])))
	if err != nil {
		logger.WithError(err).Error("trying to create the cloud resource client")

		project.Status.Conditions.SetCondition(corev1.Component{
			Name:    stage,
			Detail:  err.Error(),
			Message: "Failed to create a projects client, please check credentials",
			Status:  corev1.FailureStatus,
		})

		return err
	}

	// @step: we check if the project exists and if not create it
	obj, found, err := IsProject(ctx, client, project.Spec.ProjectName)
	if err != nil {
		logger.WithError(err).Error("trying to check for gcp project")

		project.Status.Conditions.SetCondition(corev1.Component{
			Name:    stage,
			Detail:  err.Error(),
			Message: "Failed to check for project existence",
			Status:  corev1.FailureStatus,
		})

		return err
	}
	if found {
		project.Status.ProjectID = obj.ProjectId

		// @step: we need to check the lifecycle of the project
		switch obj.LifecycleState {
		case "DELETE_REQUESTED", "DELETE_IN_PROGRESS":
			logger.Warn("gcp project has been deleted, attempting to restore")

			_, err := client.Projects.Undelete(obj.ProjectId, &cloudresourcemanager.UndeleteProjectRequest{}).
				Context(ctx).
				Do()
			if err != nil {
				logger.WithError(err).Error("trying to restore project")

				project.Status.Conditions.SetCondition(corev1.Component{
					Name:    stage,
					Detail:  err.Error(),
					Message: "Failed to restore deleted GCP project",
					Status:  corev1.FailureStatus,
				})

				return err
			}
		}

		logger.Debug("gcp project already exists, checking if it was created by us")

		// @TODO we need something to check in the project to see if we create this project
		builder, found := obj.Labels["builder"]
		if !found || builder != "kore" {
			project.Status.Conditions.SetCondition(corev1.Component{
				Name:    stage,
				Detail:  "project conflict",
				Message: "An existing project exist which wasn't created by kore",
				Status:  corev1.FailureStatus,
			})

			return errors.New("project conflict")
		}

		// @step set the status as successful
		project.Status.Conditions.SetCondition(corev1.Component{
			Name:    stage,
			Message: success,
			Status:  corev1.SuccessStatus,
		})

		return nil
	}

	logger.Info("gcp project does not exist, creating it now")

	// @step: create the project in gcp
	resp, err := client.Projects.Create(&cloudresourcemanager.Project{
		Name: project.Spec.ProjectName,
		// @QUESTION should this be the same as the name?
		ProjectId: project.Spec.ProjectName,
		Labels: map[string]string{
			"builder": "kore",
		},
		Parent: &cloudresourcemanager.ResourceId{
			Id:   org.Spec.ParentID,
			Type: org.Spec.ParentType,
		},
	}).Context(ctx).Do()

	if err != nil {
		logger.WithError(err).Error("trying to create the gcp project")

		project.Status.Conditions.SetCondition(corev1.Component{
			Name:    stage,
			Detail:  err.Error(),
			Message: "Unable to request the project in GCP",
			Status:  corev1.FailureStatus,
		})

		return err
	}

	// @step: wait for the operation to complete or fail
	if err := utils.WaitUntilComplete(ctx, 5*time.Minute, 10*time.Second, func() (bool, error) {
		status, err := client.Operations.Get(resp.Name).Context(ctx).Do()
		if err != nil {
			logger.WithError(err).Error("checking the status of the project operation")

			return false, nil
		}
		if !status.Done {
			return false, nil
		}
		if status.Error != nil {
			return false, errors.New(status.Error.Message)
		}

		return true, nil
	}); err != nil {
		project.Status.Conditions.SetCondition(corev1.Component{
			Name:    stage,
			Detail:  err.Error(),
			Message: "Unable to provision project in GCP",
			Status:  corev1.FailureStatus,
		})

		return err
	}

	// @step: we check if the project exists and if not create it
	obj, found, err = IsProject(ctx, client, project.Spec.ProjectName)
	if err != nil {
		logger.WithError(err).Error("trying to check for gcp project")

		project.Status.Conditions.SetCondition(corev1.Component{
			Name:    stage,
			Detail:  err.Error(),
			Message: "Failed to check for project existence",
			Status:  corev1.FailureStatus,
		})

		return err
	}
	if found {
		project.Status.ProjectID = obj.ProjectId
	}

	if !found {
		project.Status.Conditions.SetCondition(corev1.Component{
			Name:    stage,
			Detail:  "project not found",
			Message: "Failed to provision project",
			Status:  corev1.FailureStatus,
		})
	}

	// @step set the status as successful
	project.Status.Conditions.SetCondition(corev1.Component{
		Name:    stage,
		Message: success,
		Status:  corev1.SuccessStatus,
	})

	return nil
}

// EnsureBilling is responsible for ensuring the billing account
func (t ctrl) EnsureBilling(
	ctx context.Context,
	credentials *configv1.Secret,
	organization *gcp.Organization,
	project *gcp.Project) error {

	logger := log.WithFields(log.Fields{
		"name":       project.Name,
		"namespace":  project.Namespace,
		"project_id": project.Status.ProjectID,
	})
	stage := "Associated Billing Account"

	err := func() error {
		client, err := cloudbilling.NewService(ctx, option.WithCredentialsJSON([]byte(credentials.Spec.Data["key"])))
		if err != nil {
			logger.WithError(err).Error("trying to create cloud resource client")

			return err
		}

		uri := fmt.Sprintf("projects/%s", project.Status.ProjectID)

		resp, err := client.Projects.GetBillingInfo(uri).Context(ctx).Do()
		if err != nil {
			logger.WithError(err).Error("trying to retrieve the billing details for account")

			return err
		}

		current := t.BillingAccountName(resp.BillingAccountName)

		// @if they are the same we can return
		if current == organization.Spec.BillingAccount {
			return nil
		}

		if current == "" {
			logger.Info("billing account not set, attempting to set now")
		}
		if current != organization.Spec.BillingAccount {
			logger.Warn("project billing account differs, trying to reconcile now")
		}

		if _, err := client.Projects.UpdateBillingInfo(uri, &cloudbilling.ProjectBillingInfo{
			BillingAccountName: "billingAccounts/" + organization.Spec.BillingAccount,
			BillingEnabled:     true,
		}).Context(ctx).Do(); err != nil {
			logger.WithError(err).Error("trying to update the project billing details")

			return err
		}

		return err
	}()
	if err != nil {
		project.Status.Conditions.SetCondition(corev1.Component{
			Name:    stage,
			Detail:  err.Error(),
			Message: "Failed to link the billing account to project",
			Status:  corev1.FailureStatus,
		})

		return err
	}

	project.Status.Conditions.SetCondition(corev1.Component{
		Name:    stage,
		Message: "GCP Project has been linked billing account",
		Status:  corev1.SuccessStatus,
	})

	return nil
}

// BillingAccountName extracts the billing name from path
func (t ctrl) BillingAccountName(path string) string {
	items := strings.Split(path, "/")

	return items[len(items)-1]
}

// EnsureAPIs is responsible for ensuing the apis are enabled in the account
func (t ctrl) EnsureAPIs(ctx context.Context, credentials *configv1.Secret, project *gcp.Project) error {
	stage := "GCP Account APIs"

	logger := log.WithFields(log.Fields{
		"name":      project.Name,
		"namespace": project.Namespace,
	})

	opt := option.WithCredentialsJSON([]byte(credentials.Spec.Data["key"]))
	// @step: set the project id
	id := project.Status.ProjectID
	if id == "" {
		id = project.Name
	}

	// @step: retrieve a list of enabled services
	usage, err := serviceusage.NewService(ctx, opt)
	if err != nil {
		logger.WithError(err).Error("trying to create the service usage client")

		return err
	}
	list, err := gcputils.ListEnabledAPIs(ctx, usage, id)
	if err != nil {
		logger.WithError(err).Error("trying to retrieve list of enabled services")

		return err
	}
	logger.WithField(
		"enabled", list,
	).Debug("the following api are already enabled")

	// create a client for enabling the apis
	client, err := servicemanagement.NewService(ctx, opt)
	if err != nil {
		logger.WithError(err).Error("trying to create the service management client")

		return err
	}

	wg := &sync.WaitGroup{}

	for _, name := range t.GetRequiredAPI() {
		if utils.Contains(name, list) {
			continue
		}
		wg.Add(1)

		go func(service string) {
			defer func() {
				wg.Done()
			}()

			logger.WithField(
				"api", service,
			).Debug("attempting to enable the api in the project")

			if err := gcputils.EnableAPI(ctx, client, id, service); err != nil {
				logger.WithError(err).Error("trying to enable the api")

				project.Status.Conditions.SetCondition(corev1.Component{
					Name:    stage,
					Detail:  err.Error(),
					Message: "Failed to enable " + service + " api in the project",
					Status:  corev1.FailureStatus,
				})
			}
		}(name)
	}
	wg.Wait()

	project.Status.Conditions.SetCondition(corev1.Component{
		Name:    stage,
		Message: "Successfully enabled all the APIs in project",
		Status:  corev1.SuccessStatus,
	})

	return nil
}

func (t ctrl) GetServiceAccountName(project *gcp.Project) string {
	return project.Status.ProjectID
}

// EnsureServiceAccount is responsible for creating the service account in the project
func (t ctrl) EnsureServiceAccount(ctx context.Context, credentials *configv1.Secret, project *gcp.Project) (*iam.ServiceAccount, error) {
	stage := "GCP Service Accounts"

	var sa *iam.ServiceAccount

	account := t.GetServiceAccountName(project)
	logger := log.WithFields(log.Fields{
		"name":       project.Name,
		"namespace":  project.Namespace,
		"account":    account,
		"project_id": project.Status.ProjectID,
	})
	logger.Debug("attempting to ensure the servie account in gcp project")

	err := func() error {
		options := option.WithCredentialsJSON([]byte(credentials.Spec.Data["key"]))

		// @step: create the iam client
		client, err := iam.NewService(ctx, options)
		if err != nil {
			logger.WithError(err).Error("trying to create the client")

			return err
		}

		// @step: create the resource client
		pclient, err := cloudresourcemanager.NewService(ctx, options)
		if err != nil {
			logger.WithError(err).Error("trying to create the cloud resource client")

			return err
		}

		path := fmt.Sprintf("projects/%s", project.Status.ProjectID)

		// @step: ensure the service account exists in the project
		list, err := client.Projects.ServiceAccounts.List(path).Context(ctx).Do()
		if err != nil {
			logger.WithError(err).Error("trying to retrieve the service account list")

			return err
		}

		var found bool
		displayName := "Kore Service Account"

		// @step: check if the service account exists
		sa, found = func() (*iam.ServiceAccount, bool) {
			for _, x := range list.Accounts {
				if x.DisplayName == displayName {
					return x, true
				}
			}

			return nil, false
		}()

		if !found {
			logger.Debug("service account does not exist, creating now")

			sa, err = client.Projects.ServiceAccounts.Create(path, &iam.CreateServiceAccountRequest{
				AccountId:      account,
				ServiceAccount: &iam.ServiceAccount{DisplayName: displayName},
			}).Context(ctx).Do()

			if err != nil {
				logger.WithError(err).Error("trying to create the service account in project")

				return err
			}
		} else {
			logger.Debug("service account already exists, skipping the creation")
		}

		bindings := []*cloudresourcemanager.Binding{
			{Role: "roles/owner", Members: []string{"serviceAccount:" + sa.Email}},
		}

		logger.Debug("attempting to set the project iam policy in the service account")

		// @step: attempt to update the project iam policy
		if err := gcputils.AddBindingsToProjectIAM(ctx, pclient, bindings, project.Status.ProjectID); err != nil {
			logger.WithError(err).Error("trying to update the project iam policy")

			return err
		}

		return nil
	}()
	if err != nil {
		logger.WithError(err).Error("attempting to provision the service account")

		project.Status.Conditions.SetCondition(corev1.Component{
			Name:    stage,
			Detail:  err.Error(),
			Message: "Failed to provision the IAM credentials in the project",
			Status:  corev1.FailureStatus,
		})

		return sa, err
	}

	project.Status.Conditions.SetCondition(corev1.Component{
		Name:    stage,
		Message: "Successfully provision the IAM in project",
		Status:  corev1.SuccessStatus,
	})

	return sa, nil
}

// EnsureServiceAccountKey is responsible for ensuring the account key exists
func (t ctrl) EnsureServiceAccountKey(
	ctx context.Context,
	credentials *configv1.Secret,
	organization *gcp.Organization,
	account *iam.ServiceAccount,
	project *gcp.Project) error {

	stage := "GCP Service Credentials"

	logger := log.WithFields(log.Fields{
		"account":   account.Email,
		"name":      project.Name,
		"namespace": project.Namespace,
	})
	var key *iam.ServiceAccountKey

	err := func() error {
		// @step: create the iam client
		client, err := iam.NewService(ctx, option.WithCredentialsJSON([]byte(credentials.Spec.Data["key"])))
		if err != nil {
			logger.WithError(err).Error("trying to create iam client for project")

			return err
		}

		path := fmt.Sprintf("projects/%s/serviceAccounts/%s", project.Status.ProjectID, account.UniqueId)

		// @step: check if the service account key exists already
		resp, err := client.Projects.ServiceAccounts.Keys.List(path).Context(ctx).Do()
		if err != nil {
			logger.WithError(err).Error("trying to check if service account key exists already")

			return err
		}
		list := resp.Keys

		// @step: a boolean used to control if we should geneate a new service key
		secretName := t.GetProjectCredentialsSecretName(project)

		// @step: update the reference for the secret
		if project.Status.CredentialRef == nil {
			project.Status.CredentialRef = &v1.SecretReference{
				Name:      secretName,
				Namespace: project.Namespace,
			}
		}

		secret := &configv1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      secretName,
				Namespace: project.Namespace,
			},
		}

		// @step: do we need to generate or regenerate the credentials
		update, err := func() (bool, error) {
			// @step: we need to first check if we have credentials
			if found, err := kubernetes.GetIfExists(ctx, t.mgr.GetClient(), secret); err != nil {
				logger.WithError(err).Error("trying to check for project credentials")

				return false, err
			} else if !found {
				logger.Debug("gcp project credentials secert not found")

				return true, nil
			}

			// @step: we need to decode the values in the secret
			if err := secret.Decode(); err != nil {
				logger.WithError(err).Error("failed to decode the values in secret")

				return true, nil
			}

			// @step: if we have an existing secret if need to verify it, we need to
			// check we have all the fields
			if err := IsCredentialsValid(secret); err != nil {
				logger.WithError(err).Error("invalid credentials secret")

				return true, nil
			}

			return false, nil
		}()
		if err != nil {
			logger.WithError(err).Error("trying to check the gcp project credentials")

			return err
		}

		if !update {
			logger.Debug("skipping the creation of the gcp project service account key")

			return nil
		}

		// @step: we need to check that we've not reached an limit on the keys and if so we
		// need to delete the oldest
		// @note: the reason for the minus one is because every service account appears to have a
		// system managed service account key which cannot be managed or deleted
		if len(list)-1 >= ServiceAccountKeyMax {
			logger.Debug("service account keys has reached the max, deleting the oldest now")

			if err := t.EnsureDeleteOldestKey(ctx, credentials, secret, account, list, project); err != nil {
				logger.WithError(err).Error("trying to delete the oldest service account key from gcp")

				return err
			}
		}

		// @step: we need to generate a service account key and store in kore
		key, err = client.Projects.ServiceAccounts.Keys.Create(path, &iam.CreateServiceAccountKeyRequest{
			KeyAlgorithm:   "KEY_ALG_RSA_2048",
			PrivateKeyType: "TYPE_GOOGLE_CREDENTIALS_FILE",
		}).Context(ctx).Do()
		if err != nil {
			logger.WithError(err).Error("trying to provision the service account key")

			return err
		}

		// @step: we need to convert the time - 2020-03-12T18:46:28Z
		tm, err := time.Parse(time.RFC3339, key.ValidBeforeTime)
		if err != nil {
			logger.WithError(err).Errorf("trying to parse the expiration time from key: %s", key.ValidBeforeTime)

			return err
		}

		// @step: we need base64 decode the credentials
		decoded, err := base64.StdEncoding.DecodeString(key.PrivateKeyData)
		if err != nil {
			logger.WithError(err).Errorf("trying to base64 decode the private key data")

			return err
		}

		// @step: populate the secret for storage
		keys := map[string]string{
			ExpiryKey:           fmt.Sprintf("%d", tm.Unix()),
			ProjectIDKey:        project.Status.ProjectID,
			ProjectNameKey:      project.Spec.ProjectName,
			ServiceAccountKey:   string(decoded),
			ServiceAccountKeyID: GetServiceAccountKeyID(key.Name),
		}
		secret = CreateCredentialsSecret(project, secretName, keys)

		// @step: create or update the credentials
		if _, err := kubernetes.CreateOrUpdate(ctx, t.mgr.GetClient(), secret.Encode()); err != nil {
			logger.WithError(err).Error("trying to update or create the credentials")

			return err
		}

		return nil
	}()
	if err != nil {
		// @TODO quick hack to sort the error in a SA not being here - we need to
		// fix up the controller to allow reconcile.Result{} instead
		if strings.Contains(err.Error(), "does not exist") {
			logger.WithError(err).Error("service account is not available yet")

			return err
		}

		project.Status.Conditions.SetCondition(corev1.Component{
			Name:    stage,
			Detail:  err.Error(),
			Message: "Failed to provision the service account key in the project",
			Status:  corev1.FailureStatus,
		})

		return err
	}

	project.Status.Conditions.SetCondition(corev1.Component{
		Name:    stage,
		Message: "Successfully provision the service account in project",
		Status:  corev1.SuccessStatus,
	})

	return nil
}

// EnsureCredentialsAllocation is responsible for creating an allocation to the credentials
func (t ctrl) EnsureCredentialsAllocation(
	ctx context.Context,
	project *gcp.Project) error {

	logger := log.WithFields(log.Fields{
		"name":      project.Name,
		"namespace": project.Namespace,
	})
	logger.Debug("attempting to create the allocation for the gcp project")

	name := t.GetAllocationName(project)

	allocation := &configv1.Allocation{
		TypeMeta: metav1.TypeMeta{
			APIVersion: configv1.GroupVersion.String(),
			Kind:       "Allocation",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: project.Namespace,
		},
		Spec: configv1.AllocationSpec{
			Name:    "gcp-" + project.Name,
			Summary: "Provides credentials to team GCP Project " + project.Spec.ProjectName,
			Resource: corev1.Ownership{
				Group:     gcp.SchemeGroupVersion.Group,
				Kind:      "Project",
				Name:      project.Name,
				Namespace: project.Namespace,
				Version:   gcp.SchemeGroupVersion.Version,
			},
			Teams: []string{project.Namespace},
		},
	}

	// @step: check if the allocation exists
	// @question: should we check for conflicting names??
	found, err := kubernetes.CheckIfExists(ctx, t.mgr.GetClient(), allocation)
	if err != nil {
		logger.WithError(err).Error("trying to check for allocation")

		return err
	}
	if found {
		return nil
	}

	if _, err := kubernetes.CreateOrUpdate(ctx, t.mgr.GetClient(), allocation); err != nil {
		logger.WithError(err).Error("trying to create the project project allocation")

		return err
	}

	return nil
}

// EnsureDeleteOldestKey is responsible for deleting the oldest key
func (t ctrl) EnsureDeleteOldestKey(
	ctx context.Context,
	credentials *configv1.Secret,
	secret *configv1.Secret,
	account *iam.ServiceAccount,
	keys []*iam.ServiceAccountKey,
	project *gcp.Project) error {

	logger := log.WithFields(log.Fields{
		"account":   account.Email,
		"name":      project.Name,
		"namespace": project.Namespace,
	})

	// @step: create a client to the iam api
	client, err := iam.NewService(ctx, option.WithCredentialsJSON([]byte(credentials.Spec.Data["key"])))
	if err != nil {
		logger.WithError(err).Error("trying to create iam client for project")

		return err
	}

	// @step: filter of the current credential first and anything which out a year 9999
	// as the rest if google managed
	var filtered []*iam.ServiceAccountKey
	for _, x := range keys {
		id := GetServiceAccountKeyID(x.Name)
		if id == secret.Spec.Data[ServiceAccountKeyID] {
			continue
		}
		tm, err := time.Parse(time.RFC3339, x.ValidBeforeTime)
		if err != nil {
			return err
		}
		// @note: this is the only way i can see to distinguish between user created keys
		// the system managed one by gcp
		if tm.Year() != 9999 {
			continue
		}

		filtered = append(filtered, x)
	}

	// @step: find the oldest key in the filtered bunch
	oldest, err := findOldestServiceAccountKey(filtered)
	if err != nil {
		logger.WithError(err).Error("trying to find the oldest service account key")

		return err
	}

	// @step: attempt to delete the credential from the api
	keyID := GetServiceAccountKeyID(oldest.Name)
	path := fmt.Sprintf("projects/%s/serviceAccounts/%s/keys/%s",
		project.Status.ProjectID,
		account.UniqueId,
		keyID)

	if _, err := client.Projects.ServiceAccounts.Keys.Delete(path).Context(ctx).Do(); err != nil {
		logger.WithError(err).Error("trying to delete the service account key")

		return err
	}

	return nil
}

// EnsureCredentialsAllocationDeleted is responsible for removing the allocation
func (t ctrl) EnsureCredentialsAllocationDeleted(
	ctx context.Context,
	project *gcp.Project) error {

	logger := log.WithFields(log.Fields{
		"name":      project.Name,
		"namespace": project.Namespace,
	})
	logger.Debug("ensuring the allocation has been removed")

	name := t.GetAllocationName(project)

	allocation := &configv1.Allocation{
		TypeMeta: metav1.TypeMeta{
			APIVersion: configv1.GroupVersion.String(),
			Kind:       "Allocation",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: project.Namespace,
		},
	}

	if err := kubernetes.DeleteIfExists(ctx, t.mgr.GetClient(), allocation); err != nil {
		logger.WithError(err).Error("trying to delete the project project allocation")

		return err
	}

	return nil
}

// EnsureCredentialsDeleted is responsible for deleting the credentials
func (t ctrl) EnsureCredentialsDeleted(
	ctx context.Context,
	project *gcp.Project) error {

	logger := log.WithFields(log.Fields{
		"name":      project.Name,
		"namespace": project.Namespace,
	})
	stage := "cleanup"

	secret := &configv1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      t.GetProjectCredentialsSecretName(project),
			Namespace: project.Namespace,
		},
	}

	// @step: delete the credentials once done
	if err := kubernetes.DeleteIfExists(ctx, t.mgr.GetClient(), secret); err != nil {
		logger.WithError(err).Error("trying to delete the gcp project credentials secret")

		project.Status.Conditions.SetCondition(corev1.Component{
			Name:    stage,
			Detail:  err.Error(),
			Message: "Failed to create a projects client, please check credentials",
			Status:  corev1.FailureStatus,
		})

		return err
	}

	return nil
}

// EnsureProjectDeleted is responsible for deleting the project if it exists
func (t ctrl) EnsureProjectDeleted(
	ctx context.Context,
	credentials *configv1.Secret,
	org *gcp.Organization,
	project *gcp.Project) error {

	logger := log.WithFields(log.Fields{
		"name":       project.Name,
		"namespace":  project.Namespace,
		"project_id": project.Status.ProjectID,
	})
	stage := "deleting"

	// @step: create the client
	client, err := cloudresourcemanager.NewService(ctx, option.WithCredentialsJSON([]byte(credentials.Spec.Data["key"])))
	if err != nil {
		logger.WithError(err).Error("trying to create the cloud resource client")

		project.Status.Conditions.SetCondition(corev1.Component{
			Name:    stage,
			Detail:  err.Error(),
			Message: "Failed to create a projects client, please check credentials",
			Status:  corev1.FailureStatus,
		})

		return err
	}

	// @step: we check if the project exists and if not create it
	resource, found, err := IsProject(ctx, client, project.Spec.ProjectName)
	if err != nil {
		logger.WithError(err).Error("trying to check for project existence")

		project.Status.Conditions.SetCondition(corev1.Component{
			Name:    stage,
			Detail:  err.Error(),
			Message: "Failed to check for project existence",
			Status:  corev1.FailureStatus,
		})

		return err
	}
	if !found {
		logger.Debug("gcp project does not exist, we can skip the rest")

		return nil
	}

	logger.Info("gcp project exists, deleting it now")

	// @step: create the project in gcp
	resp, err := client.Projects.Delete(resource.ProjectId).Context(ctx).Do()
	if err != nil {
		logger.WithError(err).Error("trying to create the gcp project")

		project.Status.Conditions.SetCondition(corev1.Component{
			Name:    stage,
			Detail:  err.Error(),
			Message: "Unable to request project deletion",
			Status:  corev1.FailureStatus,
		})

		return err
	}

	if resp.HTTPStatusCode < 200 || resp.HTTPStatusCode > 299 {
		project.Status.Conditions.SetCondition(corev1.Component{
			Name:    stage,
			Detail:  fmt.Sprintf("Response code back for GCP was %d", resp.HTTPStatusCode),
			Message: "GCP has responded with an unable to delete projet",
			Status:  corev1.FailureStatus,
		})

		return errors.New("invalid delete response received from gcp")
	}

	logger.Debug("successfully deleted the project from gcp")

	return nil
}

// GetProjectCredentialsSecretName returns the secret name of the credentials for this project
func (t ctrl) GetProjectCredentialsSecretName(project *gcp.Project) string {
	return fmt.Sprintf("gcp-%s", project.Name)
}

// GetRequiredAPI returns a list of required apis
func (t ctrl) GetRequiredAPI() []string {
	return []string{
		"cloudbilling.googleapis.com",
		"cloudresourcemanager.googleapis.com",
		"compute.googleapis.com",
		"container.googleapis.com",
		"iam.googleapis.com",
		"serviceusage.googleapis.com",
	}
}

// IsProjected checks if the project name has already been projected by another team
func (t ctrl) IsProjected(ctx context.Context, project *gcp.Project) (bool, error) {
	list := &gcp.ProjectList{}

	if err := t.mgr.GetClient().List(ctx, list, client.InNamespace("")); err != nil {
		return false, err
	}

	// @step: we iterate the list and look for any projects with the same name
	// but NOT in our namespace
	for _, x := range list.Items {
		if x.Namespace == project.Namespace && x.Name == project.Name {
			continue
		}

		if x.Spec.ProjectName == project.Spec.ProjectName || x.Status.ProjectID == project.Spec.ProjectName {
			return true, nil
		}
	}

	return false, nil
}

// GetAllocationName returns the name we should use for the project allocation
func (t ctrl) GetAllocationName(project *gcp.Project) string {
	return fmt.Sprintf("gcp-%s", project.Name)
}
