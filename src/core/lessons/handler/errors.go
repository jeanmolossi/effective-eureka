package handler

import (
	"errors"
	"net/http"

	"github.com/jeanmolossi/effective-eureka/src/cmd/httputil"
	"github.com/jeanmolossi/effective-eureka/src/core/lessons/domain"
	"github.com/jeanmolossi/effective-eureka/src/core/shared"
	"github.com/labstack/echo/v4"
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

func ErrorHandler(c echo.Context, err error) error {
	var notFoundErr *domain.NotFoundErr
	var badRequestErr *shared.ValidationErr
	var echoBindErr *echo.BindingError
	var unauthorizedErr *domain.UnauthorizedErr

	switch {
	case errors.As(err, &badRequestErr):
		return c.JSON(http.StatusBadRequest, badRequestErr)
	case errors.As(err, &echoBindErr):
		return c.JSON(http.StatusBadRequest, echoBindErr)
	case errors.As(err, &notFoundErr):
		return c.JSON(int(notFoundErr.Code), notFoundErr)
	case errors.As(err, &unauthorizedErr):
		return c.JSON(int(unauthorizedErr.Code), unauthorizedErr)
	default:
		return c.JSON(http.StatusInternalServerError, httputil.HttpInternalServerErr{
			Message: err.Error(),
		})
	}
}
