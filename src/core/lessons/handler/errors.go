package handler

import (
	"github.com/jeanmolossi/effective-eureka/src/core/shared"
)

// Common Errors

// HttpLessonByIDBadRequestErr is a struct to modeling the error response when the request is bad.
type HttpLessonByIDBadRequestErr struct {
	Err string `json:"error" example:"Missing lesson_id param"`
}

// HttpLessonNotFoundErr is a struct to modeling the error response when the lesson is not found.
type HttpLessonNotFoundErr struct {
	Err string `json:"error" example:"Lesson Not Found"`
}

// CreateLesson Errors

// HttpCreateLessonBadRequestErr is a struct to modeling the error response when the request is bad.
type HttpCreateLessonBadRequestErr struct {
	Err    string              `json:"error" example:"Bad Request"`
	Errors []shared.FieldError `json:"errors"`
}

// EditLessonInfo Errors

// HttpEditLessonInfoBadRequestErr is a struct to modeling the error response when the request is bad.
type HttpEditLessonInfoBadRequestErr struct {
	Err    string              `json:"error" example:"Bad Request"`
	Errors []shared.FieldError `json:"errors"`
}
