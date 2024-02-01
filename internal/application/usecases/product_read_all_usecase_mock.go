// Code generated by mockery v2.40.1. DO NOT EDIT.

package usecases

import (
	mock "github.com/stretchr/testify/mock"
)

// IProductReadAllUseCaseMock is an autogenerated mock type for the IProductReadAllUseCaseMock type
type IProductReadAllUseCaseMock struct {
	mock.Mock
}

// Execute provides a mock function with given fields:
func (_m *IProductReadAllUseCaseMock) Execute() ([]*ProductReadAllOutput, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 []*ProductReadAllOutput
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*ProductReadAllOutput, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*ProductReadAllOutput); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ProductReadAllOutput)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIProductReadAllUseCaseMock creates a new instance of IProductReadAllUseCaseMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIProductReadAllUseCaseMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *IProductReadAllUseCaseMock {
	mock := &IProductReadAllUseCaseMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
