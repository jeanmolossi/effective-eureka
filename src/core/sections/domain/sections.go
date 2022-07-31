package domain

import "time"

type section struct {
	ModuleID  string
	CourseID  string
	SectionID string
	// SectionIndex uint16 Range: 0-65535
	SectionIndex     uint16
	SectionTitle     string
	SectionPublished bool
	SectionCreatedAt time.Time
	SectionUpdatedAt time.Time
}

func NewSection(moduleID, courseID, title string, index uint16, published bool, createdAt, updatedAt *time.Time) Section {
	var sectionCreatedAt, sectionUpdatedAt time.Time

	if createdAt != nil {
		sectionCreatedAt = *createdAt
	}

	if updatedAt != nil {
		sectionUpdatedAt = *updatedAt
	}

	return &section{
		ModuleID:         moduleID,
		CourseID:         courseID,
		SectionID:        "",
		SectionIndex:     index,
		SectionTitle:     title,
		SectionPublished: published,
		SectionCreatedAt: sectionCreatedAt,
		SectionUpdatedAt: sectionUpdatedAt,
	}
}

// GetModuleID returns the module ID
func (s *section) GetModuleID() string {
	return s.ModuleID
}

// GetCourseID returns the course ID
func (s *section) GetCourseID() string {
	return s.CourseID
}

// GetSectionID returns the section ID
func (s *section) GetSectionID() string {
	return s.SectionID
}

// GetIndex returns the section position index
func (s *section) GetIndex() uint16 {
	return s.SectionIndex
}

// GetTitle returns the section title
func (s *section) GetTitle() string {
	return s.SectionTitle
}

// IsPublished returns true if the section is published
func (s *section) IsPublished() bool {
	return s.SectionPublished
}

// GetTimestamps returns the section timestamps
func (s *section) GetTimestamps() (createdAt, updatedAt *time.Time) {
	return &s.SectionCreatedAt, &s.SectionUpdatedAt
}

// SetModuleID sets the module ID
func (s *section) SetModuleID(moduleID string) {
	s.ModuleID = moduleID
}

// SetCourseID sets the course ID
func (s *section) SetCourseID(courseID string) {
	s.CourseID = courseID
}

// SetSectionID sets the section ID
func (s *section) SetSectionID(sectionID string) {
	s.SectionID = sectionID
}

// SetIndex sets the section position index
func (s *section) SetIndex(index uint16) {
	s.SectionIndex = index
}

// SetTitle sets the section title
func (s *section) SetTitle(title string) {
	s.SectionTitle = title
}

// Publish publishes the section
func (s *section) Publish() {
	s.SectionPublished = true
}

// Unpublish unpublishes the section
func (s *section) Unpublish() {
	s.SectionPublished = false
}
