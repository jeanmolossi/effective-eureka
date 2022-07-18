package auth

import "errors"

type LoginCredentials struct {
	Username string `json:"username" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=255"`
}

func (l *LoginCredentials) GetErrorMap() map[string]map[string]error {
	return map[string]map[string]error{
		"username": {
			"required": errors.New("username is required"),
			"email":    errors.New("username must be an email"),
		},
		"password": {
			"required": errors.New("password is required"),
			"min":      errors.New("password must be at least 6 characters"),
			"max":      errors.New("password must be at most 255 characters"),
		},
	}
}
