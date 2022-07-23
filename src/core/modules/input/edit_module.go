package input

import "errors"

type EditModuleInfo struct {
	ModuleID    string `json:"-" param:"moduleID" validate:"required,uuid4"`
	CourseID    string `json:"course_id,omitempty" validate:"uuid4" example:"f0f8e8c4-8b8f-4d8e-b8e7-8f9e939ca9e8"`
	Title       string `json:"title,omitempty" validate:"max=255" example:"Effective Eureka"`
	Thumbnail   string `json:"thumbnail" example:"https://example.com/thumbnail.png"`
	Description string `json:"description" validate:"max=255" example:"Effective Eureka is a course about effective eureka."`
	Published   bool   `json:"published,omitempty" example:"true"`
}

func (e *EditModuleInfo) GetErrorMap() map[string]map[string]error {
	return map[string]map[string]error{
		"moduleid": {
			"required": errors.New("module_id is required"),
			"uuid4":    errors.New("module_id must be a valid UUID"),
		},
		"courseid": {
			"uuid4": errors.New("course_id must be a valid UUID"),
		},
		"title": {
			"max": errors.New("title must be less than 255 characters"),
		},
		"description": {
			"max": errors.New("description must be less than 255 characters"),
		},
	}
}
