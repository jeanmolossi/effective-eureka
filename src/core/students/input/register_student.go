package input

import "errors"

type StudentInfo struct {
	Email    string `json:"username" validate:"required,email" example:"john@doe.com"`
	Password string `json:"password" validate:"required,min=8,max=64" example:"123456789"`
}

func (s *StudentInfo) GetErrorMap() map[string]map[string]error {
	return map[string]map[string]error{
		"username": {
			"required": errors.New("email is required"),
			"email":    errors.New("email is not valid"),
		},
		"password": {
			"required": errors.New("password is required"),
			"min":      errors.New("password must be at least 8 characters"),
			"max":      errors.New("password must be at most 64 characters"),
		},
	}
}
