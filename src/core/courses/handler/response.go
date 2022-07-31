package handler

import (
	"github.com/jeanmolossi/effective-eureka/src/core/courses/domain"
	"github.com/jeanmolossi/effective-eureka/src/core/shared"
)

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

// HttpCourseOk is a struct to manage the CourseOk response model.
type HttpCourseOk struct {
	CourseID          string `json:"course_id,omitempty" example:"b23d00ec-69f7-4fc9-b86c-3ba91c845bff"`
	CourseTitle       string `json:"course_title,omitempty" example:"Effective Eureka"`
	CourseThumbnail   string `json:"course_thumbnail,omitempty" example:"https://effective-eureka.s3.amazonaws.com/courses/effective-eureka/thumbnail.png"`
	CourseDescription string `json:"course_description,omitempty" example:"Effective Eureka is a course about Go."`
	CoursePublished   bool   `json:"course_published,omitempty" example:"false"`
}

func NewHttpCourseOk(course domain.Course) *HttpCourseOk {
	return &HttpCourseOk{
		CourseID:          course.GetCourseID(),
		CourseTitle:       course.GetCourseTitle(),
		CourseThumbnail:   course.GetCourseThumb(),
		CourseDescription: course.GetCourseDesc(),
		CoursePublished:   course.IsCoursePublished(),
	}
}

type HttpCoursesWithMeta struct {
	Data []*HttpCourseOk `json:"data"`
	Meta shared.HttpMeta `json:"meta"`
}

func NewHttpCoursesWithMeta(courses []*HttpCourseOk, page uint16, itemsPerPage int) *HttpCoursesWithMeta {
	baseURL := "http://localhost:8080/courses"

	return &HttpCoursesWithMeta{
		Data: courses,
		Meta: shared.GetMeta(baseURL, page, itemsPerPage, len(courses)),
	}
}
