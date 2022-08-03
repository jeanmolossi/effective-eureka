package handler

import (
	"fmt"

	"github.com/jeanmolossi/effective-eureka/src/core/lessons/domain"
	"github.com/jeanmolossi/effective-eureka/src/core/shared"
)

// HttpLessonCreated is a struct to manage the LessonCreated response model.
type HttpLessonCreated struct {
	LessonID          string `json:"lesson_id" example:"05d4d9d3-01a3-4fd3-8d3e-e3178522f514"`
	LessonTitle       string `json:"lesson_title" example:"Effective Eureka"`
	LessonThumbnail   string `json:"lesson_thumbnail" example:"https://effective-eureka.s3.amazonaws.com/courses/effective-eureka/thumbnail.png"`
	LessonDescription string `json:"lesson_description" example:"Effective Eureka is a lesson about Go."`
	LessonPublished   bool   `json:"lesson_published" example:"false"`
}

func NewHttpLessonCreated(lesson domain.Lesson) *HttpLessonCreated {
	return &HttpLessonCreated{
		LessonID:          lesson.GetLessonID(),
		LessonTitle:       lesson.GetTitle(),
		LessonThumbnail:   lesson.GetThumbnail(),
		LessonDescription: lesson.GetDescription(),
		LessonPublished:   lesson.IsPublished(),
	}
}

// HttpLessonOk is a struct to manage the CourseOk response model.
type HttpLessonOk struct {
	LessonID          string `json:"lesson_id" example:"05d4d9d3-01a3-4fd3-8d3e-e3178522f514"`
	LessonTitle       string `json:"lesson_title" example:"Effective Eureka"`
	LessonThumbnail   string `json:"lesson_thumbnail" example:"https://effective-eureka.s3.amazonaws.com/courses/effective-eureka/thumbnail.png"`
	LessonDescription string `json:"lesson_description" example:"Effective Eureka is a lesson about Go."`
	LessonPublished   bool   `json:"lesson_published" example:"false"`
}

func NewHttpLessonOk(lesson domain.Lesson) *HttpLessonOk {
	return &HttpLessonOk{
		LessonID:          lesson.GetLessonID(),
		LessonTitle:       lesson.GetTitle(),
		LessonThumbnail:   lesson.GetThumbnail(),
		LessonDescription: lesson.GetDescription(),
		LessonPublished:   lesson.IsPublished(),
	}
}

type HttpLessonsWithMeta struct {
	Data []*HttpLessonOk `json:"data"`
	Meta shared.HttpMeta `json:"meta"`
}

func NewHttpLessonsWithMeta(lessons []*HttpLessonOk, params *domain.GetLessonsInSectionParams) *HttpLessonsWithMeta {
	baseURL := fmt.Sprintf("http://localhost:8080/section/%s/lessons", params.SectionID)

	return &HttpLessonsWithMeta{
		Data: lessons,
		Meta: shared.GetMeta(baseURL, params.Page, params.ItemsPerPage, len(lessons)),
	}
}
