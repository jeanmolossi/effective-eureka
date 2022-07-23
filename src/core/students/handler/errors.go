package handler

import "github.com/jeanmolossi/effective-eureka/src/core/shared"

// HttpCreateStudentBadRequestErr is a struct to modeling the error response when the request is bad.
type HttpCreateStudentBadRequestErr struct {
	Err    string              `json:"error" example:"Bad Request"`
	Errors []shared.FieldError `json:"errors"`
}

// HttpStudentInternalServerErr is a struct to modeling the error response when the server error.
type HttpStudentInternalServerErr struct {
	Err string `json:"error" example:"Internal Server Error"`
}

// HttpStudentForbiddenErr is a struct to modeling the error response when the request is forbidden.
type HttpStudentForbiddenErr struct {
	Err string `json:"error" example:"Missing authentication"`
}
