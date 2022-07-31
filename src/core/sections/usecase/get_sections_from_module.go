package usecase

import (
	"github.com/jeanmolossi/effective-eureka/src/core/sections/domain"
	ormcondition "github.com/jeanmolossi/effective-eureka/src/pkg/orm_condition"
	"github.com/jeanmolossi/effective-eureka/src/pkg/paginator"
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
	pagination := paginator.NewPaginator()

	filters := ormcondition.NewFilterConditions()
	filters.AddCondition("section_published", true)

	if input != nil {
		pagination.SetItemsPerPage(input.ItemsPerPage)
		pagination.SetPage(input.Page)

		filters.AddCondition("module_id", input.ModuleID)
		filters.AddFields(input.Fields)

		if input.NotPublished {
			filters.RemoveCondition("section_published")
		}
	}

	return g.repo.GetByModuleID(filters, pagination)
}
