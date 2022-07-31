// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	domain "github.com/jeanmolossi/effective-eureka/src/core/lessons/domain"
	factory "github.com/jeanmolossi/effective-eureka/src/core/lessons/factory"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Lesson is an autogenerated mock type for the Lesson type
type Lesson struct {
	mock.Mock
}

// Build provides a mock function with given fields:
func (_m *Lesson) Build() domain.Lesson {
	ret := _m.Called()

	var r0 domain.Lesson
	if rf, ok := ret.Get(0).(func() domain.Lesson); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.Lesson)
		}
	}

	return r0
}

// CreateLesson provides a mock function with given fields: title, description, thumb, index, published, createdAt, updatedAt
func (_m *Lesson) CreateLesson(title string, description string, thumb string, index uint16, published bool, createdAt *time.Time, updatedAt *time.Time) factory.Lesson {
	ret := _m.Called(title, description, thumb, index, published, createdAt, updatedAt)

	var r0 factory.Lesson
	if rf, ok := ret.Get(0).(func(string, string, string, uint16, bool, *time.Time, *time.Time) factory.Lesson); ok {
		r0 = rf(title, description, thumb, index, published, createdAt, updatedAt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(factory.Lesson)
		}
	}

	return r0
}

// WithLessonID provides a mock function with given fields: lessonID
func (_m *Lesson) WithLessonID(lessonID string) factory.Lesson {
	ret := _m.Called(lessonID)

	var r0 factory.Lesson
	if rf, ok := ret.Get(0).(func(string) factory.Lesson); ok {
		r0 = rf(lessonID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(factory.Lesson)
		}
	}

	return r0
}

// WithSectionID provides a mock function with given fields: sectionID
func (_m *Lesson) WithSectionID(sectionID string) factory.Lesson {
	ret := _m.Called(sectionID)

	var r0 factory.Lesson
	if rf, ok := ret.Get(0).(func(string) factory.Lesson); ok {
		r0 = rf(sectionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(factory.Lesson)
		}
	}

	return r0
}
