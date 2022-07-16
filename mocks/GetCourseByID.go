// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	domain "github.com/jeanmolossi/effective-eureka/src/core/courses/domain"
	mock "github.com/stretchr/testify/mock"
)

// GetCourseByID is an autogenerated mock type for the GetCourseByID type
type GetCourseByID struct {
	mock.Mock
}

// Run provides a mock function with given fields: courseID
func (_m *GetCourseByID) Run(courseID string) (domain.Course, error) {
	ret := _m.Called(courseID)

	var r0 domain.Course
	if rf, ok := ret.Get(0).(func(string) domain.Course); ok {
		r0 = rf(courseID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.Course)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(courseID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}