package input

import "errors"

type EditLessonInfo struct {
	LessonID    string `json:"-" param:"lessonID" validate:"required"`
	SectionID   string `json:"section_id" example:"affec47d-e496-48ed-a6d8-78a57177a752"`
	Title       string `json:"title" validate:"required,max=255" example:"Lesson 1"`
	Description string `json:"description" validate:"max=255" example:"Lesson 1 description"`
	Thumbnail   string `json:"thumbnail" example:"https://example.com/thumbnail.png"`
	Published   bool   `json:"published" example:"true"`
	Index       uint16 `json:"index" example:"1"`
}

func (c *EditLessonInfo) GetErrorMap() map[string]map[string]error {
	return map[string]map[string]error{
		"lessonid": {
			"required": errors.New("lessonID is required"),
		},
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
