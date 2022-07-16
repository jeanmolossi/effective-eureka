package usecase

import "github.com/jeanmolossi/effective-eureka/src/core/courses/domain"

type createCourse struct {
	repo domain.CourseRepository
}

// NewCreateCourse is a factory method to create a usecase to create a course.
func NewCreateCourse(repo domain.CourseRepository) domain.CreateCourse {
	return &createCourse{
		repo,
	}
}

// Run is the method with handles application to create a course.
func (c *createCourse) Run(course domain.Course) (domain.Course, error) {
	course.SetCourseID("course-id")
	return c.repo.Create(course)
}
