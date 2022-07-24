package domain

import "net/http"

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
