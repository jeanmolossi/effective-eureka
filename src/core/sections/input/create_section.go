// Package input represents all accepted inputs for the create section endpoint.
package input

import "errors"

// CreateSection represents the accepted input for the create section endpoint.
type CreateSection struct {
	ModuleID  string `json:"-" param:"moduleID" validate:"required"`
	Index     uint16 `json:"index" validate:"required" example:"1"`
	Title     string `json:"title" validate:"required,max=255" example:"Effective Eureka"`
	Published bool   `json:"published" example:"true"`
}

func (c *CreateSection) GetErrorMap() map[string]map[string]error {
	return map[string]map[string]error{
		"moduleid": {
			"required": errors.New("module_id is required"),
		},
		"index": {
			"required": errors.New("index is required"),
		},
		"title": {
			"required": errors.New("title is required"),
			"max":      errors.New("title must be less than 255 characters"),
		},
	}
}
