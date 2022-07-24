package input

import "errors"

// EditSection represents a request to edit a section input
type EditSection struct {
	SectionID string `json:"-" param:"sectionID" validate:"required"`
	ModuleID  string `json:"module_id" example:"4aa77560-9c90-4128-b308-ad5c0515b5d7"`
	Title     string `json:"title" validate:"required,max=255" example:"Effective Eureka"`
	Index     uint16 `json:"index" example:"1"`
	Published bool   `json:"published" example:"true"`
}

func (e *EditSection) GetErrorMap() map[string]map[string]error {
	return map[string]map[string]error{
		"sectionid": {
			"required": errors.New("section_id is required"),
		},
		"title": {
			"required": errors.New("title is required"),
			"max":      errors.New("title must be less than 255 characters"),
		},
	}
}
