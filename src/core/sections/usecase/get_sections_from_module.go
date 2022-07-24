package usecase

import "github.com/jeanmolossi/effective-eureka/src/core/sections/domain"

type getSectionsFromModule struct {
	repo domain.SectionsRepository
}

func NewGetSectionsFromModule(repo domain.SectionsRepository) domain.GetSectionsFromModule {
	return &getSectionsFromModule{
		repo: repo,
	}
}

// Run execute usecase to get sections from module
func (g *getSectionsFromModule) Run(moduleID string) ([]domain.Section, error) {
	return g.repo.GetByModuleID(moduleID)
}
