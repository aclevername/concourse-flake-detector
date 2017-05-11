// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"sync"

	"github.com/aclevername/concourse-flake-detector/concourse"
)

type FakeClientInterface struct {
	GetPipelineStub        func(string) (concourse.Pipeline, error)
	getPipelineMutex       sync.RWMutex
	getPipelineArgsForCall []struct {
		arg1 string
	}
	getPipelineReturns struct {
		result1 concourse.Pipeline
		result2 error
	}
	getPipelineReturnsOnCall map[int]struct {
		result1 concourse.Pipeline
		result2 error
	}
	GetBuildsStub        func(concourse.Job) ([]concourse.Build, error)
	getBuildsMutex       sync.RWMutex
	getBuildsArgsForCall []struct {
		arg1 concourse.Job
	}
	getBuildsReturns struct {
		result1 []concourse.Build
		result2 error
	}
	getBuildsReturnsOnCall map[int]struct {
		result1 []concourse.Build
		result2 error
	}
	GetResourcesStub        func(concourse.Build) (concourse.Run, error)
	getResourcesMutex       sync.RWMutex
	getResourcesArgsForCall []struct {
		arg1 concourse.Build
	}
	getResourcesReturns struct {
		result1 concourse.Run
		result2 error
	}
	getResourcesReturnsOnCall map[int]struct {
		result1 concourse.Run
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClientInterface) GetPipeline(arg1 string) (concourse.Pipeline, error) {
	fake.getPipelineMutex.Lock()
	ret, specificReturn := fake.getPipelineReturnsOnCall[len(fake.getPipelineArgsForCall)]
	fake.getPipelineArgsForCall = append(fake.getPipelineArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetPipeline", []interface{}{arg1})
	fake.getPipelineMutex.Unlock()
	if fake.GetPipelineStub != nil {
		return fake.GetPipelineStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getPipelineReturns.result1, fake.getPipelineReturns.result2
}

func (fake *FakeClientInterface) GetPipelineCallCount() int {
	fake.getPipelineMutex.RLock()
	defer fake.getPipelineMutex.RUnlock()
	return len(fake.getPipelineArgsForCall)
}

func (fake *FakeClientInterface) GetPipelineArgsForCall(i int) string {
	fake.getPipelineMutex.RLock()
	defer fake.getPipelineMutex.RUnlock()
	return fake.getPipelineArgsForCall[i].arg1
}

func (fake *FakeClientInterface) GetPipelineReturns(result1 concourse.Pipeline, result2 error) {
	fake.GetPipelineStub = nil
	fake.getPipelineReturns = struct {
		result1 concourse.Pipeline
		result2 error
	}{result1, result2}
}

func (fake *FakeClientInterface) GetPipelineReturnsOnCall(i int, result1 concourse.Pipeline, result2 error) {
	fake.GetPipelineStub = nil
	if fake.getPipelineReturnsOnCall == nil {
		fake.getPipelineReturnsOnCall = make(map[int]struct {
			result1 concourse.Pipeline
			result2 error
		})
	}
	fake.getPipelineReturnsOnCall[i] = struct {
		result1 concourse.Pipeline
		result2 error
	}{result1, result2}
}

func (fake *FakeClientInterface) GetBuilds(arg1 concourse.Job) ([]concourse.Build, error) {
	fake.getBuildsMutex.Lock()
	ret, specificReturn := fake.getBuildsReturnsOnCall[len(fake.getBuildsArgsForCall)]
	fake.getBuildsArgsForCall = append(fake.getBuildsArgsForCall, struct {
		arg1 concourse.Job
	}{arg1})
	fake.recordInvocation("GetBuilds", []interface{}{arg1})
	fake.getBuildsMutex.Unlock()
	if fake.GetBuildsStub != nil {
		return fake.GetBuildsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getBuildsReturns.result1, fake.getBuildsReturns.result2
}

func (fake *FakeClientInterface) GetBuildsCallCount() int {
	fake.getBuildsMutex.RLock()
	defer fake.getBuildsMutex.RUnlock()
	return len(fake.getBuildsArgsForCall)
}

func (fake *FakeClientInterface) GetBuildsArgsForCall(i int) concourse.Job {
	fake.getBuildsMutex.RLock()
	defer fake.getBuildsMutex.RUnlock()
	return fake.getBuildsArgsForCall[i].arg1
}

func (fake *FakeClientInterface) GetBuildsReturns(result1 []concourse.Build, result2 error) {
	fake.GetBuildsStub = nil
	fake.getBuildsReturns = struct {
		result1 []concourse.Build
		result2 error
	}{result1, result2}
}

func (fake *FakeClientInterface) GetBuildsReturnsOnCall(i int, result1 []concourse.Build, result2 error) {
	fake.GetBuildsStub = nil
	if fake.getBuildsReturnsOnCall == nil {
		fake.getBuildsReturnsOnCall = make(map[int]struct {
			result1 []concourse.Build
			result2 error
		})
	}
	fake.getBuildsReturnsOnCall[i] = struct {
		result1 []concourse.Build
		result2 error
	}{result1, result2}
}

func (fake *FakeClientInterface) GetResources(arg1 concourse.Build) (concourse.Run, error) {
	fake.getResourcesMutex.Lock()
	ret, specificReturn := fake.getResourcesReturnsOnCall[len(fake.getResourcesArgsForCall)]
	fake.getResourcesArgsForCall = append(fake.getResourcesArgsForCall, struct {
		arg1 concourse.Build
	}{arg1})
	fake.recordInvocation("GetResources", []interface{}{arg1})
	fake.getResourcesMutex.Unlock()
	if fake.GetResourcesStub != nil {
		return fake.GetResourcesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getResourcesReturns.result1, fake.getResourcesReturns.result2
}

func (fake *FakeClientInterface) GetResourcesCallCount() int {
	fake.getResourcesMutex.RLock()
	defer fake.getResourcesMutex.RUnlock()
	return len(fake.getResourcesArgsForCall)
}

func (fake *FakeClientInterface) GetResourcesArgsForCall(i int) concourse.Build {
	fake.getResourcesMutex.RLock()
	defer fake.getResourcesMutex.RUnlock()
	return fake.getResourcesArgsForCall[i].arg1
}

func (fake *FakeClientInterface) GetResourcesReturns(result1 concourse.Run, result2 error) {
	fake.GetResourcesStub = nil
	fake.getResourcesReturns = struct {
		result1 concourse.Run
		result2 error
	}{result1, result2}
}

func (fake *FakeClientInterface) GetResourcesReturnsOnCall(i int, result1 concourse.Run, result2 error) {
	fake.GetResourcesStub = nil
	if fake.getResourcesReturnsOnCall == nil {
		fake.getResourcesReturnsOnCall = make(map[int]struct {
			result1 concourse.Run
			result2 error
		})
	}
	fake.getResourcesReturnsOnCall[i] = struct {
		result1 concourse.Run
		result2 error
	}{result1, result2}
}

func (fake *FakeClientInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getPipelineMutex.RLock()
	defer fake.getPipelineMutex.RUnlock()
	fake.getBuildsMutex.RLock()
	defer fake.getBuildsMutex.RUnlock()
	fake.getResourcesMutex.RLock()
	defer fake.getResourcesMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeClientInterface) recordInvocation(key string, args []interface{}) {
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

var _ concourse.ClientInterface = new(FakeClientInterface)