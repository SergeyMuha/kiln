// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/kiln/fetcher"
	"github.com/pivotal-cf/kiln/release"
)

type ReleaseSource struct {
	DownloadReleaseStub        func(string, release.Remote, int) (release.Local, error)
	downloadReleaseMutex       sync.RWMutex
	downloadReleaseArgsForCall []struct {
		arg1 string
		arg2 release.Remote
		arg3 int
	}
	downloadReleaseReturns struct {
		result1 release.Local
		result2 error
	}
	downloadReleaseReturnsOnCall map[int]struct {
		result1 release.Local
		result2 error
	}
	GetMatchedReleaseStub        func(release.Requirement) (release.Remote, bool, error)
	getMatchedReleaseMutex       sync.RWMutex
	getMatchedReleaseArgsForCall []struct {
		arg1 release.Requirement
	}
	getMatchedReleaseReturns struct {
		result1 release.Remote
		result2 bool
		result3 error
	}
	getMatchedReleaseReturnsOnCall map[int]struct {
		result1 release.Remote
		result2 bool
		result3 error
	}
	IDStub        func() string
	iDMutex       sync.RWMutex
	iDArgsForCall []struct {
	}
	iDReturns struct {
		result1 string
	}
	iDReturnsOnCall map[int]struct {
		result1 string
	}
	PublishableStub        func() bool
	publishableMutex       sync.RWMutex
	publishableArgsForCall []struct {
	}
	publishableReturns struct {
		result1 bool
	}
	publishableReturnsOnCall map[int]struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ReleaseSource) DownloadRelease(arg1 string, arg2 release.Remote, arg3 int) (release.Local, error) {
	fake.downloadReleaseMutex.Lock()
	ret, specificReturn := fake.downloadReleaseReturnsOnCall[len(fake.downloadReleaseArgsForCall)]
	fake.downloadReleaseArgsForCall = append(fake.downloadReleaseArgsForCall, struct {
		arg1 string
		arg2 release.Remote
		arg3 int
	}{arg1, arg2, arg3})
	fake.recordInvocation("DownloadRelease", []interface{}{arg1, arg2, arg3})
	fake.downloadReleaseMutex.Unlock()
	if fake.DownloadReleaseStub != nil {
		return fake.DownloadReleaseStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.downloadReleaseReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ReleaseSource) DownloadReleaseCallCount() int {
	fake.downloadReleaseMutex.RLock()
	defer fake.downloadReleaseMutex.RUnlock()
	return len(fake.downloadReleaseArgsForCall)
}

func (fake *ReleaseSource) DownloadReleaseCalls(stub func(string, release.Remote, int) (release.Local, error)) {
	fake.downloadReleaseMutex.Lock()
	defer fake.downloadReleaseMutex.Unlock()
	fake.DownloadReleaseStub = stub
}

func (fake *ReleaseSource) DownloadReleaseArgsForCall(i int) (string, release.Remote, int) {
	fake.downloadReleaseMutex.RLock()
	defer fake.downloadReleaseMutex.RUnlock()
	argsForCall := fake.downloadReleaseArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *ReleaseSource) DownloadReleaseReturns(result1 release.Local, result2 error) {
	fake.downloadReleaseMutex.Lock()
	defer fake.downloadReleaseMutex.Unlock()
	fake.DownloadReleaseStub = nil
	fake.downloadReleaseReturns = struct {
		result1 release.Local
		result2 error
	}{result1, result2}
}

func (fake *ReleaseSource) DownloadReleaseReturnsOnCall(i int, result1 release.Local, result2 error) {
	fake.downloadReleaseMutex.Lock()
	defer fake.downloadReleaseMutex.Unlock()
	fake.DownloadReleaseStub = nil
	if fake.downloadReleaseReturnsOnCall == nil {
		fake.downloadReleaseReturnsOnCall = make(map[int]struct {
			result1 release.Local
			result2 error
		})
	}
	fake.downloadReleaseReturnsOnCall[i] = struct {
		result1 release.Local
		result2 error
	}{result1, result2}
}

func (fake *ReleaseSource) GetMatchedRelease(arg1 release.Requirement) (release.Remote, bool, error) {
	fake.getMatchedReleaseMutex.Lock()
	ret, specificReturn := fake.getMatchedReleaseReturnsOnCall[len(fake.getMatchedReleaseArgsForCall)]
	fake.getMatchedReleaseArgsForCall = append(fake.getMatchedReleaseArgsForCall, struct {
		arg1 release.Requirement
	}{arg1})
	fake.recordInvocation("GetMatchedRelease", []interface{}{arg1})
	fake.getMatchedReleaseMutex.Unlock()
	if fake.GetMatchedReleaseStub != nil {
		return fake.GetMatchedReleaseStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getMatchedReleaseReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *ReleaseSource) GetMatchedReleaseCallCount() int {
	fake.getMatchedReleaseMutex.RLock()
	defer fake.getMatchedReleaseMutex.RUnlock()
	return len(fake.getMatchedReleaseArgsForCall)
}

func (fake *ReleaseSource) GetMatchedReleaseCalls(stub func(release.Requirement) (release.Remote, bool, error)) {
	fake.getMatchedReleaseMutex.Lock()
	defer fake.getMatchedReleaseMutex.Unlock()
	fake.GetMatchedReleaseStub = stub
}

func (fake *ReleaseSource) GetMatchedReleaseArgsForCall(i int) release.Requirement {
	fake.getMatchedReleaseMutex.RLock()
	defer fake.getMatchedReleaseMutex.RUnlock()
	argsForCall := fake.getMatchedReleaseArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ReleaseSource) GetMatchedReleaseReturns(result1 release.Remote, result2 bool, result3 error) {
	fake.getMatchedReleaseMutex.Lock()
	defer fake.getMatchedReleaseMutex.Unlock()
	fake.GetMatchedReleaseStub = nil
	fake.getMatchedReleaseReturns = struct {
		result1 release.Remote
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *ReleaseSource) GetMatchedReleaseReturnsOnCall(i int, result1 release.Remote, result2 bool, result3 error) {
	fake.getMatchedReleaseMutex.Lock()
	defer fake.getMatchedReleaseMutex.Unlock()
	fake.GetMatchedReleaseStub = nil
	if fake.getMatchedReleaseReturnsOnCall == nil {
		fake.getMatchedReleaseReturnsOnCall = make(map[int]struct {
			result1 release.Remote
			result2 bool
			result3 error
		})
	}
	fake.getMatchedReleaseReturnsOnCall[i] = struct {
		result1 release.Remote
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *ReleaseSource) ID() string {
	fake.iDMutex.Lock()
	ret, specificReturn := fake.iDReturnsOnCall[len(fake.iDArgsForCall)]
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct {
	}{})
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if fake.IDStub != nil {
		return fake.IDStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.iDReturns
	return fakeReturns.result1
}

func (fake *ReleaseSource) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *ReleaseSource) IDCalls(stub func() string) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = stub
}

func (fake *ReleaseSource) IDReturns(result1 string) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 string
	}{result1}
}

func (fake *ReleaseSource) IDReturnsOnCall(i int, result1 string) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	if fake.iDReturnsOnCall == nil {
		fake.iDReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.iDReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *ReleaseSource) Publishable() bool {
	fake.publishableMutex.Lock()
	ret, specificReturn := fake.publishableReturnsOnCall[len(fake.publishableArgsForCall)]
	fake.publishableArgsForCall = append(fake.publishableArgsForCall, struct {
	}{})
	fake.recordInvocation("Publishable", []interface{}{})
	fake.publishableMutex.Unlock()
	if fake.PublishableStub != nil {
		return fake.PublishableStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.publishableReturns
	return fakeReturns.result1
}

func (fake *ReleaseSource) PublishableCallCount() int {
	fake.publishableMutex.RLock()
	defer fake.publishableMutex.RUnlock()
	return len(fake.publishableArgsForCall)
}

func (fake *ReleaseSource) PublishableCalls(stub func() bool) {
	fake.publishableMutex.Lock()
	defer fake.publishableMutex.Unlock()
	fake.PublishableStub = stub
}

func (fake *ReleaseSource) PublishableReturns(result1 bool) {
	fake.publishableMutex.Lock()
	defer fake.publishableMutex.Unlock()
	fake.PublishableStub = nil
	fake.publishableReturns = struct {
		result1 bool
	}{result1}
}

func (fake *ReleaseSource) PublishableReturnsOnCall(i int, result1 bool) {
	fake.publishableMutex.Lock()
	defer fake.publishableMutex.Unlock()
	fake.PublishableStub = nil
	if fake.publishableReturnsOnCall == nil {
		fake.publishableReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.publishableReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *ReleaseSource) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.downloadReleaseMutex.RLock()
	defer fake.downloadReleaseMutex.RUnlock()
	fake.getMatchedReleaseMutex.RLock()
	defer fake.getMatchedReleaseMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.publishableMutex.RLock()
	defer fake.publishableMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ReleaseSource) recordInvocation(key string, args []interface{}) {
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

var _ fetcher.ReleaseSource = new(ReleaseSource)
