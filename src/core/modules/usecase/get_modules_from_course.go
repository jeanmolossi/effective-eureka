package usecase

import (
	"github.com/jeanmolossi/effective-eureka/src/core/modules/domain"
	"github.com/jeanmolossi/effective-eureka/src/core/shared"
)

type getModulesFromCourse struct {
	repo domain.ModuleRepository
}

func NewGetModuleFromCourse(repo domain.ModuleRepository) domain.GetModuleFromCourse {
	return &getModulesFromCourse{repo}
}

// Run is the method to get a module by ID.
func (g *getModulesFromCourse) Run(params *domain.GetModulesParams) ([]domain.Module, error) {
	filters := shared.Filters{
		ConditionMap: map[string]interface{}{
			"module_published": true,
		},
	}

	if params != nil {
		filters.Fields = params.Fields

		if params.NotPublished {
			filters.ConditionMap = nil
		}

		filters.ConditionMap = map[string]interface{}{
			"course_id": params.CourseID,
		}
	}

	paginator := shared.PagesConfig{
		Page:         1,
		ItemsPerPage: 10,
	}

	if params.Page > 0 {
		paginator.Page = params.Page
	}

	if params.ItemsPerPage > 0 {
		paginator.ItemsPerPage = params.ItemsPerPage
	}

	return g.repo.GetByCourseID(&filters, &paginator)
}
