package usecase

import "github.com/jeanmolossi/effective-eureka/src/core/modules/domain"

type getModulesFromCourse struct {
	repo domain.ModuleRepository
}

func NewGetModuleFromCourse(repo domain.ModuleRepository) domain.GetModuleFromCourse {
	return &getModulesFromCourse{repo}
}

// Run is the method to get a module by ID.
func (g *getModulesFromCourse) Run(courseID string) ([]domain.Module, error) {
	return g.repo.GetByCourseID(courseID)
}
