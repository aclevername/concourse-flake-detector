// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"sync"

	"github.com/aclevername/concourse-flake-detector/concourse"
)

type FakeGetter struct {
	Stub        func(string) ([]byte, error)
	mutex       sync.RWMutex
	argsForCall []struct {
		arg1 string
	}
	returns struct {
		result1 []byte
		result2 error
	}
	returnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGetter) Spy(arg1 string) ([]byte, error) {
	fake.mutex.Lock()
	ret, specificReturn := fake.returnsOnCall[len(fake.argsForCall)]
	fake.argsForCall = append(fake.argsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Getter", []interface{}{arg1})
	fake.mutex.Unlock()
	if fake.Stub != nil {
		return fake.Stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.returns.result1, fake.returns.result2
}

func (fake *FakeGetter) CallCount() int {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return len(fake.argsForCall)
}

func (fake *FakeGetter) ArgsForCall(i int) string {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return fake.argsForCall[i].arg1
}

func (fake *FakeGetter) Returns(result1 []byte, result2 error) {
	fake.Stub = nil
	fake.returns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeGetter) ReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.Stub = nil
	if fake.returnsOnCall == nil {
		fake.returnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.returnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeGetter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return fake.invocations
}

func (fake *FakeGetter) recordInvocation(key string, args []interface{}) {
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

var _ concourse.Getter = new(FakeGetter).Spy