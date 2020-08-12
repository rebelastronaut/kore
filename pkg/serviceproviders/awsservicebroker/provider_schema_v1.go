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

const providerSchemaV1 = `{
	"$id": "https://appvia.io/kore/schemas/serviceprovider/aws-servicebroker/v1.json",
	"$schema": "http://json-schema.org/draft-07/schema#",
	"description": "This is a custom service provider for aws-servicebroker (https://github.com/awslabs/aws-servicebroker)",
	"type": "object",
	"additionalProperties": false,
	"required": [
		"awsAccessKeyID",
		"awsSecretAccessKey"
	],
	"properties": {
		"chartRepositoryType": {
			"type": "string",
			"enum": ["git", "helm"],
			"title": "Chart Repository Type",
			"default": "git",
			"description": "Helm repository type to use (git or helm)"
		},
		"chartRepository": {
			"type": "string",
			"minLength": 1,
			"title": "Chart Repository",
			"default": "https://github.com/appvia/aws-servicebroker",
			"description": "the repository URL of the Helm chart for the aws-servicebroker"
		},
		"chartVersion": {
			"type": "string",
			"minLength": 1,
			"title": "Chart Version",
			"description": "the version of the Helm chart for the aws-servicebroker"
		},
		"chartRepositoryRef": {
			"type": "string",
			"minLength": 1,
			"title": "Chart Repository URL",
			"description": "the Helm repository URL for the aws-servicebroker"
		},
		"chartRepositoryPath": {
			"type": "string",
			"minLength": 1,
			"title": "Chart Repository Path",
			"default": "packaging/helm/aws-servicebroker",
			"description": "the path to the chart relative to the repository root"
		},
		"region": {
			"type": "string",
			"default": "us-east-1",
			"minLength": 1,
			"title": "Region",
			"description": "the AWS region where the AWS dependencies will be created"
		},
		"tableName": {
			"type": "string",
			"default": "aws-service-broker",
			"minLength": 1,
			"title": "DynamoDB Table Name",
			"description": "the DynamoDB table name where state will be stored"
		},
		"s3BucketName": {
			"type": "string",
			"default": "awsservicebroker",
			"minLength": 1,
			"title": "S3 Bucket Name",
			"description": "the name of the S3 bucket used to store the CloudFormation templates for the service plans"
		},
		"s3BucketRegion": {
			"type": "string",
			"default": "us-east-1",
			"minLength": 1,
			"title": "S3 Bucket Region",
			"description": "the region of the S3 bucket used to store the CloudFormation templates for the service plans"
		},
		"s3BucketKey": {
			"type": "string",
			"default": "templates/latest/",
			"title": "S3 Bucket Path",
			"description": "the path in the S3 bucket used to store the CloudFormation templates for the service plans"
		},
		"awsAccessKeyID": {
			"type": "string",
			"minLength": 1,
			"title": "AWS Access Key ID",
			"description": "the AWS access key id used to create the required AWS resources"
		},
		"awsSecretAccessKey": {
			"type": "string",
			"minLength": 1,
			"title": "AWS Secret Access Key",
			"description": "the AWS secret access key used to create the required AWS resources"
		},
		"awsIAMRoleName": {
			"type": "string",
			"minLength": 1,
			"title": "AWS IAM Role Name",
			"default": "aws-service-broker",
			"description": "the IAM role name to assume when provisiong resources"
		},
		"allowEmptyCredentialSchema": {
			"type": "boolean",
			"default": false,
			"title": "Allow Empty Credential Schema",
			"description": "if true the service credentials schema doesn't have to be defined"
		},
		"defaultPlans": {
			"type": "array",
			"items": {
				"type": "string",
				"minLength": 1
			},
			"title": "Default plans",
			"description": "a list of plan names to use as default for each service kind in a format as [kind]-[plan name]"
		},
		"includeKinds": {
			"type": "array",
			"items": {
				"type": "string",
				"minLength": 1
			},
			"title": "Include Service Kinds",
			"description": "a list of service kinds to include from the catalog"
		},
		"excludeKinds": {
			"type": "array",
			"items": {
				"type": "string",
				"minLength": 1
			},
			"title": "Exclude Service Kinds",
			"description": "a list of service kinds to exclude from the catalog"
		},
		"includePlans": {
			"type": "array",
			"items": {
				"type": "string",
				"minLength": 1
			},
			"title": "Include Service Plans",
			"description": "a list of service plan names ([kind]-[plan name]) to include from the catalog"
		},
		"excludePlans": {
			"type": "array",
			"items": {
				"type": "string",
				"minLength": 1
			},
			"title": "Exclude Service Plans",
			"description": "a list of service plan names ([kind]-[plan name]) to exclude from the catalog"
		}
	}
}`
