package handler

import (
	"errors"
	"net/http"

	"github.com/jeanmolossi/effective-eureka/src/cmd/httputil"
	"github.com/jeanmolossi/effective-eureka/src/core/modules/domain"
	"github.com/jeanmolossi/effective-eureka/src/core/shared"
	"github.com/labstack/echo/v4"
)

// Common errors

// HttpBadRequestErr is a struct to manage internal server errors.
type HttpBadRequestErr struct {
	Err    string              `json:"error" example:"Bad Request"`
	Errors []shared.FieldError `json:"errors"`
}

func ErrorHandler(c echo.Context, err error) error {
	var notFoundErr *domain.NotFoundErr
	switch {
	case errors.As(err, &notFoundErr):
		return c.JSON(int(notFoundErr.Code), notFoundErr)
	default:
		return c.JSON(http.StatusInternalServerError, httputil.HttpInternalServerErr{
			Message: err.Error(),
		})
	}
}
