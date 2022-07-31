package handler

import (
	"fmt"

	"github.com/jeanmolossi/effective-eureka/src/core/sections/domain"
	"github.com/jeanmolossi/effective-eureka/src/core/shared"
)

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
func NewHttpSectionCreated(section domain.Section) *HttpSectionCreated {
	return &HttpSectionCreated{
		CourseID:         section.GetCourseID(),
		ModuleID:         section.GetModuleID(),
		SectionID:        section.GetSectionID(),
		SectionTitle:     section.GetTitle(),
		SectionIndex:     section.GetIndex(),
		SectionPublished: section.IsPublished(),
	}
}

// HttpSectionOk returns a 201 response with a message
type HttpSectionOk struct {
	CourseID         string `json:"course_id,omitempty" example:"05d4d9d3-01a3-4fd3-8d3e-e3178522f515"`
	ModuleID         string `json:"module_id,omitempty" example:"4aa77560-9c90-4128-b308-ad5c0515b5d7"`
	SectionID        string `json:"section_id,omitempty" example:"b670feb8-35d6-45b1-91c9-8586213b2688"`
	SectionTitle     string `json:"section_title,omitempty" example:"Effective Eureka"`
	SectionIndex     uint16 `json:"section_index,omitempty" example:"1"`
	SectionPublished bool   `json:"section_published,omitempty" example:"false"`
}

// NewHttpSectionOk returns a new HttpSectionOk
func NewHttpSectionOk(section domain.Section) *HttpSectionOk {
	return &HttpSectionOk{
		CourseID:         section.GetCourseID(),
		ModuleID:         section.GetModuleID(),
		SectionID:        section.GetSectionID(),
		SectionTitle:     section.GetTitle(),
		SectionIndex:     section.GetIndex(),
		SectionPublished: section.IsPublished(),
	}
}

type HttpSectionWithMeta struct {
	Data []*HttpSectionOk `json:"data"`
	Meta *shared.HttpMeta `json:"meta"`
}

func NewHttpSectionWithMeta(sections []*HttpSectionOk, input *domain.GetSectionsParams) *HttpSectionWithMeta {
	baseURL := fmt.Sprintf("http://localhost:8080/module/%s/sections", input.ModuleID)
	meta := shared.GetMeta(baseURL, input.Page, input.ItemsPerPage, len(sections))

	return &HttpSectionWithMeta{
		Data: sections,
		Meta: &meta,
	}
}
