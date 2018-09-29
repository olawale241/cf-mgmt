// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/pivotalservices/cf-mgmt/ldap"
)

type FakeManager struct {
	GetUserDNsStub        func(groupName string) ([]string, error)
	getUserDNsMutex       sync.RWMutex
	getUserDNsArgsForCall []struct {
		groupName string
	}
	getUserDNsReturns struct {
		result1 []string
		result2 error
	}
	GetUserByIDStub        func(userID string) (*ldap.User, error)
	getUserByIDMutex       sync.RWMutex
	getUserByIDArgsForCall []struct {
		userID string
	}
	getUserByIDReturns struct {
		result1 *ldap.User
		result2 error
	}
	GetUserByDNStub        func(userDN string) (*ldap.User, error)
	getUserByDNMutex       sync.RWMutex
	getUserByDNArgsForCall []struct {
		userDN string
	}
	getUserByDNReturns struct {
		result1 *ldap.User
		result2 error
	}
	CloseStub        func()
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeManager) GetUserDNs(groupName string) ([]string, error) {
	fake.getUserDNsMutex.Lock()
	fake.getUserDNsArgsForCall = append(fake.getUserDNsArgsForCall, struct {
		groupName string
	}{groupName})
	fake.recordInvocation("GetUserDNs", []interface{}{groupName})
	fake.getUserDNsMutex.Unlock()
	if fake.GetUserDNsStub != nil {
		return fake.GetUserDNsStub(groupName)
	} else {
		return fake.getUserDNsReturns.result1, fake.getUserDNsReturns.result2
	}
}

func (fake *FakeManager) GetUserDNsCallCount() int {
	fake.getUserDNsMutex.RLock()
	defer fake.getUserDNsMutex.RUnlock()
	return len(fake.getUserDNsArgsForCall)
}

func (fake *FakeManager) GetUserDNsArgsForCall(i int) string {
	fake.getUserDNsMutex.RLock()
	defer fake.getUserDNsMutex.RUnlock()
	return fake.getUserDNsArgsForCall[i].groupName
}

func (fake *FakeManager) GetUserDNsReturns(result1 []string, result2 error) {
	fake.GetUserDNsStub = nil
	fake.getUserDNsReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) GetUserByID(userID string) (*ldap.User, error) {
	fake.getUserByIDMutex.Lock()
	fake.getUserByIDArgsForCall = append(fake.getUserByIDArgsForCall, struct {
		userID string
	}{userID})
	fake.recordInvocation("GetUserByID", []interface{}{userID})
	fake.getUserByIDMutex.Unlock()
	if fake.GetUserByIDStub != nil {
		return fake.GetUserByIDStub(userID)
	} else {
		return fake.getUserByIDReturns.result1, fake.getUserByIDReturns.result2
	}
}

func (fake *FakeManager) GetUserByIDCallCount() int {
	fake.getUserByIDMutex.RLock()
	defer fake.getUserByIDMutex.RUnlock()
	return len(fake.getUserByIDArgsForCall)
}

func (fake *FakeManager) GetUserByIDArgsForCall(i int) string {
	fake.getUserByIDMutex.RLock()
	defer fake.getUserByIDMutex.RUnlock()
	return fake.getUserByIDArgsForCall[i].userID
}

func (fake *FakeManager) GetUserByIDReturns(result1 *ldap.User, result2 error) {
	fake.GetUserByIDStub = nil
	fake.getUserByIDReturns = struct {
		result1 *ldap.User
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) GetUserByDN(userDN string) (*ldap.User, error) {
	fake.getUserByDNMutex.Lock()
	fake.getUserByDNArgsForCall = append(fake.getUserByDNArgsForCall, struct {
		userDN string
	}{userDN})
	fake.recordInvocation("GetUserByDN", []interface{}{userDN})
	fake.getUserByDNMutex.Unlock()
	if fake.GetUserByDNStub != nil {
		return fake.GetUserByDNStub(userDN)
	} else {
		return fake.getUserByDNReturns.result1, fake.getUserByDNReturns.result2
	}
}

func (fake *FakeManager) GetUserByDNCallCount() int {
	fake.getUserByDNMutex.RLock()
	defer fake.getUserByDNMutex.RUnlock()
	return len(fake.getUserByDNArgsForCall)
}

func (fake *FakeManager) GetUserByDNArgsForCall(i int) string {
	fake.getUserByDNMutex.RLock()
	defer fake.getUserByDNMutex.RUnlock()
	return fake.getUserByDNArgsForCall[i].userDN
}

func (fake *FakeManager) GetUserByDNReturns(result1 *ldap.User, result2 error) {
	fake.GetUserByDNStub = nil
	fake.getUserByDNReturns = struct {
		result1 *ldap.User
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) Close() {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		fake.CloseStub()
	}
}

func (fake *FakeManager) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getUserDNsMutex.RLock()
	defer fake.getUserDNsMutex.RUnlock()
	fake.getUserByIDMutex.RLock()
	defer fake.getUserByIDMutex.RUnlock()
	fake.getUserByDNMutex.RLock()
	defer fake.getUserByDNMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeManager) recordInvocation(key string, args []interface{}) {
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

var _ ldap.Manager = new(FakeManager)
