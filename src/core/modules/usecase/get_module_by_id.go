package usecase

import "github.com/jeanmolossi/effective-eureka/src/core/modules/domain"

// getModelByID is a usecase that gets a module by ID
type getModuleByID struct {
	repo domain.ModuleRepository
}

// NewGetModuleByID returns a new getModuleByID struct wich implements
// the domain.GetModelByID interface
func NewGetModuleByID(repo domain.ModuleRepository) domain.GetModuleByID {
	return &getModuleByID{repo}
}

// Run runs the usecase
func (g *getModuleByID) Run(moduleID string) (domain.Module, error) {
	return g.repo.GetByID(moduleID)
}
