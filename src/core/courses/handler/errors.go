package handler

import "github.com/jeanmolossi/effective-eureka/src/core/shared"

// HttpCreateCourseBadRequestErr is a struct to modeling the error response when the request is bad.
type HttpCreateCourseBadRequestErr struct {
	Err    string              `json:"error" example:"Bad Request"`
	Errors []shared.FieldError `json:"errors"`
}
