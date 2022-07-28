// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// FilterConditions is an autogenerated mock type for the FilterConditions type
type FilterConditions struct {
	mock.Mock
}

// Conditions provides a mock function with given fields:
func (_m *FilterConditions) Conditions() (string, []interface{}) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 []interface{}
	if rf, ok := ret.Get(1).(func() []interface{}); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]interface{})
		}
	}

	return r0, r1
}

// HasConditions provides a mock function with given fields:
func (_m *FilterConditions) HasConditions() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// OnlyFields provides a mock function with given fields: fields
func (_m *FilterConditions) OnlyFields(fields []string) []string {
	ret := _m.Called(fields)

	var r0 []string
	if rf, ok := ret.Get(0).(func([]string) []string); ok {
		r0 = rf(fields)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}
