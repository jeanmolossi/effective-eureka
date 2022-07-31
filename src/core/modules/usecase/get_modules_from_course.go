package usecase

import (
	"github.com/jeanmolossi/effective-eureka/src/core/modules/domain"
	ormcondition "github.com/jeanmolossi/effective-eureka/src/pkg/orm_condition"
	"github.com/jeanmolossi/effective-eureka/src/pkg/paginator"
)

type getModulesFromCourse struct {
	repo domain.ModuleRepository
}

func NewGetModuleFromCourse(repo domain.ModuleRepository) domain.GetModuleFromCourse {
	return &getModulesFromCourse{repo}
}

// Run is the method to get a module by ID.
func (g *getModulesFromCourse) Run(params *domain.GetModulesParams) ([]domain.Module, error) {
	filters := ormcondition.NewFilterConditions()
	filters.AddCondition("module_published", true)

	paginator := paginator.NewPaginator()

	if params != nil {
		filters.AddCondition("course_id", params.CourseID)
		filters.AddFields(params.Fields)

		if params.NotPublished {
			filters.RemoveCondition("module_published")
		}

		paginator.SetPage(params.Page)
		paginator.SetItemsPerPage(params.ItemsPerPage)
	}

	return g.repo.GetByCourseID(filters, paginator)
}
