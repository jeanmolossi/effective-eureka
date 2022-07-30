package handler

import (
	"fmt"

	"github.com/jeanmolossi/effective-eureka/src/core/modules/domain"
	"github.com/jeanmolossi/effective-eureka/src/core/shared"
)

// HttpOkWithMessage returns a 201 response with a message
type HttpModuleCreated struct {
	CourseID          string `json:"course_id" example:"05d4d9d3-01a3-4fd3-8d3e-e3178522f515"`
	ModuleID          string `json:"module_id" example:"4aa77560-9c90-4128-b308-ad5c0515b5d7"`
	ModuleTitle       string `json:"module_title" example:"Effective Eureka"`
	ModuleDescription string `json:"module_description" example:"Effective Eureka is a course about Go."`
	ModuleThumbnail   string `json:"module_thumbnail" example:"https://effective-eureka.s3.amazonaws.com/courses/effective-eureka/thumbnail.png"`
	ModulePublished   bool   `json:"module_published" example:"false"`
}

// NewHttpModuleCreated returns a new HttpModuleCreated
func NewHttpModuleCreated(module domain.Module) *HttpModuleCreated {
	return &HttpModuleCreated{
		CourseID:          module.GetCourseID(),
		ModuleID:          module.GetModuleID(),
		ModuleTitle:       module.GetModuleTitle(),
		ModuleDescription: module.GetModuleDescription(),
		ModuleThumbnail:   module.GetModuleThumb(),
		ModulePublished:   module.IsModulePublished(),
	}
}

type HttpModuleOk struct {
	CourseID          string `json:"course_id,omitempty" example:"05d4d9d3-01a3-4fd3-8d3e-e3178522f515"`
	ModuleID          string `json:"module_id,omitempty" example:"4aa77560-9c90-4128-b308-ad5c0515b5d7"`
	ModuleTitle       string `json:"module_title,omitempty" example:"Effective Eureka"`
	ModuleDescription string `json:"module_description,omitempty" example:"Effective Eureka is a course about Go."`
	ModuleThumbnail   string `json:"module_thumbnail,omitempty" example:"https://effective-eureka.s3.amazonaws.com/courses/effective-eureka/thumbnail.png"`
	ModulePublished   bool   `json:"module_published,omitempty" example:"false"`
}

func NewHttpModuleOk(module domain.Module) *HttpModuleOk {
	return &HttpModuleOk{
		CourseID:          module.GetCourseID(),
		ModuleID:          module.GetModuleID(),
		ModuleTitle:       module.GetModuleTitle(),
		ModuleDescription: module.GetModuleDescription(),
		ModuleThumbnail:   module.GetModuleThumb(),
		ModulePublished:   module.IsModulePublished(),
	}
}

type HttpModulesWithMeta struct {
	Data []*HttpModuleOk `json:"data"`
	Meta shared.HttpMeta `json:"meta"`
}

func NewHttpModulesWithMeta(modules []*HttpModuleOk, input *domain.GetModulesParams) *HttpModulesWithMeta {
	baseURL := fmt.Sprintf("http://localhost:8080/course/%s/modules", input.CourseID)

	return &HttpModulesWithMeta{
		Data: modules,
		Meta: shared.GetMeta(baseURL, input.Page, input.ItemsPerPage, len(modules)),
	}
}
