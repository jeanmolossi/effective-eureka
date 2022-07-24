package handler

import (
	"errors"
	"net/http"

	"github.com/jeanmolossi/effective-eureka/src/cmd/httputil"
	"github.com/jeanmolossi/effective-eureka/src/core/shared"
	"github.com/labstack/echo/v4"
)

// Common Errors

// HttpCourseByIDBadRequestErr is a struct to modeling the error response when the request is bad.
type HttpCourseByIDBadRequestErr struct {
	Err string `json:"error" example:"Missing course_id param"`
}

// HttpCourseNotFoundErr is a struct to modeling the error response when the course is not found.
type HttpCourseNotFoundErr struct {
	Err string `json:"error" example:"Course Not Found"`
}

// CreateCourse Errors

// HttpCreateCourseBadRequestErr is a struct to modeling the error response when the request is bad.
type HttpCreateCourseBadRequestErr struct {
	Err    string              `json:"error" example:"Bad Request"`
	Errors []shared.FieldError `json:"errors"`
}

// EditCourseInfo Errors

// HttpEditCourseInfoBadRequestErr is a struct to modeling the error response when the request is bad.
type HttpEditCourseInfoBadRequestErr struct {
	Err    string              `json:"error" example:"Bad Request"`
	Errors []shared.FieldError `json:"errors"`
}

func ErrorHandler(c echo.Context, err error) error {
	var notFoundErr *shared.NotFoundErr
	var badRequestErr *shared.ValidationErr
	var echoBindErr *echo.BindingError
	var unauthorizedErr *shared.UnauthorizedErr

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
