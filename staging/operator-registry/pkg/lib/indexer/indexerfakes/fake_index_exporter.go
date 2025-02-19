// Code generated by counterfeiter. DO NOT EDIT.
package indexerfakes

import (
	"sync"

	"github.com/operator-framework/operator-registry/pkg/lib/indexer"
)

type FakeIndexExporter struct {
	ExportFromIndexStub        func(indexer.ExportFromIndexRequest) error
	exportFromIndexMutex       sync.RWMutex
	exportFromIndexArgsForCall []struct {
		arg1 indexer.ExportFromIndexRequest
	}
	exportFromIndexReturns struct {
		result1 error
	}
	exportFromIndexReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeIndexExporter) ExportFromIndex(arg1 indexer.ExportFromIndexRequest) error {
	fake.exportFromIndexMutex.Lock()
	ret, specificReturn := fake.exportFromIndexReturnsOnCall[len(fake.exportFromIndexArgsForCall)]
	fake.exportFromIndexArgsForCall = append(fake.exportFromIndexArgsForCall, struct {
		arg1 indexer.ExportFromIndexRequest
	}{arg1})
	stub := fake.ExportFromIndexStub
	fakeReturns := fake.exportFromIndexReturns
	fake.recordInvocation("ExportFromIndex", []interface{}{arg1})
	fake.exportFromIndexMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIndexExporter) ExportFromIndexCallCount() int {
	fake.exportFromIndexMutex.RLock()
	defer fake.exportFromIndexMutex.RUnlock()
	return len(fake.exportFromIndexArgsForCall)
}

func (fake *FakeIndexExporter) ExportFromIndexCalls(stub func(indexer.ExportFromIndexRequest) error) {
	fake.exportFromIndexMutex.Lock()
	defer fake.exportFromIndexMutex.Unlock()
	fake.ExportFromIndexStub = stub
}

func (fake *FakeIndexExporter) ExportFromIndexArgsForCall(i int) indexer.ExportFromIndexRequest {
	fake.exportFromIndexMutex.RLock()
	defer fake.exportFromIndexMutex.RUnlock()
	argsForCall := fake.exportFromIndexArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeIndexExporter) ExportFromIndexReturns(result1 error) {
	fake.exportFromIndexMutex.Lock()
	defer fake.exportFromIndexMutex.Unlock()
	fake.ExportFromIndexStub = nil
	fake.exportFromIndexReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIndexExporter) ExportFromIndexReturnsOnCall(i int, result1 error) {
	fake.exportFromIndexMutex.Lock()
	defer fake.exportFromIndexMutex.Unlock()
	fake.ExportFromIndexStub = nil
	if fake.exportFromIndexReturnsOnCall == nil {
		fake.exportFromIndexReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.exportFromIndexReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIndexExporter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.exportFromIndexMutex.RLock()
	defer fake.exportFromIndexMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeIndexExporter) recordInvocation(key string, args []interface{}) {
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

var _ indexer.IndexExporter = new(FakeIndexExporter)
