package handler

import "github.com/jeanmolossi/effective-eureka/src/core/modules/domain"

// HttpOkWithMessage returns a 201 response with a message
type HttpModuleCreated struct {
	CourseID          string `json:"course_id" example:"05d4d9d3-01a3-4fd3-8d3e-e3178522f515"`
	ModuleID          string `json:"module_id" example:"05d4d9d3-01a3-4fd3-8d3e-e3178522f514"`
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
