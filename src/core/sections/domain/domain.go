package domain

import "time"

// Section is a domain section
type Section interface {
	// GetModuleID returns the module ID
	GetModuleID() string
	// GetCourseID returns the course ID
	GetCourseID() string
	// GetSectionID returns the section ID
	GetSectionID() string
	// GetIndex returns the section position index
	GetIndex() uint16
	// GetTitle returns the section title
	GetTitle() string
	// IsPublished returns true if the section is published
	IsPublished() bool
	// GetTimestamps returns the section timestamps
	GetTimestamps() (createdAt, updatedAt *time.Time)

	// SetModuleID sets the module ID
	SetModuleID(moduleID string)
	// SetCourseID sets the course ID
	SetCourseID(courseID string)
	// SetSectionID sets the section ID
	SetSectionID(sectionID string)
	// SetIndex sets the section position index
	SetIndex(index uint16)
	// SetTitle sets the section title
	SetTitle(title string)
	// Publish publishes the section
	Publish()
	// Unpublish unpublishes the section
	Unpublish()
}

// GetSectionsFromModule returns the sections from a module
type GetSectionsFromModule interface {
	// Run execute usecase to get sections from module
	Run(moduleID string) ([]Section, error)
}

// CreateSectionInModule creates a section in a module
type CreateSectionInModule interface {
	// Run execute usecase to create section in module
	Run(section Section) (Section, error)
}

// EditSectionInfo updates a section in a module
type EditSectionInfo interface {
	// Run execute usecase to update section in module
	Run(section Section) (Section, error)
}

// GetSectionLessons returns the lessons from a section
type GetSectionLessons interface {
	// Run execute usecase to get lessons from section
	// TODO: []interface{} return should be replaced by []Lesson
	Run(sectionID string) ([]interface{}, error)
}

// SectionUpdater is a interface who provides methods to update a module. It
// works with together the SectionRepository.Edit method.
// We can handle all section properties inside that callback function.
type SectionUpdater func(section Section) (Section, error)

// SectionRepository is a interface who provides methods to manage sections.
type SectionsRepository interface {
	// IssetModule returns true if the module exists
	IssetModule(moduleID string) (string, bool)
	// GetByModuleID returns the sections from a module
	GetByModuleID(moduleID string) ([]Section, error)
	// GetByID returns the section from a module
	GetByID(sectionID string) (Section, error)
	// Create creates a section in a module
	Create(section Section) (Section, error)
	// Edit updates a section in a module
	Edit(section Section, updater SectionUpdater) (Section, error)
	// GetLessons returns the lessons from a section
	// TODO: []interface{} return should be replaced by []Lesson
	GetLessons(sectionID string) ([]interface{}, error)
}
