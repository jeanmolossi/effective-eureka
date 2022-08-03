package domain

import "errors"

type GetLessonsInSectionParams struct {
	SectionID    string   `param:"sectionID" validate:"required"`
	Fields       []string `query:"fields"`
	NotPublished bool     `query:"not_published"`
	Page         uint16   `query:"page"`
	ItemsPerPage int      `query:"items_per_page"`
}

func (c *GetLessonsInSectionParams) GetErrorMap() map[string]map[string]error {
	return map[string]map[string]error{
		"sectionid": {
			"required": errors.New("section id is required"),
		},
	}
}

type GetLessonsInSection interface {
	Run(params *GetLessonsInSectionParams) ([]Lesson, error)
}
