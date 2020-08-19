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

package aws

import (
	"context"
	"errors"
	"reflect"
	"strings"

	eksv1alpha1 "github.com/appvia/kore/pkg/apis/eks/v1alpha1"
	"github.com/appvia/kore/pkg/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/eks"
	awseks "github.com/aws/aws-sdk-go/service/eks"
	log "github.com/sirupsen/logrus"
)

var (
	// ErrClusterNotFound indicates the cluster does not exist
	ErrClusterNotFound = errors.New("eks cluster not found")
	// ErrNodeGroupNotFound indicates the nodegroup does not exist
	ErrNodeGroupNotFound = errors.New("eks nodegroup not found")
	// ErrResourceBusy indicate the resource is currently busy performing an operation
	ErrResourceBusy = errors.New("resource is busy performing an operation (upgrade, creating)")
)

// Client for aws EKS and EKS nodegroups
type Client struct {
	// cluster is the API object used
	cluster *eksv1alpha1.EKS
	// clusterName is the eks cluster name
	clusterName string
	// Sess is the AWS session
	Sess *session.Session
	// svc is the eks service
	svc *eks.EKS
}

// NewBasicClient gets an AWS session relating to a cluster
// TODO: maybe remove after refactor of nodegroup to use clusterref?
func NewBasicClient(creds *Credentials, clusterName, region string) (*Client, error) {
	sesh, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(creds.AccessKeyID, creds.SecretAccessKey, ""),
	})
	if err != nil {

		return nil, err
	}

	return &Client{
		clusterName: clusterName,
		Sess:        sesh,
		svc:         eks.New(sesh),
	}, err
}

// NewEKSClient gets an AWS and cluster session with a reference to our API object
func NewEKSClient(cred *Credentials, cluster *eksv1alpha1.EKS) (*Client, error) {

	sesh, err := session.NewSession(&aws.Config{
		Region: aws.String(cluster.Spec.Region),
		Credentials: credentials.NewStaticCredentials(
			cred.AccessKeyID,
			cred.SecretAccessKey,
			"",
		),
	})
	if err != nil {
		return nil, err
	}

	return &Client{
		clusterName: cluster.Name,
		cluster:     cluster,
		Sess:        sesh,
		svc:         eks.New(sesh),
	}, err
}

// NewEKSClientFromVPC will create a new eks client from an VPCClient object
func NewEKSClientFromVPC(c *VPCClient, clusterName string) *Client {
	return &Client{
		clusterName: clusterName,
		Sess:        c.Sess,
		svc:         eks.New(c.Sess),
	}
}

// GetIfExists checks if a cluster exists and returns it
func (c *Client) GetIfExists(ctx context.Context) (*awseks.Cluster, bool, error) {
	out, err := c.svc.DescribeClusterWithContext(ctx, &awseks.DescribeClusterInput{
		Name: aws.String(c.clusterName),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case eks.ErrCodeResourceNotFoundException:
				return nil, false, nil
			default:
				return nil, false, err
			}
		} else {
			return nil, false, err
		}
	}

	return out.Cluster, true, nil
}

// Create creates an EKS cluster
func (c *Client) Create(ctx context.Context) (*eks.Cluster, error) {
	logger := log.WithFields(log.Fields{
		"name":      c.cluster.Name,
		"namespace": c.cluster.Namespace,
	})
	logger.Debug("attempting to create the eks cluster")

	output, err := c.svc.CreateClusterWithContext(ctx, c.createClusterInput())
	if err != nil {
		return nil, err
	}

	return output.Cluster, err
}

// Delete is responsible for deleting the eks cluster
func (c *Client) Delete(ctx context.Context) error {
	logger := log.WithFields(log.Fields{
		"name":      c.cluster.Name,
		"namespace": c.cluster.Namespace,
	})
	logger.Debug("attempting to delete the eks cluster")

	// @step: get the state of the cluster
	_, err := c.svc.DeleteClusterWithContext(ctx, &eks.DeleteClusterInput{
		Name: aws.String(c.cluster.Name),
	})
	if err != nil {
		if c.IsNotFound(err) {
			return nil
		}
		logger.WithError(err).Error("trying to delete the eks cluster")

		return err
	}

	return nil
}

// Update should migrate changes to a cluster object
func (c *Client) Update(ctx context.Context) (bool, error) {
	logger := log.WithFields(log.Fields{
		"name":      c.cluster.Name,
		"namespace": c.cluster.Namespace,
	})
	logger.Debug("checking if the cluster requires an update")

	// @step: retrieve the current state of the cluster
	state, err := c.Describe(ctx)
	if err != nil {
		logger.WithError(err).Error("trying to describe the cluster")

		return false, err
	}

	// @step: Check cluster version up to date with spec
	if c.cluster.Spec.Version != "" {
		// @TODO we need to check the semvar and never try and downgrade??
		if aws.StringValue(state.Version) != c.cluster.Spec.Version {
			logger.Debug("cluster version is out of sync, attempting to update")

			if _, err := c.svc.UpdateClusterVersionWithContext(ctx, &awseks.UpdateClusterVersionInput{
				Name:    aws.String(c.cluster.Name),
				Version: aws.String(c.cluster.Spec.Version),
			}); err != nil {
				logger.WithError(err).Error("trying to request a version update")

				return false, err
			}

			return true, nil
		}
	}

	// @step: Check cluster tags
	tagsUpdate := &awseks.TagResourceInput{
		Tags:        map[string]*string{},
		ResourceArn: state.Arn,
	}
	for k, v := range c.cluster.Spec.Tags {
		if state.Tags == nil || state.Tags[k] == nil || *state.Tags[k] != v {
			tagsUpdate.Tags[k] = aws.String(v)
		}
	}
	if len(tagsUpdate.Tags) > 0 {
		logger.Info("eks cluster tagging needs updating, attempting to sync")

		if _, err := c.svc.TagResourceWithContext(ctx, tagsUpdate); err != nil {
			return false, err
		}
	}

	// @step: Check cluster VPC config
	update := &awseks.UpdateClusterConfigInput{
		Name:               aws.String(c.cluster.Name),
		ResourcesVpcConfig: &awseks.VpcConfigRequest{},
	}

	if c.cluster.Spec.AuthorizedMasterNetworks != nil {
		// @step: have the public ranges changed for the endpoint?
		if !reflect.DeepEqual(
			utils.StringsSorted(aws.StringValueSlice(state.ResourcesVpcConfig.PublicAccessCidrs)),
			utils.StringsSorted(c.cluster.Spec.AuthorizedMasterNetworks),
		) {
			update.ResourcesVpcConfig.PublicAccessCidrs = aws.StringSlice(c.cluster.Spec.AuthorizedMasterNetworks)
		}
	}

	publicEndpoint := aws.BoolValue(state.ResourcesVpcConfig.EndpointPublicAccess)

	// @check if the public endpoint has changed
	switch {
	case utils.BoolValue(c.cluster.Spec.EnablePrivateNetwork) && publicEndpoint:
		update.ResourcesVpcConfig.EndpointPublicAccess = aws.Bool(false)
	case !utils.BoolValue(c.cluster.Spec.EnablePrivateNetwork) && !publicEndpoint:
		update.ResourcesVpcConfig.EndpointPublicAccess = aws.Bool(true)
	}

	// @check if the private endpoint has changed
	if !aws.BoolValue(state.ResourcesVpcConfig.EndpointPrivateAccess) {
		update.ResourcesVpcConfig.EndpointPrivateAccess = aws.Bool(true)
	}

	// has anything been changed?
	if utils.IsEmpty(update.ResourcesVpcConfig) {
		return false, nil
	}

	logger.Debug("eks cluster vpc configuration has drifted, attempting to sync")

	if _, err := c.svc.UpdateClusterConfigWithContext(ctx, update); err != nil {
		return false, err
	}

	return false, nil
}

// VerifyCredentials is responsible for verifying AWS creds
func (c *Client) VerifyCredentials() error {
	// TODO: see https://github.com/appvia/kore/issues/498

	return nil
}

// Describe returns the AWS EKS output
func (c *Client) Describe(ctx context.Context) (*eks.Cluster, error) {
	d, err := c.svc.DescribeClusterWithContext(ctx, &awseks.DescribeClusterInput{
		Name: aws.String(c.clusterName),
	})
	if err != nil {
		return nil, err
	}

	return d.Cluster, nil
}

// DeleteNodeGroup will remove a nodegroup from a cluster
func (c *Client) DeleteNodeGroup(ctx context.Context, group *eksv1alpha1.EKSNodeGroup) error {
	logger := log.WithFields(log.Fields{
		"name":      group.Name,
		"namespace": group.Namespace,
	})
	logger.Debug("attempting to delete the eks nodegroup")

	if _, err := c.svc.DeleteNodegroupWithContext(ctx, &eks.DeleteNodegroupInput{
		ClusterName:   aws.String(group.Spec.Cluster.Name),
		NodegroupName: aws.String(group.Name),
	}); err != nil {
		return err
	}

	return nil
}

// CreateNodeGroup will create a node group for the EKS cluster
func (c *Client) CreateNodeGroup(ctx context.Context, group *eksv1alpha1.EKSNodeGroup) error {
	// @step: check if the nodegroup exists already
	existing, err := c.NodeGroupExists(ctx, group)
	if err != nil {
		return err
	}
	if !existing {
		scalingConfig := &eks.NodegroupScalingConfig{
			DesiredSize: aws.Int64(group.Spec.DesiredSize),
			MaxSize:     aws.Int64(group.Spec.MaxSize),
			MinSize:     aws.Int64(group.Spec.MinSize),
		}
		// If we're not auto-scaling, we need to keep the min and max the same as desired to
		// force the size to be managed by kore - else any manual node group scale up/downs
		// directly in AWS will simply be undone by kore.
		if !group.Spec.EnableAutoscaler {
			scalingConfig.MaxSize = scalingConfig.DesiredSize
			scalingConfig.MinSize = scalingConfig.DesiredSize
		}
		input := &eks.CreateNodegroupInput{
			AmiType:       aws.String(group.Spec.AMIType),
			ClusterName:   aws.String(group.Spec.Cluster.Name),
			DiskSize:      aws.Int64(group.Spec.DiskSize),
			InstanceTypes: aws.StringSlice([]string{group.Spec.InstanceType}),
			NodeRole:      aws.String(group.Status.NodeIAMRole),
			NodegroupName: aws.String(group.Name),
			Subnets:       aws.StringSlice(group.Spec.Subnets),
			Version:       aws.String(group.Spec.Version),
			ScalingConfig: scalingConfig,
		}
		if group.Spec.EC2SSHKey != "" {
			input.RemoteAccess = &eks.RemoteAccessConfig{
				Ec2SshKey:            aws.String(group.Spec.EC2SSHKey),
				SourceSecurityGroups: aws.StringSlice(group.Spec.SSHSourceSecurityGroups),
			}
		}
		if len(group.Spec.Tags) > 0 {
			input.Tags = aws.StringMap(group.Spec.Tags)
		}
		if len(group.Spec.Labels) > 0 {
			input.Labels = aws.StringMap(group.Spec.Labels)
		}

		if _, err := c.svc.CreateNodegroup(input); err != nil {
			return err
		}
	}

	return nil
}

// NodeGroupExists is responsible for checking if the nodegroup exists
func (c *Client) NodeGroupExists(ctx context.Context, nodegroup *eksv1alpha1.EKSNodeGroup) (exists bool, err error) {
	_, err = c.svc.DescribeNodegroupWithContext(ctx, &eks.DescribeNodegroupInput{
		ClusterName:   aws.String(nodegroup.Spec.Cluster.Name),
		NodegroupName: aws.String(nodegroup.Name),
	})
	if err != nil {
		if !c.IsNotFound(err) {
			return false, err
		}

		return false, nil
	}

	return true, nil
}

// DescribeNodeGroup retrieve the nodegroup
func (c *Client) DescribeNodeGroup(ctx context.Context, group *eksv1alpha1.EKSNodeGroup) (*awseks.Nodegroup, error) {
	req, err := c.svc.DescribeNodegroupWithContext(ctx, &eks.DescribeNodegroupInput{
		ClusterName:   aws.String(group.Spec.Cluster.Name),
		NodegroupName: aws.String(group.Name),
	})
	if err != nil {
		return nil, err
	}
	if req.Nodegroup == nil {
		return nil, ErrNodeGroupNotFound
	}

	return req.Nodegroup, nil
}

// UpdateNodeGroup is responsible for checking for a drift and applying an update if required
func (c *Client) UpdateNodeGroup(ctx context.Context, group *eksv1alpha1.EKSNodeGroup) (bool, error) {
	logger := log.WithFields(log.Fields{
		"name":      group.Name,
		"namespace": group.Namespace,
	})
	state, err := c.DescribeNodeGroup(ctx, group)
	if err != nil {
		logger.WithError(err).Error("trying to retrieve the eks nodegroup")

		return false, err
	}

	if group.Spec.Version != "" && group.Spec.Version != aws.StringValue(state.Version) {
		logger.WithFields(log.Fields{
			"current":  aws.StringValue(state.Version),
			"expected": group.Spec.Version,
		}).Debug("attempting to update the nodegroup node version")

		if _, err := c.svc.UpdateNodegroupVersionWithContext(ctx, &awseks.UpdateNodegroupVersionInput{
			ClusterName:   aws.String(group.Spec.Cluster.Name),
			Force:         aws.Bool(true),
			NodegroupName: aws.String(group.Name),
			Version:       aws.String(group.Spec.Version),
		}); err != nil {
			logger.WithError(err).Error("trying to updade the node version")

			return false, err
		}

		return true, nil
	}

	maxSize := group.Spec.MaxSize
	minSize := group.Spec.MinSize
	desiredSize := group.Spec.DesiredSize
	if !group.Spec.EnableAutoscaler {
		// If we're not auto-scaling, we need to keep the min and max the same as desired to
		// force the size to be managed by kore - else any manual node group scale up/downs
		// directly in AWS will simply be undone by kore.
		maxSize = group.Spec.DesiredSize
		minSize = group.Spec.DesiredSize
	} else {
		// If we are auto-scaling, we should leave the desired size well alone! Don't want
		// to be fighting with the auto-scaler no matter what our spec says... this leaves
		// a VERY small window between retrieving the state above and setting the state below
		// where if an auto-scale happens and our spec has a changed max/min size at the same
		// time, you'd get a slightly odd result here where we might set an old desired size.
		// But it's a very small window indeed, and will sort itself out fairly promptly.
		desiredSize = aws.Int64Value(state.ScalingConfig.DesiredSize)

		// We have to make sure the desired size is always within min and max
		if desiredSize < minSize {
			desiredSize = minSize
		} else if desiredSize > maxSize {
			desiredSize = maxSize
		}
	}
	if aws.Int64Value(state.ScalingConfig.MinSize) != minSize ||
		aws.Int64Value(state.ScalingConfig.MaxSize) != maxSize ||
		aws.Int64Value(state.ScalingConfig.DesiredSize) != desiredSize {

		if _, err := c.svc.UpdateNodegroupConfigWithContext(ctx, &awseks.UpdateNodegroupConfigInput{
			ClusterName:   aws.String(group.Spec.Cluster.Name),
			NodegroupName: aws.String(group.Name),
			ScalingConfig: &awseks.NodegroupScalingConfig{
				DesiredSize: aws.Int64(desiredSize),
				MinSize:     aws.Int64(minSize),
				MaxSize:     aws.Int64(maxSize),
			},
		}); err != nil {
			return false, err
		}

		return true, nil
	}

	return false, nil
}

// ListNodeGroups get a list of the nodegroups
func (c *Client) ListNodeGroups() (nodegroups []string, err error) {
	nodegroups = make([]string, 0)
	ngo, err := c.svc.ListNodegroups(&eks.ListNodegroupsInput{
		ClusterName: &c.clusterName,
	})
	if err != nil {

		return nodegroups, err
	}
	for _, ng := range ngo.Nodegroups {
		nodegroups = append(nodegroups, *ng)
	}

	return nodegroups, nil
}

// GetEKSNodeGroupStatus the status of an existing node group
func (c *Client) GetEKSNodeGroupStatus(nodegroup *eksv1alpha1.EKSNodeGroup) (status string, err error) {
	out, err := c.svc.DescribeNodegroup(&eks.DescribeNodegroupInput{
		ClusterName:   &c.clusterName,
		NodegroupName: &nodegroup.Name,
	})

	return *out.Nodegroup.Status, err
}

// ListEKSNodeGroupNodes returns a list of the current nodes which make up this node group
func (c *Client) ListEKSNodeGroupNodes(group *eksv1alpha1.EKSNodeGroup) ([]*ec2.Instance, error) {
	ec2client := ec2.New(c.Sess)
	instances, err := ec2client.DescribeInstances(&ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("instance-state-name"),
				Values: aws.StringSlice([]string{"running"}),
			},
			{
				Name:   aws.String("tag:eks:cluster-name"),
				Values: []*string{aws.String(group.Spec.Cluster.Name)},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	var ret []*ec2.Instance
	for _, r := range instances.Reservations {
		ret = append(ret, r.Instances...)
	}
	return ret, nil
}

// SetEKSNodeGroupNodeTags requests the supplied set of tags to be set on all of the supplied instance IDs
func (c *Client) SetEKSNodeGroupNodeTags(ids []string, tags map[string]string) error {
	var ec2tags []*ec2.Tag
	for k, v := range tags {
		ec2tags = append(ec2tags, &ec2.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		})
	}
	ec2client := ec2.New(c.Sess)
	_, err := ec2client.CreateTags(&ec2.CreateTagsInput{
		Resources: aws.StringSlice(ids),
		Tags:      ec2tags,
	})
	return err
}

// HasTagWithValue checks whether the supplied instance is tagged with the specified key/value pair
func (c *Client) HasTagWithValue(instance *ec2.Instance, key string, value string) bool {
	for _, t := range instance.Tags {
		if *t.Key == key && *t.Value == value {
			return true
		}
	}
	return false
}

// createClusterInput is used to generate the EKS cluster definition
func (c *Client) createClusterInput() *awseks.CreateClusterInput {
	tags := map[string]*string{
		utils.Label("name"):  aws.String(c.cluster.Name),
		utils.Label("owned"): aws.String("true"),
		utils.Label("team"):  aws.String(c.cluster.Namespace),
	}
	for k, v := range c.cluster.Spec.Tags {
		tags[k] = aws.String(v)
	}

	enablePublicEndpoint := true
	if utils.BoolValue(c.cluster.Spec.EnablePrivateNetwork) {
		enablePublicEndpoint = false
	}

	d := &awseks.CreateClusterInput{
		Name:    aws.String(c.cluster.Name),
		RoleArn: aws.String(c.cluster.Status.RoleARN),
		Version: aws.String(c.cluster.Spec.Version),
		ResourcesVpcConfig: &awseks.VpcConfigRequest{
			SecurityGroupIds:      aws.StringSlice(c.cluster.Spec.SecurityGroupIDs),
			SubnetIds:             aws.StringSlice(c.cluster.Spec.SubnetIDs),
			EndpointPublicAccess:  aws.Bool(enablePublicEndpoint),
			EndpointPrivateAccess: aws.Bool(true),
		},
		Tags: tags,
	}

	for _, x := range c.cluster.Spec.AuthorizedMasterNetworks {
		d.ResourcesVpcConfig.PublicAccessCidrs = append(d.ResourcesVpcConfig.PublicAccessCidrs, aws.String(x))
	}

	return d
}

// IsNotFound checks if the aws error was an not found resource
func (c *Client) IsNotFound(err error) bool {
	if aerr, ok := err.(awserr.Error); ok {
		if aerr.Code() == eks.ErrCodeResourceNotFoundException {
			return true
		}
	}

	return false
}

// IsInvalidParameterException checks if the error was a invalid parameter
func (c *Client) IsInvalidParameterException(err error) bool {
	if aerr, ok := err.(awserr.Error); ok {
		if aerr.Code() == eks.ErrCodeInvalidParameterException && strings.Contains(aerr.Message(), "does not exist") {
			return true
		}
	}

	return false
}
