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

package metadata

import (
	"fmt"
	"time"

	costsv1 "github.com/appvia/kore/pkg/apis/costs/v1beta1"
)

// Metadata allows access to cloud service metadata such as instance types and prices
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . Metadata
type Metadata interface {
	// Clouds retrieves the list of supported clouds
	Clouds() ([]string, error)
	// MapProviderToCloud maps a provider to a cloud
	MapProviderToCloud(provider string) (string, error)
	// Regions retrieves the list of available regions, organised by continent, for the specified cloud
	Regions(cloud string) (*costsv1.ContinentList, error)
	// RegionZones retrieves the list of available AZs in the given region, for the specified cloud
	RegionZones(cloud string, region string) ([]string, error)
	// InstanceTypes retrieves the list of available instance types for the specified cloud and region
	InstanceTypes(cloud string, region string) (*costsv1.InstanceTypeList, error)
	// InstanceType gets the metadata for a specific selected instance type for the specified cloud and region
	InstanceType(cloud string, region string, instanceType string) (*costsv1.InstanceType, error)
	// KubernetesVersions retrieves the list of supported kubernetes versions for the specified cloud and region
	KubernetesVersions(cloud string, region string) ([]string, error)
	// KubernetesControlPlanCost retrieves the price in microdollars per hour of a Kubernetes
	// control plane in the specific cloud and region
	KubernetesControlPlaneCost(cloud string, region string) (int64, error)
	// KubernetesExposedServiceCost retrieves the price in microdollars per hour of an exposed service (i.e.
	// HTTP load balancer)
	KubernetesExposedServiceCost(cloud string, region string) (int64, error)
	// PricesAvailable identifies if the metadata service can supply pricing info currently
	PricesAvailable() bool
}

// New creates a new instance of the metadata API
func New(config *Config) Metadata {
	cloudinfo := NewCloudInfo(config.CloudinfoURL)
	return &metadataImpl{
		cloudinfo:          cloudinfo,
		cloudinfoAvailable: false,
	}
}

type metadataImpl struct {
	cloudinfo          Cloudinfo
	cloudinfoAvailable bool
	cloudinfoCheckDue  time.Time
}

func (m *metadataImpl) info() Cloudinfo {
	// Check every 30 seconds if cloudinfo is available.
	if m.cloudinfoCheckDue.IsZero() || m.cloudinfoCheckDue.Before(time.Now()) {
		m.cloudinfoAvailable = m.cloudinfo.Ready()
		m.cloudinfoCheckDue = time.Now().Add(30 * time.Second)
	}
	if m.cloudinfoAvailable {
		return m.cloudinfo
	}
	return &staticData{}
}

func (m *metadataImpl) PricesAvailable() bool {
	_ = m.info()
	return m.cloudinfoAvailable
}

func (m *metadataImpl) Clouds() ([]string, error) {
	return []string{CloudGCP, CloudAWS, CloudAzure}, nil
}

func (m *metadataImpl) MapProviderToCloud(provider string) (string, error) {
	cloud := getCloudForClusterProvider(provider)
	if cloud == "" {
		return "", fmt.Errorf("unknown Kubernetes provider %s, cannot determine cloud provider", provider)
	}
	return cloud, nil
}

func (m *metadataImpl) Regions(cloud string) (*costsv1.ContinentList, error) {
	continents, err := m.info().KubernetesRegions(cloud)
	if err != nil {
		return nil, err
	}
	if continents == nil {
		return nil, nil
	}
	result := &costsv1.ContinentList{}
	result.Items = append(result.Items, continents...)
	return result, nil
}

func (m *metadataImpl) RegionZones(cloud string, region string) ([]string, error) {
	return m.info().KubernetesRegionAZs(cloud, region)
}

func (m *metadataImpl) InstanceTypes(cloud string, region string) (*costsv1.InstanceTypeList, error) {
	instanceTypes, err := m.info().KubernetesInstanceTypes(cloud, region)
	if err != nil {
		return nil, err
	}
	if instanceTypes == nil {
		return nil, nil
	}
	result := &costsv1.InstanceTypeList{}
	result.Items = append(result.Items, instanceTypes...)
	return result, nil
}

func (m *metadataImpl) InstanceType(cloud string, region string, instanceType string) (*costsv1.InstanceType, error) {
	return m.info().KubernetesInstanceType(cloud, region, instanceType)
}

func (m *metadataImpl) KubernetesVersions(cloud string, region string) ([]string, error) {
	return m.info().KubernetesVersions(cloud, region)
}

func (m *metadataImpl) KubernetesControlPlaneCost(cloud string, region string) (int64, error) {
	// @TODO: Determine this from the providers. For now, AWS and GCP both charge 10c/hr worldwide, Azure 0c.
	switch cloud {
	case CloudGCP:
		return 100000, nil
	case CloudAWS:
		return 100000, nil
	}
	return 0, nil
}

func (m *metadataImpl) KubernetesExposedServiceCost(cloud string, region string) (int64, error) {
	// @TODO: Determine this from the providers. For now, just return a typicalish load
	// balancer cost of ~2.5c/hr
	return 25000, nil
}
