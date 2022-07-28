// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	domain "github.com/jeanmolossi/effective-eureka/src/core/modules/domain"
	mock "github.com/stretchr/testify/mock"
)

// EditModuleInfo is an autogenerated mock type for the EditModuleInfo type
type EditModuleInfo struct {
	mock.Mock
}

// Run provides a mock function with given fields: module
func (_m *EditModuleInfo) Run(module domain.Module) (domain.Module, error) {
	ret := _m.Called(module)

	var r0 domain.Module
	if rf, ok := ret.Get(0).(func(domain.Module) domain.Module); ok {
		r0 = rf(module)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.Module)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Module) error); ok {
		r1 = rf(module)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
