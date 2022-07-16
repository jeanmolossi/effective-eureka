// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	domain "github.com/jeanmolossi/effective-eureka/src/core/courses/domain"
	mock "github.com/stretchr/testify/mock"
)

// CourseRepository is an autogenerated mock type for the CourseRepository type
type CourseRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: course
func (_m *CourseRepository) Create(course domain.Course) (domain.Course, error) {
	ret := _m.Called(course)

	var r0 domain.Course
	if rf, ok := ret.Get(0).(func(domain.Course) domain.Course); ok {
		r0 = rf(course)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.Course)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Course) error); ok {
		r1 = rf(course)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Edit provides a mock function with given fields: course
func (_m *CourseRepository) Edit(course domain.Course) (domain.Course, error) {
	ret := _m.Called(course)

	var r0 domain.Course
	if rf, ok := ret.Get(0).(func(domain.Course) domain.Course); ok {
		r0 = rf(course)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.Course)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Course) error); ok {
		r1 = rf(course)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: courseID
func (_m *CourseRepository) GetByID(courseID string) (domain.Course, error) {
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

// GetByStudentID provides a mock function with given fields: studentID
func (_m *CourseRepository) GetByStudentID(studentID string) ([]domain.Course, error) {
	ret := _m.Called(studentID)

	var r0 []domain.Course
	if rf, ok := ret.Get(0).(func(string) []domain.Course); ok {
		r0 = rf(studentID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Course)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(studentID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
