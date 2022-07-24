// Package factory is used to create modules providing the necessary dependencies.
package factory

import (
	"time"

	"github.com/jeanmolossi/effective-eureka/src/core/sections/domain"
)

type Section interface {
	CreateSection(moduleID, title string, index uint16, published bool, createdAt, updatedAt *time.Time) Section
	WithID(sectionID string) Section
	WithCourseID(courseID string) Section
	Build() domain.Section
}

type section struct {
	domain.Section
}

func NewSection() Section {
	return &section{
		domain.NewSection("", "", "", 0, false, nil, nil),
	}
}

func (s *section) CreateSection(moduleID, title string, index uint16, published bool, createdAt, updatedAt *time.Time) Section {
	if createdAt != nil && updatedAt != nil {
		s.Section = domain.NewSection(
			moduleID,
			"",
			title,
			index,
			published,
			createdAt,
			updatedAt,
		)
	} else {
		s.Section.SetModuleID(moduleID)
		s.Section.SetTitle(title)
		s.Section.SetIndex(index)

		if published {
			s.Section.Publish()
		} else {
			s.Section.Unpublish()
		}
	}

	return s
}

func (s *section) WithID(sectionID string) Section {
	s.Section.SetSectionID(sectionID)
	return s
}

func (s *section) WithCourseID(courseID string) Section {
	s.Section.SetCourseID(courseID)
	return s
}

func (s *section) Build() domain.Section {
	return s.Section
}
