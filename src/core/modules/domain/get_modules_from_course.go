package domain

import "errors"

type GetModulesParams struct {
	CourseID     string   `param:"courseID" validate:"required"`
	Fields       []string `query:"fields"`
	NotPublished bool     `query:"not_published"`
	Page         uint16   `query:"page"`
	ItemsPerPage int      `query:"items_per_page"`
}

func (c *GetModulesParams) GetErrorMap() map[string]map[string]error {
	return map[string]map[string]error{
		"courseid": {
			"required": errors.New("course id is required"),
		},
	}
}

// GetModuleByID is a interface who provides methods to get a module by ID.
type GetModuleFromCourse interface {
	// Run is the method to get a module by ID.
	Run(input *GetModulesParams) ([]Module, error)
}
