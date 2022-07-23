// Package domain is a package to manage the application domain.
//
// That package is the module domain.
package domain

type Module interface {
	// GetCourseID returns parent course ID.
	GetCourseID() string
	// GetModuleID returns module ID.
	GetModuleID() string
	// GetModuleTitle returns module title.
	GetModuleTitle() string
	// GetModuleDescription returns module description.
	GetModuleDescription() string
	// GetModuleThumb returns module thumb.
	GetModuleThumb() string
	// IsModulePublished returns true if module is published.
	IsModulePublished() bool

	// SetCourseID sets parent course ID.
	SetCourseID(courseID string)
	// SetModuleID sets module ID.
	SetModuleID(moduleID string)
	// SetModuleTitle sets module title.
	SetModuleTitle(title string)
	// SetModuleDescription sets module description.
	SetModuleDescription(desc string)
	// SetModuleThumb sets module thumb.
	SetModuleThumb(thumb string)
	// PublishModule publishes module.
	PublishModule()
	// UnpublishModule unpublishes module.
	UnpublishModule()
}

// GetModuleByID is a interface who provides methods to get a module by ID.
type GetModuleFromCourse interface {
	// Run is the method to get a module by ID.
	Run(courseID string) ([]Module, error)
}

// GetModuleByID is a interface who provides methods to get a module by ID.
type GetModuleByID interface {
	// Run is the method to get a module by ID.
	Run(moduleID string) (Module, error)
}

// CreateModuleInCourse is interface segregation to create a module in course.
type CreateModuleInCourse interface {
	// Run is the method with handles application to create a module in course.
	Run(module Module) (Module, error)
}

// EditModuleInfo is interface segregation to edit a module info.
type EditModuleInfo interface {
	// Run is the method with handles application to edit a module info.
	Run(module Module) (Module, error)
}

// ModuleUpdater is a interface who provides methods to update a module. It
// works with together the ModuleRepository.Edit method.
// We can handle all module properties inside that callback function.
type ModuleUpdater func(module Module) (Module, error)

// ModuleRepository is the interface to manage modules on database.
// To persist and handle saved modules we should implement that interface.
type ModuleRepository interface {
	// GetByID returns a module by ID.
	GetByID(moduleID string) (Module, error)
	// GetByCourseID returns a list of modules by course ID.
	GetByCourseID(courseID string) ([]Module, error)
	// Create creates a new module.
	Create(module Module) (Module, error)
	// Edit updates a module.
	Edit(module Module) (Module, error)
}
