// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/pool-resource/out"
)

type FakeLockHandler struct {
	GrabAvailableLockStub        func() (lock string, version string, err error)
	grabAvailableLockMutex       sync.RWMutex
	grabAvailableLockArgsForCall []struct{}
	grabAvailableLockReturns struct {
		result1 string
		result2 string
		result3 error
	}
	UnclaimLockStub        func(lock string) (version string, err error)
	unclaimLockMutex       sync.RWMutex
	unclaimLockArgsForCall []struct {
		lock string
	}
	unclaimLockReturns struct {
		result1 string
		result2 error
	}
	AddLockStub        func(lock string, contents []byte) (version string, err error)
	addLockMutex       sync.RWMutex
	addLockArgsForCall []struct {
		lock     string
		contents []byte
	}
	addLockReturns struct {
		result1 string
		result2 error
	}
	SetupStub        func() error
	setupMutex       sync.RWMutex
	setupArgsForCall []struct{}
	setupReturns struct {
		result1 error
	}
	BroadcastLockPoolStub        func() error
	broadcastLockPoolMutex       sync.RWMutex
	broadcastLockPoolArgsForCall []struct{}
	broadcastLockPoolReturns struct {
		result1 error
	}
	ResetLockStub        func() error
	resetLockMutex       sync.RWMutex
	resetLockArgsForCall []struct{}
	resetLockReturns struct {
		result1 error
	}
}

func (fake *FakeLockHandler) GrabAvailableLock() (lock string, version string, err error) {
	fake.grabAvailableLockMutex.Lock()
	fake.grabAvailableLockArgsForCall = append(fake.grabAvailableLockArgsForCall, struct{}{})
	fake.grabAvailableLockMutex.Unlock()
	if fake.GrabAvailableLockStub != nil {
		return fake.GrabAvailableLockStub()
	} else {
		return fake.grabAvailableLockReturns.result1, fake.grabAvailableLockReturns.result2, fake.grabAvailableLockReturns.result3
	}
}

func (fake *FakeLockHandler) GrabAvailableLockCallCount() int {
	fake.grabAvailableLockMutex.RLock()
	defer fake.grabAvailableLockMutex.RUnlock()
	return len(fake.grabAvailableLockArgsForCall)
}

func (fake *FakeLockHandler) GrabAvailableLockReturns(result1 string, result2 string, result3 error) {
	fake.GrabAvailableLockStub = nil
	fake.grabAvailableLockReturns = struct {
		result1 string
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeLockHandler) UnclaimLock(lock string) (version string, err error) {
	fake.unclaimLockMutex.Lock()
	fake.unclaimLockArgsForCall = append(fake.unclaimLockArgsForCall, struct {
		lock string
	}{lock})
	fake.unclaimLockMutex.Unlock()
	if fake.UnclaimLockStub != nil {
		return fake.UnclaimLockStub(lock)
	} else {
		return fake.unclaimLockReturns.result1, fake.unclaimLockReturns.result2
	}
}

func (fake *FakeLockHandler) UnclaimLockCallCount() int {
	fake.unclaimLockMutex.RLock()
	defer fake.unclaimLockMutex.RUnlock()
	return len(fake.unclaimLockArgsForCall)
}

func (fake *FakeLockHandler) UnclaimLockArgsForCall(i int) string {
	fake.unclaimLockMutex.RLock()
	defer fake.unclaimLockMutex.RUnlock()
	return fake.unclaimLockArgsForCall[i].lock
}

func (fake *FakeLockHandler) UnclaimLockReturns(result1 string, result2 error) {
	fake.UnclaimLockStub = nil
	fake.unclaimLockReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeLockHandler) AddLock(lock string, contents []byte) (version string, err error) {
	fake.addLockMutex.Lock()
	fake.addLockArgsForCall = append(fake.addLockArgsForCall, struct {
		lock     string
		contents []byte
	}{lock, contents})
	fake.addLockMutex.Unlock()
	if fake.AddLockStub != nil {
		return fake.AddLockStub(lock, contents)
	} else {
		return fake.addLockReturns.result1, fake.addLockReturns.result2
	}
}

func (fake *FakeLockHandler) AddLockCallCount() int {
	fake.addLockMutex.RLock()
	defer fake.addLockMutex.RUnlock()
	return len(fake.addLockArgsForCall)
}

func (fake *FakeLockHandler) AddLockArgsForCall(i int) (string, []byte) {
	fake.addLockMutex.RLock()
	defer fake.addLockMutex.RUnlock()
	return fake.addLockArgsForCall[i].lock, fake.addLockArgsForCall[i].contents
}

func (fake *FakeLockHandler) AddLockReturns(result1 string, result2 error) {
	fake.AddLockStub = nil
	fake.addLockReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeLockHandler) Setup() error {
	fake.setupMutex.Lock()
	fake.setupArgsForCall = append(fake.setupArgsForCall, struct{}{})
	fake.setupMutex.Unlock()
	if fake.SetupStub != nil {
		return fake.SetupStub()
	} else {
		return fake.setupReturns.result1
	}
}

func (fake *FakeLockHandler) SetupCallCount() int {
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	return len(fake.setupArgsForCall)
}

func (fake *FakeLockHandler) SetupReturns(result1 error) {
	fake.SetupStub = nil
	fake.setupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLockHandler) BroadcastLockPool() error {
	fake.broadcastLockPoolMutex.Lock()
	fake.broadcastLockPoolArgsForCall = append(fake.broadcastLockPoolArgsForCall, struct{}{})
	fake.broadcastLockPoolMutex.Unlock()
	if fake.BroadcastLockPoolStub != nil {
		return fake.BroadcastLockPoolStub()
	} else {
		return fake.broadcastLockPoolReturns.result1
	}
}

func (fake *FakeLockHandler) BroadcastLockPoolCallCount() int {
	fake.broadcastLockPoolMutex.RLock()
	defer fake.broadcastLockPoolMutex.RUnlock()
	return len(fake.broadcastLockPoolArgsForCall)
}

func (fake *FakeLockHandler) BroadcastLockPoolReturns(result1 error) {
	fake.BroadcastLockPoolStub = nil
	fake.broadcastLockPoolReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLockHandler) ResetLock() error {
	fake.resetLockMutex.Lock()
	fake.resetLockArgsForCall = append(fake.resetLockArgsForCall, struct{}{})
	fake.resetLockMutex.Unlock()
	if fake.ResetLockStub != nil {
		return fake.ResetLockStub()
	} else {
		return fake.resetLockReturns.result1
	}
}

func (fake *FakeLockHandler) ResetLockCallCount() int {
	fake.resetLockMutex.RLock()
	defer fake.resetLockMutex.RUnlock()
	return len(fake.resetLockArgsForCall)
}

func (fake *FakeLockHandler) ResetLockReturns(result1 error) {
	fake.ResetLockStub = nil
	fake.resetLockReturns = struct {
		result1 error
	}{result1}
}

var _ out.LockHandler = new(FakeLockHandler)
