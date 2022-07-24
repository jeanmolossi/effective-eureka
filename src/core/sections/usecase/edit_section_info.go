package usecase

import (
	"github.com/jeanmolossi/effective-eureka/src/core/sections/domain"
)

type editSectionInfo struct {
	repo domain.SectionsRepository
}

func NewEditSectionInfo(repo domain.SectionsRepository) domain.EditSectionInfo {
	return &editSectionInfo{repo}
}

// Run execute usecase to update section in module
func (e *editSectionInfo) Run(section domain.Section) (domain.Section, error) {
	return e.repo.Edit(section, e.updater(section))
}

func (e *editSectionInfo) updater(newSection domain.Section) domain.SectionUpdater {
	return func(currentSection domain.Section) (domain.Section, error) {
		if newSection.GetModuleID() != "" {
			currentSection.SetModuleID(newSection.GetModuleID())
		}

		if newSection.GetTitle() != "" {
			currentSection.SetTitle(newSection.GetTitle())
		}

		if newSection.GetIndex() != 0 {
			currentSection.SetIndex(newSection.GetIndex())
		}

		if newSection.IsPublished() {
			currentSection.Publish()
		} else {
			currentSection.Unpublish()
		}

		return currentSection, nil
	}
}
