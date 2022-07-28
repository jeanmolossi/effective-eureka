// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	domain "github.com/jeanmolossi/effective-eureka/src/core/students/domain"
	mock "github.com/stretchr/testify/mock"
)

// StudentRepository is an autogenerated mock type for the StudentRepository type
type StudentRepository struct {
	mock.Mock
}

// CreateStudent provides a mock function with given fields: student
func (_m *StudentRepository) CreateStudent(student domain.Student) (domain.Student, error) {
	ret := _m.Called(student)

	var r0 domain.Student
	if rf, ok := ret.Get(0).(func(domain.Student) domain.Student); ok {
		r0 = rf(student)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.Student)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Student) error); ok {
		r1 = rf(student)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteStudent provides a mock function with given fields: studentID
func (_m *StudentRepository) DeleteStudent(studentID string) error {
	ret := _m.Called(studentID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(studentID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetStudentByEmail provides a mock function with given fields: studentEmail
func (_m *StudentRepository) GetStudentByEmail(studentEmail string) (domain.Student, error) {
	ret := _m.Called(studentEmail)

	var r0 domain.Student
	if rf, ok := ret.Get(0).(func(string) domain.Student); ok {
		r0 = rf(studentEmail)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.Student)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(studentEmail)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStudentByID provides a mock function with given fields: studentID
func (_m *StudentRepository) GetStudentByID(studentID string) (domain.Student, error) {
	ret := _m.Called(studentID)

	var r0 domain.Student
	if rf, ok := ret.Get(0).(func(string) domain.Student); ok {
		r0 = rf(studentID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.Student)
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

// UpdateStudent provides a mock function with given fields: student
func (_m *StudentRepository) UpdateStudent(student domain.Student) (domain.Student, error) {
	ret := _m.Called(student)

	var r0 domain.Student
	if rf, ok := ret.Get(0).(func(domain.Student) domain.Student); ok {
		r0 = rf(student)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.Student)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Student) error); ok {
		r1 = rf(student)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
