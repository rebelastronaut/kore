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

package helpers

import (
	configv1 "github.com/appvia/kore/pkg/apis/config/v1"
	"github.com/appvia/kore/pkg/utils/cloud/aws"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreateSecretRef creates and returns a secret
func CreateSecretRef(namespace, name string) *configv1.Secret {
	return &configv1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
	}
}

// GetAWSCreds will get the AWS credentials from a kore secret and account ID
func GetAWSCreds(credentials *configv1.Secret, accountID string) (*aws.Credentials, error) {
	err := credentials.Decode()
	if err != nil {
		return nil, err
	}
	return &aws.Credentials{
		AccountID:       accountID,
		AccessKeyID:     credentials.Spec.Data["access_key_id"],
		SecretAccessKey: credentials.Spec.Data["access_secret_key"],
	}, nil
}
