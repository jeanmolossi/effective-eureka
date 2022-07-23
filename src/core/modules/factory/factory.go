// Package factory is used to create modules providing the necessary dependencies.
package factory

import "github.com/jeanmolossi/effective-eureka/src/core/modules/domain"

type Module interface {
	CreateModule(courseID, title, thumb, desc string, published bool) Module
	WithID(id string) Module
	Build() domain.Module
}

type module struct {
	domain.Module
}

func NewModule() Module {
	return &module{
		domain.NewModule("", "", "", "", false),
	}
}

func (m *module) CreateModule(courseID, title, thumb, desc string, published bool) Module {
	m.Module.SetCourseID(courseID)
	m.Module.SetModuleTitle(title)
	m.Module.SetModuleThumb(thumb)
	m.Module.SetModuleDescription(desc)

	if published {
		m.Module.PublishModule()
	} else {
		m.Module.UnpublishModule()
	}

	return m
}

func (m *module) WithID(id string) Module {
	return m
}

func (m *module) Build() domain.Module {
	return m.Module
}
