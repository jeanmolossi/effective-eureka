// Package usecase is a package to manage the application usecase.
//
// That package is the module usecase.
package usecase

import "github.com/jeanmolossi/effective-eureka/src/core/modules/domain"

// createModuleInCourse is a struct to create a module in course.
// It implements domain.CreateModuleInCourse interface.
type createModuleInCourse struct {
	repo domain.ModuleRepository
}

// NewCreateModuleInCourse is a function to create a createModuleInCourse struct.
// That one implements domain.CreateModuleInCourse interface.
func NewCreateModuleInCourse(repo domain.ModuleRepository) domain.CreateModuleInCourse {
	return &createModuleInCourse{repo}
}

// Run is the method with handles application to create a module in course.
func (c *createModuleInCourse) Run(module domain.Module) (domain.Module, error) {
	return c.repo.Create(module)
}
