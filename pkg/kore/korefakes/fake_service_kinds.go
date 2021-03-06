// Code generated by counterfeiter. DO NOT EDIT.
package korefakes

import (
	"context"
	"sync"

	v1 "github.com/appvia/kore/pkg/apis/services/v1"
	"github.com/appvia/kore/pkg/kore"
)

type FakeServiceKinds struct {
	CheckDeleteStub        func(context.Context, *v1.ServiceKind, ...kore.DeleteOptionFunc) error
	checkDeleteMutex       sync.RWMutex
	checkDeleteArgsForCall []struct {
		arg1 context.Context
		arg2 *v1.ServiceKind
		arg3 []kore.DeleteOptionFunc
	}
	checkDeleteReturns struct {
		result1 error
	}
	checkDeleteReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteStub        func(context.Context, string, ...kore.DeleteOptionFunc) (*v1.ServiceKind, error)
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 []kore.DeleteOptionFunc
	}
	deleteReturns struct {
		result1 *v1.ServiceKind
		result2 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 *v1.ServiceKind
		result2 error
	}
	GetStub        func(context.Context, string) (*v1.ServiceKind, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getReturns struct {
		result1 *v1.ServiceKind
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 *v1.ServiceKind
		result2 error
	}
	HasStub        func(context.Context, string) (bool, error)
	hasMutex       sync.RWMutex
	hasArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	hasReturns struct {
		result1 bool
		result2 error
	}
	hasReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	ListStub        func(context.Context, ...func(v1.ServiceKind) bool) (*v1.ServiceKindList, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		arg1 context.Context
		arg2 []func(v1.ServiceKind) bool
	}
	listReturns struct {
		result1 *v1.ServiceKindList
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 *v1.ServiceKindList
		result2 error
	}
	UpdateStub        func(context.Context, *v1.ServiceKind) error
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 context.Context
		arg2 *v1.ServiceKind
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

func (fake *FakeServiceKinds) CheckDelete(arg1 context.Context, arg2 *v1.ServiceKind, arg3 ...kore.DeleteOptionFunc) error {
	fake.checkDeleteMutex.Lock()
	ret, specificReturn := fake.checkDeleteReturnsOnCall[len(fake.checkDeleteArgsForCall)]
	fake.checkDeleteArgsForCall = append(fake.checkDeleteArgsForCall, struct {
		arg1 context.Context
		arg2 *v1.ServiceKind
		arg3 []kore.DeleteOptionFunc
	}{arg1, arg2, arg3})
	fake.recordInvocation("CheckDelete", []interface{}{arg1, arg2, arg3})
	fake.checkDeleteMutex.Unlock()
	if fake.CheckDeleteStub != nil {
		return fake.CheckDeleteStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.checkDeleteReturns
	return fakeReturns.result1
}

func (fake *FakeServiceKinds) CheckDeleteCallCount() int {
	fake.checkDeleteMutex.RLock()
	defer fake.checkDeleteMutex.RUnlock()
	return len(fake.checkDeleteArgsForCall)
}

func (fake *FakeServiceKinds) CheckDeleteCalls(stub func(context.Context, *v1.ServiceKind, ...kore.DeleteOptionFunc) error) {
	fake.checkDeleteMutex.Lock()
	defer fake.checkDeleteMutex.Unlock()
	fake.CheckDeleteStub = stub
}

func (fake *FakeServiceKinds) CheckDeleteArgsForCall(i int) (context.Context, *v1.ServiceKind, []kore.DeleteOptionFunc) {
	fake.checkDeleteMutex.RLock()
	defer fake.checkDeleteMutex.RUnlock()
	argsForCall := fake.checkDeleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeServiceKinds) CheckDeleteReturns(result1 error) {
	fake.checkDeleteMutex.Lock()
	defer fake.checkDeleteMutex.Unlock()
	fake.CheckDeleteStub = nil
	fake.checkDeleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceKinds) CheckDeleteReturnsOnCall(i int, result1 error) {
	fake.checkDeleteMutex.Lock()
	defer fake.checkDeleteMutex.Unlock()
	fake.CheckDeleteStub = nil
	if fake.checkDeleteReturnsOnCall == nil {
		fake.checkDeleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.checkDeleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceKinds) Delete(arg1 context.Context, arg2 string, arg3 ...kore.DeleteOptionFunc) (*v1.ServiceKind, error) {
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

func (fake *FakeServiceKinds) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeServiceKinds) DeleteCalls(stub func(context.Context, string, ...kore.DeleteOptionFunc) (*v1.ServiceKind, error)) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeServiceKinds) DeleteArgsForCall(i int) (context.Context, string, []kore.DeleteOptionFunc) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeServiceKinds) DeleteReturns(result1 *v1.ServiceKind, result2 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 *v1.ServiceKind
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceKinds) DeleteReturnsOnCall(i int, result1 *v1.ServiceKind, result2 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 *v1.ServiceKind
			result2 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 *v1.ServiceKind
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceKinds) Get(arg1 context.Context, arg2 string) (*v1.ServiceKind, error) {
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

func (fake *FakeServiceKinds) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeServiceKinds) GetCalls(stub func(context.Context, string) (*v1.ServiceKind, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeServiceKinds) GetArgsForCall(i int) (context.Context, string) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeServiceKinds) GetReturns(result1 *v1.ServiceKind, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 *v1.ServiceKind
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceKinds) GetReturnsOnCall(i int, result1 *v1.ServiceKind, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 *v1.ServiceKind
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 *v1.ServiceKind
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceKinds) Has(arg1 context.Context, arg2 string) (bool, error) {
	fake.hasMutex.Lock()
	ret, specificReturn := fake.hasReturnsOnCall[len(fake.hasArgsForCall)]
	fake.hasArgsForCall = append(fake.hasArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Has", []interface{}{arg1, arg2})
	fake.hasMutex.Unlock()
	if fake.HasStub != nil {
		return fake.HasStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.hasReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServiceKinds) HasCallCount() int {
	fake.hasMutex.RLock()
	defer fake.hasMutex.RUnlock()
	return len(fake.hasArgsForCall)
}

func (fake *FakeServiceKinds) HasCalls(stub func(context.Context, string) (bool, error)) {
	fake.hasMutex.Lock()
	defer fake.hasMutex.Unlock()
	fake.HasStub = stub
}

func (fake *FakeServiceKinds) HasArgsForCall(i int) (context.Context, string) {
	fake.hasMutex.RLock()
	defer fake.hasMutex.RUnlock()
	argsForCall := fake.hasArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeServiceKinds) HasReturns(result1 bool, result2 error) {
	fake.hasMutex.Lock()
	defer fake.hasMutex.Unlock()
	fake.HasStub = nil
	fake.hasReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceKinds) HasReturnsOnCall(i int, result1 bool, result2 error) {
	fake.hasMutex.Lock()
	defer fake.hasMutex.Unlock()
	fake.HasStub = nil
	if fake.hasReturnsOnCall == nil {
		fake.hasReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.hasReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceKinds) List(arg1 context.Context, arg2 ...func(v1.ServiceKind) bool) (*v1.ServiceKindList, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		arg1 context.Context
		arg2 []func(v1.ServiceKind) bool
	}{arg1, arg2})
	fake.recordInvocation("List", []interface{}{arg1, arg2})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServiceKinds) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeServiceKinds) ListCalls(stub func(context.Context, ...func(v1.ServiceKind) bool) (*v1.ServiceKindList, error)) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *FakeServiceKinds) ListArgsForCall(i int) (context.Context, []func(v1.ServiceKind) bool) {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	argsForCall := fake.listArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeServiceKinds) ListReturns(result1 *v1.ServiceKindList, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 *v1.ServiceKindList
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceKinds) ListReturnsOnCall(i int, result1 *v1.ServiceKindList, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 *v1.ServiceKindList
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 *v1.ServiceKindList
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceKinds) Update(arg1 context.Context, arg2 *v1.ServiceKind) error {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 context.Context
		arg2 *v1.ServiceKind
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

func (fake *FakeServiceKinds) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeServiceKinds) UpdateCalls(stub func(context.Context, *v1.ServiceKind) error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = stub
}

func (fake *FakeServiceKinds) UpdateArgsForCall(i int) (context.Context, *v1.ServiceKind) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	argsForCall := fake.updateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeServiceKinds) UpdateReturns(result1 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceKinds) UpdateReturnsOnCall(i int, result1 error) {
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

func (fake *FakeServiceKinds) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkDeleteMutex.RLock()
	defer fake.checkDeleteMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.hasMutex.RLock()
	defer fake.hasMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeServiceKinds) recordInvocation(key string, args []interface{}) {
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

var _ kore.ServiceKinds = new(FakeServiceKinds)
