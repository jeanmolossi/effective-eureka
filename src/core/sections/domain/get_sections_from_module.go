package domain

import "errors"

type GetSectionsParams struct {
	ModuleID     string   `param:"module_id" validate:"required"`
	Fields       []string `query:"fields"`
	NotPublished bool     `query:"not_published"`
	Page         uint16   `query:"page"`
	ItemsPerPage int      `query:"items_per_page"`
}

func (g *GetSectionsParams) GetErrorMap() map[string]map[string]error {
	return map[string]map[string]error{
		"moduleid": {
			"required": errors.New("module_id is required"),
		},
	}
}

// GetSectionsFromModule returns the sections from a module
type GetSectionsFromModule interface {
	// Run execute usecase to get sections from module
	Run(moduleID string) ([]Section, error)
}
