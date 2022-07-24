package handler

import (
	"errors"
	"net/http"

	"github.com/jeanmolossi/effective-eureka/src/cmd/httputil"
	"github.com/jeanmolossi/effective-eureka/src/core/sections/domain"
	"github.com/jeanmolossi/effective-eureka/src/core/shared"
	"github.com/labstack/echo/v4"
)

// Common errors

// HttpBadRequestErr is a struct to manage internal server errors.
type HttpBadRequestErr struct {
	Err    string              `json:"error" example:"Bad Request"`
	Errors []shared.FieldError `json:"errors"`
}

func (e *HttpBadRequestErr) Error() string {
	return e.Err
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
