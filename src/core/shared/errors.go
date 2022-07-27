package shared

import (
	"errors"
	"net/http"

	"github.com/jeanmolossi/effective-eureka/src/cmd/httputil"
	"github.com/labstack/echo/v4"
)

type NotFoundErr struct {
	// Range: 0 through 65535.
	Code uint16 `json:"-"`
	// The error message.
	Message string `json:"error" example:"Not Found"`
}

// Error returns the error message.
func (e *NotFoundErr) Error() string {
	return e.Message
}

func NewNotFoundErr(err error) *NotFoundErr {
	return &NotFoundErr{
		Code:    http.StatusNotFound,
		Message: err.Error(),
	}
}

type UnauthorizedErr struct {
	// Range: 0 through 65535.
	Code uint16 `json:"-"`
	// The error message.
	Message string `json:"error" example:"Unauthorized"`
}

// Error returns the error message.
func (e *UnauthorizedErr) Error() string {
	return e.Message
}

func NewUnauthorizedErr(err error) *UnauthorizedErr {
	return &UnauthorizedErr{
		Code:    http.StatusUnauthorized,
		Message: err.Error(),
	}
}

type BadRequestErr struct {
	// Range: 0 through 65535.
	Code uint16 `json:"-"`
	// The error message.
	Message string `json:"error" example:"Bad Request"`
}

// Error returns the error message.
func (e *BadRequestErr) Error() string {
	return e.Message
}

func NewBadRequestErr(err error) *BadRequestErr {
	return &BadRequestErr{
		Code:    http.StatusBadRequest,
		Message: err.Error(),
	}
}

func ErrorHandler(c echo.Context, err error) error {
	var notFoundErr *NotFoundErr
	var badRequestErr *ValidationErr
	var unauthorizedErr *UnauthorizedErr
	var echoBindErr *echo.BindingError

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
