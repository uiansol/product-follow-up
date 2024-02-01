// Code generated by mockery v2.40.1. DO NOT EDIT.

package usecases

import (
	mock "github.com/stretchr/testify/mock"
)

// IProductReadUseCaseMock is an autogenerated mock type for the IProductReadUseCase type
type IProductReadUseCaseMock struct {
	mock.Mock
}

// Execute provides a mock function with given fields: productReadInput
func (_m *IProductReadUseCaseMock) Execute(productReadInput ProductReadInput) (ProductReadOutput, error) {
	ret := _m.Called(productReadInput)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 ProductReadOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(ProductReadInput) (ProductReadOutput, error)); ok {
		return rf(productReadInput)
	}
	if rf, ok := ret.Get(0).(func(ProductReadInput) ProductReadOutput); ok {
		r0 = rf(productReadInput)
	} else {
		r0 = ret.Get(0).(ProductReadOutput)
	}

	if rf, ok := ret.Get(1).(func(ProductReadInput) error); ok {
		r1 = rf(productReadInput)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIProductReadUseCaseMock creates a new instance of IProductReadUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIProductReadUseCaseMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *IProductReadUseCaseMock {
	mock := &IProductReadUseCaseMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}