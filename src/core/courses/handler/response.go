package handler

import "github.com/jeanmolossi/effective-eureka/src/core/courses/domain"

// HttpCourseCreated is a struct to manage the CourseCreated response model.
type HttpCourseCreated struct {
	CourseID        string `json:"course_id" example:"05d4d9d3-01a3-4fd3-8d3e-e3178522f514"`
	CoursePublished bool   `json:"course_published" example:"false"`
}

func NewHttpCourseCreated(course domain.Course) *HttpCourseCreated {
	return &HttpCourseCreated{
		CourseID:        course.GetCourseID(),
		CoursePublished: course.IsCoursePublished(),
	}
}
