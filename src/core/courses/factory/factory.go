// Package factory is used to create courses providing the necessary dependencies.
package factory

import "github.com/jeanmolossi/effective-eureka/src/core/courses/domain"

// Course is an interface wich defines the behavior of a course.
// It's used to create courses providing the necessary dependencies.
type Course interface {
	// CreateCourse creates a new course with basic information.
	CreateCourse(title, thumb, description string, published bool) Course
	// WithID adds course ID to the course.
	WithID(id string) Course
	// Build will return the course built.
	Build() domain.Course
}

// course is a struct who implements Course interface and maintains the course state.
type course struct {
	domain.Course
}

// NewCourse returns a new instance of Course.
//
// Usage:
//	// new course factory instance with basic info
// 	courseFactory := factory.NewCourse().CreateCourse(
//		"Course Title",				// title
//		"https://www.example.com/title-thumb.jpg",	// thumb
//		"Course Description",			// description
//		true)					// published
//	// add course ID and build that
// 	course := courseFactory.WithID("1").Build()
func NewCourse() Course {
	return &course{
		domain.NewCourse("", "", "", false),
	}
}

// CreateCourse creates a new course with basic information.
func (c *course) CreateCourse(title, thumb, description string, published bool) Course {
	c.Course.SetCourseTitle(title)
	c.Course.SetCourseThumb(thumb)
	c.Course.SetCourseDesc(description)
	// only sets course as published if that is the case
	if published {
		c.Course.PublishCourse()
	}

	return c
}

// WithID adds course ID to the course.
func (c *course) WithID(id string) Course {
	c.Course.SetCourseID(id)
	return c
}

// Build will return the course built.
func (c *course) Build() domain.Course {
	return c.Course
}
