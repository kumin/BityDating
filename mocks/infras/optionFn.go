// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	infras "github.com/kumin/BityDating/infras"
	mock "github.com/stretchr/testify/mock"
)

// OptionFn is an autogenerated mock type for the optionFn type
type OptionFn struct {
	mock.Mock
}

// Execute provides a mock function with given fields: opt
func (_m *OptionFn) Execute(opt *infras.MysqlOption) {
	_m.Called(opt)
}

// NewOptionFn creates a new instance of OptionFn. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOptionFn(t interface {
	mock.TestingT
	Cleanup(func())
}) *OptionFn {
	mock := &OptionFn{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
