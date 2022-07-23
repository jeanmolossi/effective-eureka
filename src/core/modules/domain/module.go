package domain

type module struct {
	courseID    string
	moduleID    string
	title       string
	thumb       string
	description string
	published   bool
}

func NewModule(courseID, title, thumb, description string, published bool) Module {
	return &module{
		courseID:    courseID,
		title:       title,
		thumb:       thumb,
		description: description,
		published:   published,
	}
}

// GetCourseID returns parent course ID.
func (m *module) GetCourseID() string {
	return m.courseID
}

// GetModuleID returns module ID.
func (m *module) GetModuleID() string {
	return m.moduleID
}

// GetModuleTitle returns module title.
func (m *module) GetModuleTitle() string {
	return m.title
}

// GetModuleDescription returns module description.
func (m *module) GetModuleDescription() string {
	return m.description
}

// GetModuleThumb returns module thumb.
func (m *module) GetModuleThumb() string {
	return m.thumb
}

// IsModulePublished returns true if module is published.
func (m *module) IsModulePublished() bool {
	return m.published
}

// SetCourseID sets parent course ID.
func (m *module) SetCourseID(courseID string) {
	m.courseID = courseID
}

// SetModuleID sets module ID.
func (m *module) SetModuleID(moduleID string) {
	m.moduleID = moduleID
}

// SetModuleTitle sets module title.
func (m *module) SetModuleTitle(title string) {
	m.title = title
}

// SetModuleDescription sets module description.
func (m *module) SetModuleDescription(desc string) {
	m.description = desc
}

// SetModuleThumb sets module thumb.
func (m *module) SetModuleThumb(thumb string) {
	m.thumb = thumb
}

// PublishModule publishes module.
func (m *module) PublishModule() {
	m.published = true
}

// UnpublishModule unpublishes module.
func (m *module) UnpublishModule() {
	m.published = false
}
