package usecase

import (
	"errors"

	"github.com/jeanmolossi/effective-eureka/src/core/modules/domain"
)

type editModuleInfo struct {
	repo domain.ModuleRepository
}

func NewEditModuleInfo(repo domain.ModuleRepository) domain.EditModuleInfo {
	return &editModuleInfo{repo}
}

func (e *editModuleInfo) Run(module domain.Module) (domain.Module, error) {
	if !e.repo.IssetCourseID(module.GetCourseID()) {
		return nil, domain.NewNotFoundErr(errors.New("course not found"))
	}

	return e.repo.Edit(module.GetModuleID(), e.updater(module))
}

func (e *editModuleInfo) updater(newModule domain.Module) domain.ModuleUpdater {
	return func(currentModule domain.Module) (domain.Module, error) {
		if newModule.GetCourseID() != "" {
			currentModule.SetCourseID(newModule.GetCourseID())
		}

		if newModule.GetModuleTitle() != "" {
			currentModule.SetModuleTitle(newModule.GetModuleTitle())
		}

		if newModule.GetModuleThumb() != "" {
			currentModule.SetModuleThumb(newModule.GetModuleThumb())
		}

		if newModule.GetModuleDescription() != "" {
			currentModule.SetModuleDescription(newModule.GetModuleDescription())
		}

		if newModule.IsModulePublished() {
			currentModule.PublishModule()
		} else {
			currentModule.UnpublishModule()
		}

		return currentModule, nil
	}
}
