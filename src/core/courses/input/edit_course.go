package input

import "errors"

type EditCourseInfo struct {
	Title       string `json:"title" validate:"max=255" example:"Effective Eureka"`
	Thumbnail   string `json:"thumbnail"  example:"https://example.com/thumbnail.png"`
	Description string `json:"description" validate:"max=255" example:"Effective Eureka is a course about effective eureka."`
}

func (e *EditCourseInfo) GetErrorMap() map[string]map[string]error {
	return map[string]map[string]error{
		"title": {
			"max": errors.New("title must be less than 255 characters"),
		},
		"description": {
			"max": errors.New("description must be less than 255 characters"),
		},
	}
}
