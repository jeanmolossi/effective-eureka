package handler

import "github.com/jeanmolossi/effective-eureka/src/core/sections/domain"

// HttpSectionCreated returns a 201 response with a message
type HttpSectionCreated struct {
	CourseID         string `json:"course_id" example:"05d4d9d3-01a3-4fd3-8d3e-e3178522f515"`
	ModuleID         string `json:"module_id" example:"4aa77560-9c90-4128-b308-ad5c0515b5d7"`
	SectionID        string `json:"section_id" example:"4aa77560-9c90-4128-b308-ad5c0515b5d7"`
	SectionTitle     string `json:"section_title" example:"Effective Eureka"`
	SectionIndex     uint16 `json:"section_index" example:"1"`
	SectionPublished bool   `json:"section_published" example:"false"`
}

// NewHttpModuleCreated returns a new HttpSectionCreated
func NewHttpModuleCreated(section domain.Section) *HttpSectionCreated {
	return &HttpSectionCreated{
		CourseID:         section.GetCourseID(),
		ModuleID:         section.GetModuleID(),
		SectionID:        section.GetSectionID(),
		SectionTitle:     section.GetTitle(),
		SectionIndex:     section.GetIndex(),
		SectionPublished: section.IsPublished(),
	}
}
