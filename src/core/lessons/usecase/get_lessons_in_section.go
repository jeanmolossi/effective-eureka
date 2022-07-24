package usecase

import (
	"github.com/jeanmolossi/effective-eureka/src/core/lessons/domain"
)

type getLessonsInSection struct {
	repo domain.LessonRepository
}

func NewGetLessonsInSection(repo domain.LessonRepository) domain.GetLessonsInSection {
	return &getLessonsInSection{repo}
}

func (g *getLessonsInSection) Run(sectionID string) ([]domain.Lesson, error) {
	lessons, err := g.repo.GetLessonsFromSection(sectionID)
	if err != nil {
		return nil, err
	}

	return lessons, nil
}
