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
		Code:    http.StatusOK,
		Message: err.Error(),
	}
}
