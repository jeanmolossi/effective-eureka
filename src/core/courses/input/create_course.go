// Package input represents all accepted inputs for the create course endpoint.
package input

import "errors"

// CreateCourseRepresents the accepted input for the create course endpoint.
type CreateCourse struct {
	Title       string `json:"title" validate:"required,max=255" example:"Effective Eureka"`
	Description string `json:"description" validate:"max=255" example:"This is a catalog video manager API."`
	Thumbnail   string `json:"thumbnail" example:"https://effective-eureka.s3.amazonaws.com/courses/thumbnail/1.jpg"`
	Published   bool   `json:"published" example:"true"`
}

func (c *CreateCourse) GetErrorMap() map[string]map[string]error {
	return map[string]map[string]error{
		"title": {
			"required": errors.New("title is required"),
			"max":      errors.New("title must be less than 255 characters"),
		},
		"description": {
			"max": errors.New("description must be less than 255 characters"),
		},
	}
}
