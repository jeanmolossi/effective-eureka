package handler

import (
	"fmt"

	"github.com/jeanmolossi/effective-eureka/src/core/lessons/domain"
	"github.com/jeanmolossi/effective-eureka/src/core/shared"
)

// HttpLessonCreated is a struct to manage the LessonCreated response model.
type HttpLessonCreated struct {
	ID           string `json:"id" example:"05d4d9d3-01a3-4fd3-8d3e-e3178522f514"`
	Title        string `json:"title" example:"Effective Eureka"`
	Thumbnail    string `json:"thumbnail" example:"https://effective-eureka.s3.amazonaws.com/courses/effective-eureka/thumbnail.png"`
	VideoPreview string `json:"video_preview" example:"https://effective-eureka.s3.amazonaws.com/courses/effective-eureka/videoPreview.png"`
	Video        string `json:"video" example:"https://effective-eureka.s3.amazonaws.com/courses/effective-eureka/video.mp4"`
	Description  string `json:"description" example:"Effective Eureka is a lesson about Go."`
	Published    bool   `json:"published" example:"false"`
}

func NewHttpLessonCreated(lesson domain.Lesson) *HttpLessonCreated {
	return &HttpLessonCreated{
		ID:           lesson.GetLessonID(),
		Title:        lesson.GetTitle(),
		Thumbnail:    lesson.GetThumbnail(),
		VideoPreview: lesson.GetVideoPreview(),
		Video:        lesson.GetVideo(),
		Description:  lesson.GetDescription(),
		Published:    lesson.IsPublished(),
	}
}

// HttpLessonOk is a struct to manage the CourseOk response model.
type HttpLessonOk struct {
	ID           string `json:"id,omitempty" example:"05d4d9d3-01a3-4fd3-8d3e-e3178522f514"`
	Title        string `json:"title,omitempty" example:"Effective Eureka"`
	Thumbnail    string `json:"thumbnail,omitempty" example:"https://effective-eureka.s3.amazonaws.com/courses/effective-eureka/thumbnail.png"`
	VideoPreview string `json:"video_preview,omitempty" example:"https://effective-eureka.s3.amazonaws.com/courses/effective-eureka/thumbnail.png"`
	Video        string `json:"video,omitempty" example:"https://effective-eureka.s3.amazonaws.com/courses/effective-eureka/thumbnail.png"`
	Description  string `json:"description,omitempty" example:"Effective Eureka is a lesson about Go."`
	Published    bool   `json:"published,omitempty" example:"false"`
}

func NewHttpLessonOk(lesson domain.Lesson) *HttpLessonOk {
	return &HttpLessonOk{
		ID:           lesson.GetLessonID(),
		Title:        lesson.GetTitle(),
		Thumbnail:    lesson.GetThumbnail(),
		VideoPreview: lesson.GetVideoPreview(),
		Video:        lesson.GetVideo(),
		Description:  lesson.GetDescription(),
		Published:    lesson.IsPublished(),
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
