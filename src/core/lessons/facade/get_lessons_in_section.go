package facade

import (
	"github.com/jeanmolossi/effective-eureka/src/core/lessons/domain"
	"github.com/jeanmolossi/effective-eureka/src/core/lessons/handler"
	"github.com/jeanmolossi/effective-eureka/src/core/lessons/repository"
	"github.com/jeanmolossi/effective-eureka/src/core/lessons/usecase"
	"github.com/jeanmolossi/effective-eureka/src/pkg/logger"
	"gorm.io/gorm"
)

type GetLessonsInSection interface {
	Run(params *domain.GetLessonsInSectionParams) (*handler.HttpLessonsWithMeta, error)
}

type getLessonsInSection struct {
	getLessons domain.GetLessonsInSection

	logger logger.Logger
}

func NewGetLessonsInSection(db *gorm.DB) GetLessonsInSection {
	repo := repository.NewRepository(db)
	getLessons := usecase.NewGetLessonsInSection(repo)

	return &getLessonsInSection{
		getLessons,

		logger.NewLogger(),
	}
}

func (g *getLessonsInSection) Run(params *domain.GetLessonsInSectionParams) (*handler.HttpLessonsWithMeta, error) {
	lessons, err := g.getLessons.Run(params)
	if err != nil {
		return nil, err
	}

	httpLessons := make([]*handler.HttpLessonOk, len(lessons))
	for i, lesson := range lessons {
		httpLessons[i] = handler.NewHttpLessonOk(lesson)
	}

	return handler.NewHttpLessonsWithMeta(httpLessons, params), nil
}
