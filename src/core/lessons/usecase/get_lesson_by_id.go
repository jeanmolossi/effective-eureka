package usecase

import (
	"github.com/jeanmolossi/effective-eureka/src/core/lessons/domain"
	ormcondition "github.com/jeanmolossi/effective-eureka/src/pkg/orm_condition"
)

type getLessonByID struct {
	repo domain.LessonRepository
}

func NewGetLessonByID(repo domain.LessonRepository) domain.GetLesson {
	return &getLessonByID{repo}
}

func (g *getLessonByID) Run(params *domain.GetLessonParams) (domain.Lesson, error) {
	filters := ormcondition.NewFilterConditions()

	if params != nil {
		filters.AddCondition("lesson_id", params.LessonID)
		filters.AddFields(params.Fields)
	}

	return g.repo.GetLesson(filters)
}
