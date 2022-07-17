// Package usecase is a package to manage courses in Go.
package usecase

import "github.com/jeanmolossi/effective-eureka/src/core/courses/domain"

type getCourseByID struct {
	repo domain.CourseRepository
}

// NewGetCourseByID is a factory method to create a usecase to get a course by ID.
func NewGetCourseByID(repo domain.CourseRepository) domain.GetCourseByID {
	return &getCourseByID{repo}
}

// Run is the method to get a course by ID. That implements the usecase.GetCourseByID interface.
func (g *getCourseByID) Run(courseID string) (domain.Course, error) {
	return g.repo.GetByID(courseID)
}
