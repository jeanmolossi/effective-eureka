// Package usecase is a package to manage courses in Go.
package usecase

import "github.com/jeanmolossi/effective-eureka/src/core/courses/domain"

type getCourseByID struct {
	repo domain.CourseRepository
}

func NewGetCourseByID() domain.GetCourseByID {
	return &getCourseByID{}
}

func (g *getCourseByID) Run(courseID string) (domain.Course, error) {
	return nil, nil
}
