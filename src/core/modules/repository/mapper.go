package repository

import "github.com/jeanmolossi/effective-eureka/src/core/modules/domain"

func DomainToModel(module domain.Module) *ModuleModel {
	return &ModuleModel{
		CourseID:          module.GetCourseID(),
		ModuleTitle:       module.GetModuleTitle(),
		ModuleThumb:       module.GetModuleThumb(),
		ModuleDescription: module.GetModuleDescription(),
		ModulePublished:   module.IsModulePublished(),
	}
}

func ModelToDomain(model *ModuleModel) domain.Module {
	if model == nil {
		return domain.NewModule("", "", "", "", false)
	}

	module := domain.NewModule(
		model.CourseID,
		model.ModuleTitle,
		model.ModuleThumb,
		model.ModuleDescription,
		model.ModulePublished,
	)
	module.SetModuleID(model.ModuleID)

	return module
}
