package input

import "errors"

type CreateLesson struct {
	SectionID   string `json:"-" param:"sectionID" validate:"required"`
	Title       string `json:"title" validate:"required,max=255" example:"Lesson 1"`
	Description string `json:"description" validate:"max=255" example:"Lesson 1 description"`
	Thumbnail   string `json:"thumbnail" example:"https://example.com/thumbnail.png"`
	Index       uint16 `json:"index" example:"1"`
	Published   bool   `json:"published" example:"true"`
}

func (c *CreateLesson) GetErrorMap() map[string]map[string]error {
	return map[string]map[string]error{
		"sectionid": {
			"required": errors.New("sectionID is required"),
		},
		"title": {
			"required": errors.New("title is required"),
			"max":      errors.New("title is too long"),
		},
		"description": {
			"max": errors.New("description is too long"),
		},
	}
}
