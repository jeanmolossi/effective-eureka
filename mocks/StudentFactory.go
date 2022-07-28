// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	domain "github.com/jeanmolossi/effective-eureka/src/core/students/domain"
	factory "github.com/jeanmolossi/effective-eureka/src/core/students/factory"

	mock "github.com/stretchr/testify/mock"
)

// StudentFactory is an autogenerated mock type for the StudentFactory type
type StudentFactory struct {
	mock.Mock
}

// Build provides a mock function with given fields:
func (_m *StudentFactory) Build() domain.Student {
	ret := _m.Called()

	var r0 domain.Student
	if rf, ok := ret.Get(0).(func() domain.Student); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.Student)
		}
	}

	return r0
}

// CreateStudent provides a mock function with given fields: email, password
func (_m *StudentFactory) CreateStudent(email string, password string) factory.StudentFactory {
	ret := _m.Called(email, password)

	var r0 factory.StudentFactory
	if rf, ok := ret.Get(0).(func(string, string) factory.StudentFactory); ok {
		r0 = rf(email, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(factory.StudentFactory)
		}
	}

	return r0
}

// WithID provides a mock function with given fields: id
func (_m *StudentFactory) WithID(id string) factory.StudentFactory {
	ret := _m.Called(id)

	var r0 factory.StudentFactory
	if rf, ok := ret.Get(0).(func(string) factory.StudentFactory); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(factory.StudentFactory)
		}
	}

	return r0
}
