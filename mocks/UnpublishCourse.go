// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UnpublishCourse is an autogenerated mock type for the UnpublishCourse type
type UnpublishCourse struct {
	mock.Mock
}

// Run provides a mock function with given fields: courseID
func (_m *UnpublishCourse) Run(courseID string) error {
	ret := _m.Called(courseID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(courseID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
