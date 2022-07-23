package handler

import "github.com/jeanmolossi/effective-eureka/src/core/shared"

// Common errors

// HttpBadRequestErr is a struct to manage internal server errors.
type HttpBadRequestErr struct {
	Err    string              `json:"error" example:"Bad Request"`
	Errors []shared.FieldError `json:"errors"`
}
