// This file was generated by counterfeiter
package configurerfakes

import (
	"net"
	"sync"

	"github.com/teddyking/netsetgo/configurer"
)

type FakeBridgeCreator struct {
	CreateStub        func(string, net.IP, *net.IPNet) (*net.Interface, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 string
		arg2 net.IP
		arg3 *net.IPNet
	}
	createReturns struct {
		result1 *net.Interface
		result2 error
	}
	AttachStub        func(bridge, hostVeth *net.Interface) error
	attachMutex       sync.RWMutex
	attachArgsForCall []struct {
		bridge   *net.Interface
		hostVeth *net.Interface
	}
	attachReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBridgeCreator) Create(arg1 string, arg2 net.IP, arg3 *net.IPNet) (*net.Interface, error) {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 string
		arg2 net.IP
		arg3 *net.IPNet
	}{arg1, arg2, arg3})
	fake.recordInvocation("Create", []interface{}{arg1, arg2, arg3})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1, arg2, arg3)
	} else {
		return fake.createReturns.result1, fake.createReturns.result2
	}
}

func (fake *FakeBridgeCreator) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeBridgeCreator) CreateArgsForCall(i int) (string, net.IP, *net.IPNet) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].arg1, fake.createArgsForCall[i].arg2, fake.createArgsForCall[i].arg3
}

func (fake *FakeBridgeCreator) CreateReturns(result1 *net.Interface, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 *net.Interface
		result2 error
	}{result1, result2}
}

func (fake *FakeBridgeCreator) Attach(bridge *net.Interface, hostVeth *net.Interface) error {
	fake.attachMutex.Lock()
	fake.attachArgsForCall = append(fake.attachArgsForCall, struct {
		bridge   *net.Interface
		hostVeth *net.Interface
	}{bridge, hostVeth})
	fake.recordInvocation("Attach", []interface{}{bridge, hostVeth})
	fake.attachMutex.Unlock()
	if fake.AttachStub != nil {
		return fake.AttachStub(bridge, hostVeth)
	} else {
		return fake.attachReturns.result1
	}
}

func (fake *FakeBridgeCreator) AttachCallCount() int {
	fake.attachMutex.RLock()
	defer fake.attachMutex.RUnlock()
	return len(fake.attachArgsForCall)
}

func (fake *FakeBridgeCreator) AttachArgsForCall(i int) (*net.Interface, *net.Interface) {
	fake.attachMutex.RLock()
	defer fake.attachMutex.RUnlock()
	return fake.attachArgsForCall[i].bridge, fake.attachArgsForCall[i].hostVeth
}

func (fake *FakeBridgeCreator) AttachReturns(result1 error) {
	fake.AttachStub = nil
	fake.attachReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBridgeCreator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.attachMutex.RLock()
	defer fake.attachMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeBridgeCreator) recordInvocation(key string, args []interface{}) {
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

var _ configurer.BridgeCreator = new(FakeBridgeCreator)