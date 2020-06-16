// Code generated by counterfeiter. DO NOT EDIT.
package korefakes

import (
	"context"
	"sync"

	v1 "github.com/appvia/kore/pkg/apis/services/v1"
	"github.com/appvia/kore/pkg/kore"
)

type FakeServices struct {
	DeleteStub        func(context.Context, string, ...kore.DeleteOptionFunc) (*v1.Service, error)
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 []kore.DeleteOptionFunc
	}
	deleteReturns struct {
		result1 *v1.Service
		result2 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 *v1.Service
		result2 error
	}
	GetStub        func(context.Context, string) (*v1.Service, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getReturns struct {
		result1 *v1.Service
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 *v1.Service
		result2 error
	}
	ListStub        func(context.Context) (*v1.ServiceList, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		arg1 context.Context
	}
	listReturns struct {
		result1 *v1.ServiceList
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 *v1.ServiceList
		result2 error
	}
	ListFilteredStub        func(context.Context, func(v1.Service) bool) (*v1.ServiceList, error)
	listFilteredMutex       sync.RWMutex
	listFilteredArgsForCall []struct {
		arg1 context.Context
		arg2 func(v1.Service) bool
	}
	listFilteredReturns struct {
		result1 *v1.ServiceList
		result2 error
	}
	listFilteredReturnsOnCall map[int]struct {
		result1 *v1.ServiceList
		result2 error
	}
	UpdateStub        func(context.Context, *v1.Service) error
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 context.Context
		arg2 *v1.Service
	}
	updateReturns struct {
		result1 error
	}
	updateReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeServices) Delete(arg1 context.Context, arg2 string, arg3 ...kore.DeleteOptionFunc) (*v1.Service, error) {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 []kore.DeleteOptionFunc
	}{arg1, arg2, arg3})
	fake.recordInvocation("Delete", []interface{}{arg1, arg2, arg3})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deleteReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServices) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeServices) DeleteCalls(stub func(context.Context, string, ...kore.DeleteOptionFunc) (*v1.Service, error)) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeServices) DeleteArgsForCall(i int) (context.Context, string, []kore.DeleteOptionFunc) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeServices) DeleteReturns(result1 *v1.Service, result2 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 *v1.Service
		result2 error
	}{result1, result2}
}

func (fake *FakeServices) DeleteReturnsOnCall(i int, result1 *v1.Service, result2 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 *v1.Service
			result2 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 *v1.Service
		result2 error
	}{result1, result2}
}

func (fake *FakeServices) Get(arg1 context.Context, arg2 string) (*v1.Service, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Get", []interface{}{arg1, arg2})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServices) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeServices) GetCalls(stub func(context.Context, string) (*v1.Service, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeServices) GetArgsForCall(i int) (context.Context, string) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeServices) GetReturns(result1 *v1.Service, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 *v1.Service
		result2 error
	}{result1, result2}
}

func (fake *FakeServices) GetReturnsOnCall(i int, result1 *v1.Service, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 *v1.Service
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 *v1.Service
		result2 error
	}{result1, result2}
}

func (fake *FakeServices) List(arg1 context.Context) (*v1.ServiceList, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("List", []interface{}{arg1})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServices) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeServices) ListCalls(stub func(context.Context) (*v1.ServiceList, error)) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *FakeServices) ListArgsForCall(i int) context.Context {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	argsForCall := fake.listArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeServices) ListReturns(result1 *v1.ServiceList, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 *v1.ServiceList
		result2 error
	}{result1, result2}
}

func (fake *FakeServices) ListReturnsOnCall(i int, result1 *v1.ServiceList, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 *v1.ServiceList
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 *v1.ServiceList
		result2 error
	}{result1, result2}
}

func (fake *FakeServices) ListFiltered(arg1 context.Context, arg2 func(v1.Service) bool) (*v1.ServiceList, error) {
	fake.listFilteredMutex.Lock()
	ret, specificReturn := fake.listFilteredReturnsOnCall[len(fake.listFilteredArgsForCall)]
	fake.listFilteredArgsForCall = append(fake.listFilteredArgsForCall, struct {
		arg1 context.Context
		arg2 func(v1.Service) bool
	}{arg1, arg2})
	fake.recordInvocation("ListFiltered", []interface{}{arg1, arg2})
	fake.listFilteredMutex.Unlock()
	if fake.ListFilteredStub != nil {
		return fake.ListFilteredStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listFilteredReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServices) ListFilteredCallCount() int {
	fake.listFilteredMutex.RLock()
	defer fake.listFilteredMutex.RUnlock()
	return len(fake.listFilteredArgsForCall)
}

func (fake *FakeServices) ListFilteredCalls(stub func(context.Context, func(v1.Service) bool) (*v1.ServiceList, error)) {
	fake.listFilteredMutex.Lock()
	defer fake.listFilteredMutex.Unlock()
	fake.ListFilteredStub = stub
}

func (fake *FakeServices) ListFilteredArgsForCall(i int) (context.Context, func(v1.Service) bool) {
	fake.listFilteredMutex.RLock()
	defer fake.listFilteredMutex.RUnlock()
	argsForCall := fake.listFilteredArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeServices) ListFilteredReturns(result1 *v1.ServiceList, result2 error) {
	fake.listFilteredMutex.Lock()
	defer fake.listFilteredMutex.Unlock()
	fake.ListFilteredStub = nil
	fake.listFilteredReturns = struct {
		result1 *v1.ServiceList
		result2 error
	}{result1, result2}
}

func (fake *FakeServices) ListFilteredReturnsOnCall(i int, result1 *v1.ServiceList, result2 error) {
	fake.listFilteredMutex.Lock()
	defer fake.listFilteredMutex.Unlock()
	fake.ListFilteredStub = nil
	if fake.listFilteredReturnsOnCall == nil {
		fake.listFilteredReturnsOnCall = make(map[int]struct {
			result1 *v1.ServiceList
			result2 error
		})
	}
	fake.listFilteredReturnsOnCall[i] = struct {
		result1 *v1.ServiceList
		result2 error
	}{result1, result2}
}

func (fake *FakeServices) Update(arg1 context.Context, arg2 *v1.Service) error {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 context.Context
		arg2 *v1.Service
	}{arg1, arg2})
	fake.recordInvocation("Update", []interface{}{arg1, arg2})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.updateReturns
	return fakeReturns.result1
}

func (fake *FakeServices) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeServices) UpdateCalls(stub func(context.Context, *v1.Service) error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = stub
}

func (fake *FakeServices) UpdateArgsForCall(i int) (context.Context, *v1.Service) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	argsForCall := fake.updateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeServices) UpdateReturns(result1 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServices) UpdateReturnsOnCall(i int, result1 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeServices) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.listFilteredMutex.RLock()
	defer fake.listFilteredMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeServices) recordInvocation(key string, args []interface{}) {
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

var _ kore.Services = new(FakeServices)
