// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	domain "github.com/jeanmolossi/effective-eureka/src/core/lessons/domain"
	mock "github.com/stretchr/testify/mock"
)

// EditLessonInfo is an autogenerated mock type for the EditLessonInfo type
type EditLessonInfo struct {
	mock.Mock
}

// EditLesson provides a mock function with given fields: lesson
func (_m *EditLessonInfo) EditLesson(lesson domain.Lesson) (domain.Lesson, error) {
	ret := _m.Called(lesson)

	var r0 domain.Lesson
	if rf, ok := ret.Get(0).(func(domain.Lesson) domain.Lesson); ok {
		r0 = rf(lesson)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.Lesson)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Lesson) error); ok {
		r1 = rf(lesson)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}