// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// AccountInfo is an autogenerated mock type for the AccountInfo type
type AccountInfo struct {
	mock.Mock
}

// NewAccountInfo creates a new instance of AccountInfo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAccountInfo(t interface {
	mock.TestingT
	Cleanup(func())
}) *AccountInfo {
	mock := &AccountInfo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
