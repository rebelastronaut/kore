// Code generated by counterfeiter. DO NOT EDIT.
package korefakes

import (
	"sync"

	"github.com/appvia/kore/pkg/kore"
)

type FakeTeam struct {
	AllocationsStub        func() kore.Allocations
	allocationsMutex       sync.RWMutex
	allocationsArgsForCall []struct {
	}
	allocationsReturns struct {
		result1 kore.Allocations
	}
	allocationsReturnsOnCall map[int]struct {
		result1 kore.Allocations
	}
	CloudStub        func() kore.Cloud
	cloudMutex       sync.RWMutex
	cloudArgsForCall []struct {
	}
	cloudReturns struct {
		result1 kore.Cloud
	}
	cloudReturnsOnCall map[int]struct {
		result1 kore.Cloud
	}
	ClustersStub        func() kore.Clusters
	clustersMutex       sync.RWMutex
	clustersArgsForCall []struct {
	}
	clustersReturns struct {
		result1 kore.Clusters
	}
	clustersReturnsOnCall map[int]struct {
		result1 kore.Clusters
	}
	KubernetesStub        func() kore.Kubernetes
	kubernetesMutex       sync.RWMutex
	kubernetesArgsForCall []struct {
	}
	kubernetesReturns struct {
		result1 kore.Kubernetes
	}
	kubernetesReturnsOnCall map[int]struct {
		result1 kore.Kubernetes
	}
	MembersStub        func() kore.TeamMembers
	membersMutex       sync.RWMutex
	membersArgsForCall []struct {
	}
	membersReturns struct {
		result1 kore.TeamMembers
	}
	membersReturnsOnCall map[int]struct {
		result1 kore.TeamMembers
	}
	NamespaceClaimsStub        func() kore.NamespaceClaims
	namespaceClaimsMutex       sync.RWMutex
	namespaceClaimsArgsForCall []struct {
	}
	namespaceClaimsReturns struct {
		result1 kore.NamespaceClaims
	}
	namespaceClaimsReturnsOnCall map[int]struct {
		result1 kore.NamespaceClaims
	}
	SecretsStub        func() kore.Secrets
	secretsMutex       sync.RWMutex
	secretsArgsForCall []struct {
	}
	secretsReturns struct {
		result1 kore.Secrets
	}
	secretsReturnsOnCall map[int]struct {
		result1 kore.Secrets
	}
	ServiceCredentialsStub        func() kore.ServiceCredentials
	serviceCredentialsMutex       sync.RWMutex
	serviceCredentialsArgsForCall []struct {
	}
	serviceCredentialsReturns struct {
		result1 kore.ServiceCredentials
	}
	serviceCredentialsReturnsOnCall map[int]struct {
		result1 kore.ServiceCredentials
	}
	ServicesStub        func() kore.Services
	servicesMutex       sync.RWMutex
	servicesArgsForCall []struct {
	}
	servicesReturns struct {
		result1 kore.Services
	}
	servicesReturnsOnCall map[int]struct {
		result1 kore.Services
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTeam) Allocations() kore.Allocations {
	fake.allocationsMutex.Lock()
	ret, specificReturn := fake.allocationsReturnsOnCall[len(fake.allocationsArgsForCall)]
	fake.allocationsArgsForCall = append(fake.allocationsArgsForCall, struct {
	}{})
	fake.recordInvocation("Allocations", []interface{}{})
	fake.allocationsMutex.Unlock()
	if fake.AllocationsStub != nil {
		return fake.AllocationsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.allocationsReturns
	return fakeReturns.result1
}

func (fake *FakeTeam) AllocationsCallCount() int {
	fake.allocationsMutex.RLock()
	defer fake.allocationsMutex.RUnlock()
	return len(fake.allocationsArgsForCall)
}

func (fake *FakeTeam) AllocationsCalls(stub func() kore.Allocations) {
	fake.allocationsMutex.Lock()
	defer fake.allocationsMutex.Unlock()
	fake.AllocationsStub = stub
}

func (fake *FakeTeam) AllocationsReturns(result1 kore.Allocations) {
	fake.allocationsMutex.Lock()
	defer fake.allocationsMutex.Unlock()
	fake.AllocationsStub = nil
	fake.allocationsReturns = struct {
		result1 kore.Allocations
	}{result1}
}

func (fake *FakeTeam) AllocationsReturnsOnCall(i int, result1 kore.Allocations) {
	fake.allocationsMutex.Lock()
	defer fake.allocationsMutex.Unlock()
	fake.AllocationsStub = nil
	if fake.allocationsReturnsOnCall == nil {
		fake.allocationsReturnsOnCall = make(map[int]struct {
			result1 kore.Allocations
		})
	}
	fake.allocationsReturnsOnCall[i] = struct {
		result1 kore.Allocations
	}{result1}
}

func (fake *FakeTeam) Cloud() kore.Cloud {
	fake.cloudMutex.Lock()
	ret, specificReturn := fake.cloudReturnsOnCall[len(fake.cloudArgsForCall)]
	fake.cloudArgsForCall = append(fake.cloudArgsForCall, struct {
	}{})
	fake.recordInvocation("Cloud", []interface{}{})
	fake.cloudMutex.Unlock()
	if fake.CloudStub != nil {
		return fake.CloudStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.cloudReturns
	return fakeReturns.result1
}

func (fake *FakeTeam) CloudCallCount() int {
	fake.cloudMutex.RLock()
	defer fake.cloudMutex.RUnlock()
	return len(fake.cloudArgsForCall)
}

func (fake *FakeTeam) CloudCalls(stub func() kore.Cloud) {
	fake.cloudMutex.Lock()
	defer fake.cloudMutex.Unlock()
	fake.CloudStub = stub
}

func (fake *FakeTeam) CloudReturns(result1 kore.Cloud) {
	fake.cloudMutex.Lock()
	defer fake.cloudMutex.Unlock()
	fake.CloudStub = nil
	fake.cloudReturns = struct {
		result1 kore.Cloud
	}{result1}
}

func (fake *FakeTeam) CloudReturnsOnCall(i int, result1 kore.Cloud) {
	fake.cloudMutex.Lock()
	defer fake.cloudMutex.Unlock()
	fake.CloudStub = nil
	if fake.cloudReturnsOnCall == nil {
		fake.cloudReturnsOnCall = make(map[int]struct {
			result1 kore.Cloud
		})
	}
	fake.cloudReturnsOnCall[i] = struct {
		result1 kore.Cloud
	}{result1}
}

func (fake *FakeTeam) Clusters() kore.Clusters {
	fake.clustersMutex.Lock()
	ret, specificReturn := fake.clustersReturnsOnCall[len(fake.clustersArgsForCall)]
	fake.clustersArgsForCall = append(fake.clustersArgsForCall, struct {
	}{})
	fake.recordInvocation("Clusters", []interface{}{})
	fake.clustersMutex.Unlock()
	if fake.ClustersStub != nil {
		return fake.ClustersStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.clustersReturns
	return fakeReturns.result1
}

func (fake *FakeTeam) ClustersCallCount() int {
	fake.clustersMutex.RLock()
	defer fake.clustersMutex.RUnlock()
	return len(fake.clustersArgsForCall)
}

func (fake *FakeTeam) ClustersCalls(stub func() kore.Clusters) {
	fake.clustersMutex.Lock()
	defer fake.clustersMutex.Unlock()
	fake.ClustersStub = stub
}

func (fake *FakeTeam) ClustersReturns(result1 kore.Clusters) {
	fake.clustersMutex.Lock()
	defer fake.clustersMutex.Unlock()
	fake.ClustersStub = nil
	fake.clustersReturns = struct {
		result1 kore.Clusters
	}{result1}
}

func (fake *FakeTeam) ClustersReturnsOnCall(i int, result1 kore.Clusters) {
	fake.clustersMutex.Lock()
	defer fake.clustersMutex.Unlock()
	fake.ClustersStub = nil
	if fake.clustersReturnsOnCall == nil {
		fake.clustersReturnsOnCall = make(map[int]struct {
			result1 kore.Clusters
		})
	}
	fake.clustersReturnsOnCall[i] = struct {
		result1 kore.Clusters
	}{result1}
}

func (fake *FakeTeam) Kubernetes() kore.Kubernetes {
	fake.kubernetesMutex.Lock()
	ret, specificReturn := fake.kubernetesReturnsOnCall[len(fake.kubernetesArgsForCall)]
	fake.kubernetesArgsForCall = append(fake.kubernetesArgsForCall, struct {
	}{})
	fake.recordInvocation("Kubernetes", []interface{}{})
	fake.kubernetesMutex.Unlock()
	if fake.KubernetesStub != nil {
		return fake.KubernetesStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.kubernetesReturns
	return fakeReturns.result1
}

func (fake *FakeTeam) KubernetesCallCount() int {
	fake.kubernetesMutex.RLock()
	defer fake.kubernetesMutex.RUnlock()
	return len(fake.kubernetesArgsForCall)
}

func (fake *FakeTeam) KubernetesCalls(stub func() kore.Kubernetes) {
	fake.kubernetesMutex.Lock()
	defer fake.kubernetesMutex.Unlock()
	fake.KubernetesStub = stub
}

func (fake *FakeTeam) KubernetesReturns(result1 kore.Kubernetes) {
	fake.kubernetesMutex.Lock()
	defer fake.kubernetesMutex.Unlock()
	fake.KubernetesStub = nil
	fake.kubernetesReturns = struct {
		result1 kore.Kubernetes
	}{result1}
}

func (fake *FakeTeam) KubernetesReturnsOnCall(i int, result1 kore.Kubernetes) {
	fake.kubernetesMutex.Lock()
	defer fake.kubernetesMutex.Unlock()
	fake.KubernetesStub = nil
	if fake.kubernetesReturnsOnCall == nil {
		fake.kubernetesReturnsOnCall = make(map[int]struct {
			result1 kore.Kubernetes
		})
	}
	fake.kubernetesReturnsOnCall[i] = struct {
		result1 kore.Kubernetes
	}{result1}
}

func (fake *FakeTeam) Members() kore.TeamMembers {
	fake.membersMutex.Lock()
	ret, specificReturn := fake.membersReturnsOnCall[len(fake.membersArgsForCall)]
	fake.membersArgsForCall = append(fake.membersArgsForCall, struct {
	}{})
	fake.recordInvocation("Members", []interface{}{})
	fake.membersMutex.Unlock()
	if fake.MembersStub != nil {
		return fake.MembersStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.membersReturns
	return fakeReturns.result1
}

func (fake *FakeTeam) MembersCallCount() int {
	fake.membersMutex.RLock()
	defer fake.membersMutex.RUnlock()
	return len(fake.membersArgsForCall)
}

func (fake *FakeTeam) MembersCalls(stub func() kore.TeamMembers) {
	fake.membersMutex.Lock()
	defer fake.membersMutex.Unlock()
	fake.MembersStub = stub
}

func (fake *FakeTeam) MembersReturns(result1 kore.TeamMembers) {
	fake.membersMutex.Lock()
	defer fake.membersMutex.Unlock()
	fake.MembersStub = nil
	fake.membersReturns = struct {
		result1 kore.TeamMembers
	}{result1}
}

func (fake *FakeTeam) MembersReturnsOnCall(i int, result1 kore.TeamMembers) {
	fake.membersMutex.Lock()
	defer fake.membersMutex.Unlock()
	fake.MembersStub = nil
	if fake.membersReturnsOnCall == nil {
		fake.membersReturnsOnCall = make(map[int]struct {
			result1 kore.TeamMembers
		})
	}
	fake.membersReturnsOnCall[i] = struct {
		result1 kore.TeamMembers
	}{result1}
}

func (fake *FakeTeam) NamespaceClaims() kore.NamespaceClaims {
	fake.namespaceClaimsMutex.Lock()
	ret, specificReturn := fake.namespaceClaimsReturnsOnCall[len(fake.namespaceClaimsArgsForCall)]
	fake.namespaceClaimsArgsForCall = append(fake.namespaceClaimsArgsForCall, struct {
	}{})
	fake.recordInvocation("NamespaceClaims", []interface{}{})
	fake.namespaceClaimsMutex.Unlock()
	if fake.NamespaceClaimsStub != nil {
		return fake.NamespaceClaimsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.namespaceClaimsReturns
	return fakeReturns.result1
}

func (fake *FakeTeam) NamespaceClaimsCallCount() int {
	fake.namespaceClaimsMutex.RLock()
	defer fake.namespaceClaimsMutex.RUnlock()
	return len(fake.namespaceClaimsArgsForCall)
}

func (fake *FakeTeam) NamespaceClaimsCalls(stub func() kore.NamespaceClaims) {
	fake.namespaceClaimsMutex.Lock()
	defer fake.namespaceClaimsMutex.Unlock()
	fake.NamespaceClaimsStub = stub
}

func (fake *FakeTeam) NamespaceClaimsReturns(result1 kore.NamespaceClaims) {
	fake.namespaceClaimsMutex.Lock()
	defer fake.namespaceClaimsMutex.Unlock()
	fake.NamespaceClaimsStub = nil
	fake.namespaceClaimsReturns = struct {
		result1 kore.NamespaceClaims
	}{result1}
}

func (fake *FakeTeam) NamespaceClaimsReturnsOnCall(i int, result1 kore.NamespaceClaims) {
	fake.namespaceClaimsMutex.Lock()
	defer fake.namespaceClaimsMutex.Unlock()
	fake.NamespaceClaimsStub = nil
	if fake.namespaceClaimsReturnsOnCall == nil {
		fake.namespaceClaimsReturnsOnCall = make(map[int]struct {
			result1 kore.NamespaceClaims
		})
	}
	fake.namespaceClaimsReturnsOnCall[i] = struct {
		result1 kore.NamespaceClaims
	}{result1}
}

func (fake *FakeTeam) Secrets() kore.Secrets {
	fake.secretsMutex.Lock()
	ret, specificReturn := fake.secretsReturnsOnCall[len(fake.secretsArgsForCall)]
	fake.secretsArgsForCall = append(fake.secretsArgsForCall, struct {
	}{})
	fake.recordInvocation("Secrets", []interface{}{})
	fake.secretsMutex.Unlock()
	if fake.SecretsStub != nil {
		return fake.SecretsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.secretsReturns
	return fakeReturns.result1
}

func (fake *FakeTeam) SecretsCallCount() int {
	fake.secretsMutex.RLock()
	defer fake.secretsMutex.RUnlock()
	return len(fake.secretsArgsForCall)
}

func (fake *FakeTeam) SecretsCalls(stub func() kore.Secrets) {
	fake.secretsMutex.Lock()
	defer fake.secretsMutex.Unlock()
	fake.SecretsStub = stub
}

func (fake *FakeTeam) SecretsReturns(result1 kore.Secrets) {
	fake.secretsMutex.Lock()
	defer fake.secretsMutex.Unlock()
	fake.SecretsStub = nil
	fake.secretsReturns = struct {
		result1 kore.Secrets
	}{result1}
}

func (fake *FakeTeam) SecretsReturnsOnCall(i int, result1 kore.Secrets) {
	fake.secretsMutex.Lock()
	defer fake.secretsMutex.Unlock()
	fake.SecretsStub = nil
	if fake.secretsReturnsOnCall == nil {
		fake.secretsReturnsOnCall = make(map[int]struct {
			result1 kore.Secrets
		})
	}
	fake.secretsReturnsOnCall[i] = struct {
		result1 kore.Secrets
	}{result1}
}

func (fake *FakeTeam) ServiceCredentials() kore.ServiceCredentials {
	fake.serviceCredentialsMutex.Lock()
	ret, specificReturn := fake.serviceCredentialsReturnsOnCall[len(fake.serviceCredentialsArgsForCall)]
	fake.serviceCredentialsArgsForCall = append(fake.serviceCredentialsArgsForCall, struct {
	}{})
	fake.recordInvocation("ServiceCredentials", []interface{}{})
	fake.serviceCredentialsMutex.Unlock()
	if fake.ServiceCredentialsStub != nil {
		return fake.ServiceCredentialsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.serviceCredentialsReturns
	return fakeReturns.result1
}

func (fake *FakeTeam) ServiceCredentialsCallCount() int {
	fake.serviceCredentialsMutex.RLock()
	defer fake.serviceCredentialsMutex.RUnlock()
	return len(fake.serviceCredentialsArgsForCall)
}

func (fake *FakeTeam) ServiceCredentialsCalls(stub func() kore.ServiceCredentials) {
	fake.serviceCredentialsMutex.Lock()
	defer fake.serviceCredentialsMutex.Unlock()
	fake.ServiceCredentialsStub = stub
}

func (fake *FakeTeam) ServiceCredentialsReturns(result1 kore.ServiceCredentials) {
	fake.serviceCredentialsMutex.Lock()
	defer fake.serviceCredentialsMutex.Unlock()
	fake.ServiceCredentialsStub = nil
	fake.serviceCredentialsReturns = struct {
		result1 kore.ServiceCredentials
	}{result1}
}

func (fake *FakeTeam) ServiceCredentialsReturnsOnCall(i int, result1 kore.ServiceCredentials) {
	fake.serviceCredentialsMutex.Lock()
	defer fake.serviceCredentialsMutex.Unlock()
	fake.ServiceCredentialsStub = nil
	if fake.serviceCredentialsReturnsOnCall == nil {
		fake.serviceCredentialsReturnsOnCall = make(map[int]struct {
			result1 kore.ServiceCredentials
		})
	}
	fake.serviceCredentialsReturnsOnCall[i] = struct {
		result1 kore.ServiceCredentials
	}{result1}
}

func (fake *FakeTeam) Services() kore.Services {
	fake.servicesMutex.Lock()
	ret, specificReturn := fake.servicesReturnsOnCall[len(fake.servicesArgsForCall)]
	fake.servicesArgsForCall = append(fake.servicesArgsForCall, struct {
	}{})
	fake.recordInvocation("Services", []interface{}{})
	fake.servicesMutex.Unlock()
	if fake.ServicesStub != nil {
		return fake.ServicesStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.servicesReturns
	return fakeReturns.result1
}

func (fake *FakeTeam) ServicesCallCount() int {
	fake.servicesMutex.RLock()
	defer fake.servicesMutex.RUnlock()
	return len(fake.servicesArgsForCall)
}

func (fake *FakeTeam) ServicesCalls(stub func() kore.Services) {
	fake.servicesMutex.Lock()
	defer fake.servicesMutex.Unlock()
	fake.ServicesStub = stub
}

func (fake *FakeTeam) ServicesReturns(result1 kore.Services) {
	fake.servicesMutex.Lock()
	defer fake.servicesMutex.Unlock()
	fake.ServicesStub = nil
	fake.servicesReturns = struct {
		result1 kore.Services
	}{result1}
}

func (fake *FakeTeam) ServicesReturnsOnCall(i int, result1 kore.Services) {
	fake.servicesMutex.Lock()
	defer fake.servicesMutex.Unlock()
	fake.ServicesStub = nil
	if fake.servicesReturnsOnCall == nil {
		fake.servicesReturnsOnCall = make(map[int]struct {
			result1 kore.Services
		})
	}
	fake.servicesReturnsOnCall[i] = struct {
		result1 kore.Services
	}{result1}
}

func (fake *FakeTeam) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.allocationsMutex.RLock()
	defer fake.allocationsMutex.RUnlock()
	fake.cloudMutex.RLock()
	defer fake.cloudMutex.RUnlock()
	fake.clustersMutex.RLock()
	defer fake.clustersMutex.RUnlock()
	fake.kubernetesMutex.RLock()
	defer fake.kubernetesMutex.RUnlock()
	fake.membersMutex.RLock()
	defer fake.membersMutex.RUnlock()
	fake.namespaceClaimsMutex.RLock()
	defer fake.namespaceClaimsMutex.RUnlock()
	fake.secretsMutex.RLock()
	defer fake.secretsMutex.RUnlock()
	fake.serviceCredentialsMutex.RLock()
	defer fake.serviceCredentialsMutex.RUnlock()
	fake.servicesMutex.RLock()
	defer fake.servicesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTeam) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ kore.Team = new(FakeTeam)
