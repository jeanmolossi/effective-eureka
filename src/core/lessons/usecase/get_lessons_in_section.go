package usecase

import (
	"errors"

	"github.com/jeanmolossi/effective-eureka/src/core/lessons/domain"
	ormcondition "github.com/jeanmolossi/effective-eureka/src/pkg/orm_condition"
	"github.com/jeanmolossi/effective-eureka/src/pkg/paginator"
)

type getLessonsInSection struct {
	repo domain.LessonRepository
}

func NewGetLessonsInSection(repo domain.LessonRepository) domain.GetLessonsInSection {
	return &getLessonsInSection{repo}
}

func (g *getLessonsInSection) Run(params *domain.GetLessonsInSectionParams) ([]domain.Lesson, error) {
	if params == nil {
		return nil, domain.NewBadRequestErr(
			errors.New("you should provide search params"),
		)
	}

	filters := ormcondition.NewFilterConditions()
	filters.AddCondition("section_id", params.SectionID)
	filters.AddFields(params.Fields)

	pagination := paginator.NewPaginator()
	pagination.SetPage(params.Page)
	pagination.SetItemsPerPage(params.ItemsPerPage)

	lessons, err := g.repo.GetLessonsFromSection(filters, pagination)
	if err != nil {
		return nil, err
	}

	return lessons, nil
}
