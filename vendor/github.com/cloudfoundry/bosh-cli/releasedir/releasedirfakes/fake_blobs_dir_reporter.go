// This file was generated by counterfeiter
package releasedirfakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-cli/releasedir"
)

type FakeBlobsDirReporter struct {
	BlobDownloadStartedStub        func(path string, size int64, blobID, sha1 string)
	blobDownloadStartedMutex       sync.RWMutex
	blobDownloadStartedArgsForCall []struct {
		path   string
		size   int64
		blobID string
		sha1   string
	}
	BlobDownloadFinishedStub        func(path, blobID string, err error)
	blobDownloadFinishedMutex       sync.RWMutex
	blobDownloadFinishedArgsForCall []struct {
		path   string
		blobID string
		err    error
	}
	BlobUploadStartedStub        func(path string, size int64, sha1 string)
	blobUploadStartedMutex       sync.RWMutex
	blobUploadStartedArgsForCall []struct {
		path string
		size int64
		sha1 string
	}
	BlobUploadFinishedStub        func(path, blobID string, err error)
	blobUploadFinishedMutex       sync.RWMutex
	blobUploadFinishedArgsForCall []struct {
		path   string
		blobID string
		err    error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBlobsDirReporter) BlobDownloadStarted(path string, size int64, blobID string, sha1 string) {
	fake.blobDownloadStartedMutex.Lock()
	fake.blobDownloadStartedArgsForCall = append(fake.blobDownloadStartedArgsForCall, struct {
		path   string
		size   int64
		blobID string
		sha1   string
	}{path, size, blobID, sha1})
	fake.recordInvocation("BlobDownloadStarted", []interface{}{path, size, blobID, sha1})
	fake.blobDownloadStartedMutex.Unlock()
	if fake.BlobDownloadStartedStub != nil {
		fake.BlobDownloadStartedStub(path, size, blobID, sha1)
	}
}

func (fake *FakeBlobsDirReporter) BlobDownloadStartedCallCount() int {
	fake.blobDownloadStartedMutex.RLock()
	defer fake.blobDownloadStartedMutex.RUnlock()
	return len(fake.blobDownloadStartedArgsForCall)
}

func (fake *FakeBlobsDirReporter) BlobDownloadStartedArgsForCall(i int) (string, int64, string, string) {
	fake.blobDownloadStartedMutex.RLock()
	defer fake.blobDownloadStartedMutex.RUnlock()
	return fake.blobDownloadStartedArgsForCall[i].path, fake.blobDownloadStartedArgsForCall[i].size, fake.blobDownloadStartedArgsForCall[i].blobID, fake.blobDownloadStartedArgsForCall[i].sha1
}

func (fake *FakeBlobsDirReporter) BlobDownloadFinished(path string, blobID string, err error) {
	fake.blobDownloadFinishedMutex.Lock()
	fake.blobDownloadFinishedArgsForCall = append(fake.blobDownloadFinishedArgsForCall, struct {
		path   string
		blobID string
		err    error
	}{path, blobID, err})
	fake.recordInvocation("BlobDownloadFinished", []interface{}{path, blobID, err})
	fake.blobDownloadFinishedMutex.Unlock()
	if fake.BlobDownloadFinishedStub != nil {
		fake.BlobDownloadFinishedStub(path, blobID, err)
	}
}

func (fake *FakeBlobsDirReporter) BlobDownloadFinishedCallCount() int {
	fake.blobDownloadFinishedMutex.RLock()
	defer fake.blobDownloadFinishedMutex.RUnlock()
	return len(fake.blobDownloadFinishedArgsForCall)
}

func (fake *FakeBlobsDirReporter) BlobDownloadFinishedArgsForCall(i int) (string, string, error) {
	fake.blobDownloadFinishedMutex.RLock()
	defer fake.blobDownloadFinishedMutex.RUnlock()
	return fake.blobDownloadFinishedArgsForCall[i].path, fake.blobDownloadFinishedArgsForCall[i].blobID, fake.blobDownloadFinishedArgsForCall[i].err
}

func (fake *FakeBlobsDirReporter) BlobUploadStarted(path string, size int64, sha1 string) {
	fake.blobUploadStartedMutex.Lock()
	fake.blobUploadStartedArgsForCall = append(fake.blobUploadStartedArgsForCall, struct {
		path string
		size int64
		sha1 string
	}{path, size, sha1})
	fake.recordInvocation("BlobUploadStarted", []interface{}{path, size, sha1})
	fake.blobUploadStartedMutex.Unlock()
	if fake.BlobUploadStartedStub != nil {
		fake.BlobUploadStartedStub(path, size, sha1)
	}
}

func (fake *FakeBlobsDirReporter) BlobUploadStartedCallCount() int {
	fake.blobUploadStartedMutex.RLock()
	defer fake.blobUploadStartedMutex.RUnlock()
	return len(fake.blobUploadStartedArgsForCall)
}

func (fake *FakeBlobsDirReporter) BlobUploadStartedArgsForCall(i int) (string, int64, string) {
	fake.blobUploadStartedMutex.RLock()
	defer fake.blobUploadStartedMutex.RUnlock()
	return fake.blobUploadStartedArgsForCall[i].path, fake.blobUploadStartedArgsForCall[i].size, fake.blobUploadStartedArgsForCall[i].sha1
}

func (fake *FakeBlobsDirReporter) BlobUploadFinished(path string, blobID string, err error) {
	fake.blobUploadFinishedMutex.Lock()
	fake.blobUploadFinishedArgsForCall = append(fake.blobUploadFinishedArgsForCall, struct {
		path   string
		blobID string
		err    error
	}{path, blobID, err})
	fake.recordInvocation("BlobUploadFinished", []interface{}{path, blobID, err})
	fake.blobUploadFinishedMutex.Unlock()
	if fake.BlobUploadFinishedStub != nil {
		fake.BlobUploadFinishedStub(path, blobID, err)
	}
}

func (fake *FakeBlobsDirReporter) BlobUploadFinishedCallCount() int {
	fake.blobUploadFinishedMutex.RLock()
	defer fake.blobUploadFinishedMutex.RUnlock()
	return len(fake.blobUploadFinishedArgsForCall)
}

func (fake *FakeBlobsDirReporter) BlobUploadFinishedArgsForCall(i int) (string, string, error) {
	fake.blobUploadFinishedMutex.RLock()
	defer fake.blobUploadFinishedMutex.RUnlock()
	return fake.blobUploadFinishedArgsForCall[i].path, fake.blobUploadFinishedArgsForCall[i].blobID, fake.blobUploadFinishedArgsForCall[i].err
}

func (fake *FakeBlobsDirReporter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.blobDownloadStartedMutex.RLock()
	defer fake.blobDownloadStartedMutex.RUnlock()
	fake.blobDownloadFinishedMutex.RLock()
	defer fake.blobDownloadFinishedMutex.RUnlock()
	fake.blobUploadStartedMutex.RLock()
	defer fake.blobUploadStartedMutex.RUnlock()
	fake.blobUploadFinishedMutex.RLock()
	defer fake.blobUploadFinishedMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeBlobsDirReporter) recordInvocation(key string, args []interface{}) {
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

var _ releasedir.BlobsDirReporter = new(FakeBlobsDirReporter)
