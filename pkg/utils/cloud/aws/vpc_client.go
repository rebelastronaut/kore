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
	"fmt"
	"net"
	"sort"
	"strconv"
	"strings"

	"github.com/appvia/kore/pkg/utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	log "github.com/sirupsen/logrus"
)

const (
	AZLimit = 3
)

// VPCClient for aws VPC
type VPCClient struct {
	// credentials are the aws credentials
	credentials Credentials
	// Sess is the AWS session
	Sess *session.Session
	// The svc to access EC2 resources
	svc *ec2.EC2
	// The VPC to act on
	VPC VPC
}

// NewVPCClient gets an AWS and session with a reference to a matching VPC
func NewVPCClient(creds Credentials, vpc VPC) (*VPCClient, error) {
	sess := getNewSession(creds, vpc.Region)

	// TODO: verify the current CIDR is big enough for the required subnets given:
	// - the constants PrivateNetworkMaskSize, PublicNetworkMaskSize
	// - the expected number of AZ's to use...
	// for now it must be a /15 (assuming three az's)
	_, _, err := net.ParseCIDR(vpc.CidrBlock)
	if err != nil {
		return nil, fmt.Errorf("invalid netmask provided %s", vpc.CidrBlock)
	}
	bitStr := strings.Split(vpc.CidrBlock, "/")[1]
	bits, _ := strconv.ParseInt(bitStr, 10, 8)
	if bits < 16 {
		return nil, fmt.Errorf("vpc cidr too small to create 3x /%d and 3x /%d subnets", PrivateNetworkMaskSize, PublicNetworkMaskSize)
	}
	return &VPCClient{
		credentials: creds,
		Sess:        sess,
		VPC:         vpc,
		svc:         ec2.New(sess),
	}, nil
}

// Ensure will create or update a VPC with ALL required global resources
func (c *VPCClient) Ensure() (ready bool, _ *VPCResult, _ error) {
	// Check if the VPC exists
	found, err := c.Exists()
	if err != nil {
		return false, nil, err
	}
	// Now check it's resources global resources exist
	if !found {
		// time to create
		o, err := c.svc.CreateVpc(&ec2.CreateVpcInput{CidrBlock: aws.String(c.VPC.CidrBlock)})
		if err != nil {
			return false, nil, fmt.Errorf("error creating a new aws vpc %s - %s", c.VPC.Name, err)
		}
		err = createTags(
			*c.svc,
			c.VPC.Name,
			*o.Vpc.VpcId,
			c.VPC.Tags,
		)
		if err != nil {
			return false, nil, fmt.Errorf("error tagging new aws vpc %s, id %s - %s", c.VPC.Name, *o.Vpc.VpcId, err)
		}
		c.VPC.awsObj = o.Vpc
	}

	// Next ensure VPC params set - EnableDnsSupport
	_, err = c.svc.ModifyVpcAttribute(&ec2.ModifyVpcAttributeInput{
		EnableDnsSupport: &ec2.AttributeBooleanValue{
			Value: aws.Bool(true),
		},
		VpcId: c.VPC.awsObj.VpcId,
	})
	if err != nil {
		return false, nil, err
	}
	// Next ensure VPC params set - EnableDnsHostnames
	// Only one at a time, see https://github.com/aws/aws-sdk-go/issues/415
	_, err = c.svc.ModifyVpcAttribute(&ec2.ModifyVpcAttributeInput{
		EnableDnsHostnames: &ec2.AttributeBooleanValue{
			Value: aws.Bool(true),
		},
		VpcId: c.VPC.awsObj.VpcId,
	})
	if err != nil {
		return false, nil, err
	}

	// ensure we have an internet gateway and attach
	igw, err := EnsureInternetGateway(*c.svc, c.VPC)
	if err != nil {
		return false, nil, err
	}

	azs, err := c.getAZs(AZLimit)
	if err != nil {
		return false, nil, err
	}

	// First discover any public subnets or create
	// The public networks will use the very first subnets from the VPC
	vpcStartIP, _, err := net.ParseCIDR(c.VPC.CidrBlock)
	if err != nil {
		return false, nil, err
	}
	publicSubnets, err := EnsurePublicSubnets(
		*c.svc,
		c.VPC,
		azs,
		vpcStartIP,
		*igw.InternetGatewayId,
	)
	if err != nil {
		return false, nil, err
	}

	// Get next network address for the internal subnets from the last of the public addresses
	// Get last networ address from public addresses:
	lastPublicSubnet := publicSubnets[len(publicSubnets)-1]
	_, lastPublicNet, err := net.ParseCIDR(*lastPublicSubnet.CidrBlock)
	if err != nil {
		return false, nil, fmt.Errorf("bad ciddr on last aws public subnet %s - %s", *lastPublicSubnet.CidrBlock, err)
	}
	// Get next network of the private size from the last public network
	privateNet, err := utils.GetSubnetFromLast(lastPublicNet, PrivateNetworkMaskSize)
	if err != nil {
		return false, nil, fmt.Errorf("error trying to work next subnet of size %d from %s - %s", PrivateNetworkMaskSize, *lastPublicSubnet.CidrBlock, err)
	}

	privateSubnets, natGateways, ready, err := EnsurePrivateSubnets(*c.svc, c.VPC, azs, privateNet.IP, publicSubnets)
	if err != nil || !ready {
		return ready, nil, err
	}

	// create security group for master control plane...
	securityGroup, err := EnsureSecurityGroup(*c.svc, c.VPC, SecurityGroupTypeEKSCluster, "eks required group for allowing communication with master nodes")
	if err != nil {
		return false, nil, fmt.Errorf("error finding or creating security group for eks master comms - %s", err)
	}

	return true, &VPCResult{
		VPC:                         c.VPC.awsObj,
		PublicSubnets:               publicSubnets,
		PrivateSubnets:              privateSubnets,
		NATGateways:                 natGateways,
		ControlPlaneSecurityGroupID: aws.StringValue(securityGroup.GroupId),
	}, nil
}

// Exists checks if a vpc exists
func (c *VPCClient) Exists() (bool, error) {
	if c.VPC.awsObj != nil {

		return true, nil
	}
	o, err := c.svc.DescribeVpcs(&ec2.DescribeVpcsInput{
		Filters: []*ec2.Filter{getEc2TagNameFilter(c.VPC.Name)},
	})
	if err != nil {

		return false, err
	}
	if len(o.Vpcs) == 1 {
		// Cache the VPC
		c.VPC.awsObj = o.Vpcs[0]

		return true, nil
	}
	if len(o.Vpcs) > 1 {

		return false, fmt.Errorf("Multiple matching VPCs")
	}

	return false, nil
}

// Delete will clear up all VPC resources
func (c *VPCClient) Delete(ctx context.Context) (ready bool, _ error) {
	exists, err := c.Exists()
	if err != nil {
		return false, err
	}
	if !exists {
		return true, nil
	}

	if !IsKoreManaged(c.VPC.awsObj.Tags) {
		return true, nil
	}

	azs, err := c.getAZs(AZLimit)
	if err != nil {
		return false, err
	}

	// @step: we need to delete any lingering ENI which have not been cleaned up
	if err := DeleteLingeringResources(ctx, c); err != nil {
		return false, err
	}
	ready, err = DeletePrivateSubnets(*c.svc, c.VPC, azs)
	if err != nil || !ready {
		return ready, err
	}
	if err := DeletePublicSubnets(*c.svc, c.VPC, azs); err != nil {
		return false, err
	}
	if err := DeleteInternetGateway(*c.svc, c.VPC); err != nil {
		return false, err
	}
	if err := DeleteSecurityGroup(*c.svc, c.VPC, SecurityGroupTypeEKSCluster); err != nil {
		return false, err
	}

	_, err = c.svc.DeleteVpc(&ec2.DeleteVpcInput{
		VpcId: c.VPC.awsObj.VpcId,
	})
	if err != nil {
		return false, fmt.Errorf("failed to delete VPC %s: %w", c.VPC.Name, err)
	}

	return true, nil
}

// DeleteLingeringResources is responsible for fixing up what aws does not
// https://github.com/aws/amazon-vpc-cni-k8s/issues/69
// https://github.com/weaveworks/eksctl/issues/1325
func DeleteLingeringResources(ctx context.Context, client *VPCClient) error {
	if err := DeleteLingeringENI(ctx, client); err != nil {
		return err
	}

	return DeleteLingeringSecurityGroups(ctx, client)
}

// DeleteLingeringSecurityGroups is related to https://github.com/aws/amazon-vpc-cni-k8s/issues/69
func DeleteLingeringSecurityGroups(ctx context.Context, client *VPCClient) error {
	vpcid := aws.StringValue(client.VPC.awsObj.VpcId)

	resp, err := client.svc.DescribeSecurityGroupsWithContext(ctx, &ec2.DescribeSecurityGroupsInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("tag:aws:eks:cluster-name"),
				Values: aws.StringSlice([]string{client.VPC.Name}),
			},
			{
				Name:   aws.String("vpc-id"),
				Values: aws.StringSlice([]string{vpcid}),
			},
		},
	})
	if err != nil {
		return err
	}

	for _, x := range resp.SecurityGroups {
		log.WithFields(log.Fields{
			"securitygroup-id": aws.StringValue(x.GroupId),
			"vpc-id":           aws.StringValue(client.VPC.awsObj.VpcId),
		}).Debug("deleting the lingering security group")

		if _, err := client.svc.DeleteSecurityGroupWithContext(ctx, &ec2.DeleteSecurityGroupInput{
			GroupId: x.GroupId,
		}); err != nil {
			return err
		}
	}

	return nil
}

// DeleteLingeringENI removes any rouge node ENIs from the VPC
// https://github.com/aws/amazon-vpc-cni-k8s/issues/69
func DeleteLingeringENI(ctx context.Context, client *VPCClient) error {
	vpcid := aws.StringValue(client.VPC.awsObj.VpcId)

	logger := log.WithFields(log.Fields{
		"vpc-id": vpcid,
	})
	logger.Debug("checking for any lingering eni")

	// @step: we find any lingering network interfaces on the VPC
	resp, err := client.svc.DescribeNetworkInterfacesWithContext(ctx, &ec2.DescribeNetworkInterfacesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("tag-key"),
				Values: aws.StringSlice([]string{"node.k8s.amazonaws.com/instance_id"}),
			},
			{
				Name:   aws.String("vpc-id"),
				Values: aws.StringSlice([]string{vpcid}),
			},
		},
	})
	if err != nil {
		return err
	}

	logger.WithField("size", len(resp.NetworkInterfaces)).Debug("found the following attached eni")

	for _, x := range resp.NetworkInterfaces {

		logger.WithFields(log.Fields{
			"eni-id":     aws.StringValue(x.NetworkInterfaceId),
			"status":     aws.StringValue(x.Status),
			"eni-vpc-id": aws.StringValue(x.VpcId),
		}).Debug("found the lingering network interface")

		if vpcid != aws.StringValue(x.VpcId) {
			continue
		}
		if aws.StringValue(x.Status) != strings.ToLower(ec2.StateAvailable) {
			continue
		}

		_, err := client.svc.DeleteNetworkInterfaceWithContext(ctx, &ec2.DeleteNetworkInterfaceInput{
			NetworkInterfaceId: x.NetworkInterfaceId,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *VPCClient) getAZs(limit int) ([]string, error) {
	res, err := c.svc.DescribeAvailabilityZones(&ec2.DescribeAvailabilityZonesInput{})
	if err != nil {
		return nil, fmt.Errorf("failed to get availability zones: %w", err)
	}

	var azs []string
	for _, az := range res.AvailabilityZones {
		azs = append(azs, *az.ZoneId)
	}

	sort.Strings(azs)

	return azs[0:limit], nil

}
