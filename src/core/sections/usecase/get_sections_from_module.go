package usecase

import (
	"github.com/jeanmolossi/effective-eureka/src/core/sections/domain"
	"github.com/jeanmolossi/effective-eureka/src/core/shared"
)

type getSectionsFromModule struct {
	repo domain.SectionsRepository
}

func NewGetSectionsFromModule(repo domain.SectionsRepository) domain.GetSectionsFromModule {
	return &getSectionsFromModule{
		repo: repo,
	}
}

// Run execute usecase to get sections from module
func (g *getSectionsFromModule) Run(input *domain.GetSectionsParams) ([]domain.Section, error) {
	filters := shared.Filters{
		ConditionMap: map[string]interface{}{
			"section_published": true,
		},
	}

	if input != nil {
		if input.NotPublished {
			filters.ConditionMap = nil
		}

		filters.ConditionMap = map[string]interface{}{
			"module_id": input.ModuleID,
		}
	}

	pagination := shared.PagesConfig{
		Page:         input.Page,
		ItemsPerPage: input.ItemsPerPage,
	}

	return g.repo.GetByModuleID(&filters, &pagination)
}
