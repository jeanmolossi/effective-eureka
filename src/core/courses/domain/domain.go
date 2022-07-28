// Package domain is a package to manage the application domain.
//
// That package is the course domain.
package domain

import "github.com/jeanmolossi/effective-eureka/src/core/shared"

// Course is a interface who provides methods to manage courses.
type Course interface {
	// GetCourseID returns the course ID.
	GetCourseID() string
	// GetCourseTitle returns the course title.
	GetCourseTitle() string
	// GetCourseThumb returns the course thumb.
	GetCourseThumb() string
	// GetCourseDesc returns the course description.
	GetCourseDesc() string
	// IsCoursePublished returns true if the course is published.
	IsCoursePublished() bool

	// SetCourseID sets the course ID.
	SetCourseID(courseID string)
	// SetCourseTitle sets the course title.
	SetCourseTitle(courseTitle string)
	// SetCourseThumb sets the course thumb.
	SetCourseThumb(courseThumb string)
	// SetCourseDesc sets the course description.
	SetCourseDesc(courseDesc string)
	// PublishCourse publishes the course.
	PublishCourse()
	// UnpublishCourse unpublishes the course.
	UnpublishCourse()
}

// GetCourseByID is a interface who provides methods to get a course by ID.
type GetCourseByID interface {
	// Run is the method to get a course by ID.
	Run(courseID string) (Course, error)
}

type GetCoursesFromStudent interface {
	Run(studentID string) ([]Course, error)
}

// CreateCourse is interface segregation to create a course.
type CreateCourse interface {
	// Run is the method with handles application to create a course.
	Run(course Course) (Course, error)
}

// EditCourseInfo is interface segregation to edit a course info.
type EditCourseInfo interface {
	// Run is the method with handles application to edit a course info.
	Run(course Course) (Course, error)
}

type PublishCourse interface {
	Run(courseID string) error
}

type UnpublishCourse interface {
	Run(courseID string) error
}

// CourseUpdater is an interface to update a course. It works
// together with the CourseRepository.Edit method.
// We can handle all course properties inside that callback function.
type CourseUpdater func(course Course) (Course, error)

// CourseRepository is the interface to manage courses on database.
// To persist and handle saved courses we should implement that interface.
type CourseRepository interface {
	// GetByID returns a course by ID.
	GetByID(courseID string) (Course, error)
	// GetByStudentID returns a list of courses from a student.
	GetByStudentID(studentID string) ([]Course, error)
	// GetCourses returns a list of courses.
	GetCourses(filters shared.FilterConditions) ([]Course, error)
	// Create creates a new course.
	Create(course Course) (Course, error)
	// EditInfo edits the course info.
	Edit(courseID string, updater CourseUpdater) (Course, error)
}
