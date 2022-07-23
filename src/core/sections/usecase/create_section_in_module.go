package usecase

import (
	"github.com/jeanmolossi/effective-eureka/src/core/sections/domain"
)

type createSectionInModule struct {
	repo domain.SectionsRepository
}

func NewCreateSectionInModule(repo domain.SectionsRepository) domain.CreateSectionInModule {
	return &createSectionInModule{
		repo: repo,
	}
}

// Run execute usecase to create section in module
func (c *createSectionInModule) Run(section domain.Section) (domain.Section, error) {
	return c.repo.Create(section)
}
