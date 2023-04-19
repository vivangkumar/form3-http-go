// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"net/http"
	"sync"
)

type FakeBaseClient struct {
	DeleteStub        func(context.Context, string, map[string]string) (*http.Response, error)
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 map[string]string
	}
	deleteReturns struct {
		result1 *http.Response
		result2 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 *http.Response
		result2 error
	}
	GetStub        func(context.Context, string, map[string]string, any) (*http.Response, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 map[string]string
		arg4 any
	}
	getReturns struct {
		result1 *http.Response
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 *http.Response
		result2 error
	}
	PostStub        func(context.Context, string, any, any) (*http.Response, error)
	postMutex       sync.RWMutex
	postArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 any
		arg4 any
	}
	postReturns struct {
		result1 *http.Response
		result2 error
	}
	postReturnsOnCall map[int]struct {
		result1 *http.Response
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBaseClient) Delete(arg1 context.Context, arg2 string, arg3 map[string]string) (*http.Response, error) {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 map[string]string
	}{arg1, arg2, arg3})
	stub := fake.DeleteStub
	fakeReturns := fake.deleteReturns
	fake.recordInvocation("Delete", []interface{}{arg1, arg2, arg3})
	fake.deleteMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBaseClient) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeBaseClient) DeleteCalls(stub func(context.Context, string, map[string]string) (*http.Response, error)) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeBaseClient) DeleteArgsForCall(i int) (context.Context, string, map[string]string) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeBaseClient) DeleteReturns(result1 *http.Response, result2 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeBaseClient) DeleteReturnsOnCall(i int, result1 *http.Response, result2 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 *http.Response
			result2 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeBaseClient) Get(arg1 context.Context, arg2 string, arg3 map[string]string, arg4 any) (*http.Response, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 map[string]string
		arg4 any
	}{arg1, arg2, arg3, arg4})
	stub := fake.GetStub
	fakeReturns := fake.getReturns
	fake.recordInvocation("Get", []interface{}{arg1, arg2, arg3, arg4})
	fake.getMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBaseClient) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeBaseClient) GetCalls(stub func(context.Context, string, map[string]string, any) (*http.Response, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeBaseClient) GetArgsForCall(i int) (context.Context, string, map[string]string, any) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeBaseClient) GetReturns(result1 *http.Response, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeBaseClient) GetReturnsOnCall(i int, result1 *http.Response, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 *http.Response
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeBaseClient) Post(arg1 context.Context, arg2 string, arg3 any, arg4 any) (*http.Response, error) {
	fake.postMutex.Lock()
	ret, specificReturn := fake.postReturnsOnCall[len(fake.postArgsForCall)]
	fake.postArgsForCall = append(fake.postArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 any
		arg4 any
	}{arg1, arg2, arg3, arg4})
	stub := fake.PostStub
	fakeReturns := fake.postReturns
	fake.recordInvocation("Post", []interface{}{arg1, arg2, arg3, arg4})
	fake.postMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBaseClient) PostCallCount() int {
	fake.postMutex.RLock()
	defer fake.postMutex.RUnlock()
	return len(fake.postArgsForCall)
}

func (fake *FakeBaseClient) PostCalls(stub func(context.Context, string, any, any) (*http.Response, error)) {
	fake.postMutex.Lock()
	defer fake.postMutex.Unlock()
	fake.PostStub = stub
}

func (fake *FakeBaseClient) PostArgsForCall(i int) (context.Context, string, any, any) {
	fake.postMutex.RLock()
	defer fake.postMutex.RUnlock()
	argsForCall := fake.postArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeBaseClient) PostReturns(result1 *http.Response, result2 error) {
	fake.postMutex.Lock()
	defer fake.postMutex.Unlock()
	fake.PostStub = nil
	fake.postReturns = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeBaseClient) PostReturnsOnCall(i int, result1 *http.Response, result2 error) {
	fake.postMutex.Lock()
	defer fake.postMutex.Unlock()
	fake.PostStub = nil
	if fake.postReturnsOnCall == nil {
		fake.postReturnsOnCall = make(map[int]struct {
			result1 *http.Response
			result2 error
		})
	}
	fake.postReturnsOnCall[i] = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeBaseClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.postMutex.RLock()
	defer fake.postMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBaseClient) recordInvocation(key string, args []interface{}) {
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