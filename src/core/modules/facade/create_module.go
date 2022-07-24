package facade

import (
	"github.com/jeanmolossi/effective-eureka/src/core/modules/domain"
	"github.com/jeanmolossi/effective-eureka/src/core/modules/factory"
	"github.com/jeanmolossi/effective-eureka/src/core/modules/handler"
	"github.com/jeanmolossi/effective-eureka/src/core/modules/input"
	"github.com/jeanmolossi/effective-eureka/src/core/modules/repository"
	"github.com/jeanmolossi/effective-eureka/src/core/modules/usecase"
	"github.com/jeanmolossi/effective-eureka/src/pkg/logger"
	"gorm.io/gorm"
)

type CreateModule interface {
	Input() *input.CreateModule
	Run() (*handler.HttpModuleCreated, error)
}

type createModule struct {
	input *input.CreateModule

	create domain.CreateModuleInCourse

	logger logger.Logger
}

func NewCreateModule(db *gorm.DB) CreateModule {
	repo := repository.NewRepository(db)
	create := usecase.NewCreateModuleInCourse(repo)

	return &createModule{
		new(input.CreateModule),

		create,

		logger.NewLogger(),
	}
}

func (c *createModule) Input() *input.CreateModule {
	return c.input
}

func (c *createModule) Run() (*handler.HttpModuleCreated, error) {
	module := factory.NewModule().CreateModule(
		c.input.CourseID,
		c.input.Title,
		c.input.Thumbnail,
		c.input.Description,
		c.input.Published,
	)

	newModule, err := c.create.Run(module.Build())
	if err != nil {

		return nil, err
	}

	return handler.NewHttpModuleCreated(newModule), nil
}
