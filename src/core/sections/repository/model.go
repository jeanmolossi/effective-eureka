package repository

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LazyModuleModel struct {
	ModuleID string `json:"module_id" gorm:"column:module_id"`
	CourseID string `json:"course_id" gorm:"column:course_id"`
}

type SectionModel struct {
	ModuleID         string    `json:"module_id" gorm:"column:module_id"`
	CourseID         string    `json:"course_id" gorm:"column:course_id"`
	SectionID        string    `json:"section_id" gorm:"primaryKey;column:section_id"`
	SectionIndex     uint16    `json:"section_index" gorm:"column:section_index"`
	SectionTitle     string    `json:"section_title" gorm:"column:section_title"`
	SectionPublished bool      `json:"section_published" gorm:"column:section_published"`
	SectionCreatedAt time.Time `json:"section_created_at" gorm:"column:section_created_at"`
	SectionUpdatedAt time.Time `json:"section_updated_at" gorm:"column:section_updated_at"`
}

// BeforeCreate is a hook to set the created_at, updated_at fields and generate
// uuid random.
func (s *SectionModel) BeforeCreate(tx *gorm.DB) error {
	s.SectionID = uuid.NewString()
	s.SectionCreatedAt = time.Now()
	s.SectionUpdatedAt = time.Now()

	return nil
}

// BeforeUpdate is a hook to set the updated_at field.
func (s *SectionModel) BeforeUpdate(tx *gorm.DB) error {
	s.SectionUpdatedAt = time.Now()
	return nil
}
