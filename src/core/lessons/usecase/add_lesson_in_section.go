package usecase

import "github.com/jeanmolossi/effective-eureka/src/core/lessons/domain"

type addLessonInSection struct {
	repo domain.LessonRepository
}

func NewAddLessonInSection(repo domain.LessonRepository) domain.AddLessonInSection {
	return &addLessonInSection{
		repo: repo,
	}
}

func (a *addLessonInSection) AddLesson(lesson domain.Lesson) (domain.Lesson, error) {
	return a.repo.Create(lesson)
}
