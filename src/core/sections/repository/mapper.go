package repository

import "github.com/jeanmolossi/effective-eureka/src/core/sections/domain"

func ModelToDomain(model *SectionModel) domain.Section {
	if model == nil {
		return domain.NewSection("", "", "", 0, false, nil, nil)
	}

	section := domain.NewSection(
		model.ModuleID,
		model.CourseID,
		model.SectionTitle,
		model.SectionIndex,
		model.SectionPublished,
		&model.SectionCreatedAt,
		&model.SectionUpdatedAt,
	)
	section.SetSectionID(model.SectionID)

	return section
}

func DomainToModel(section domain.Section) *SectionModel {
	createdAt, updatedAt := section.GetTimestamps()

	return &SectionModel{
		ModuleID:         section.GetModuleID(),
		CourseID:         section.GetCourseID(),
		SectionID:        section.GetSectionID(),
		SectionIndex:     section.GetIndex(),
		SectionTitle:     section.GetTitle(),
		SectionPublished: section.IsPublished(),
		SectionCreatedAt: *createdAt,
		SectionUpdatedAt: *updatedAt,
	}
}
