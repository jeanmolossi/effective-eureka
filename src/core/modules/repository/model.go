package repository

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ModuleModel struct {
	CourseID          string    `json:"course_id" gorm:"column:course_id"`
	ModuleID          string    `json:"module_id" gorm:"primaryKey;column:module_id"`
	ModuleTitle       string    `json:"module_title" gorm:"column:module_title"`
	ModuleDescription string    `json:"module_description" gorm:"column:module_description"`
	ModuleThumb       string    `json:"module_thumb" gorm:"column:module_thumb"`
	ModulePublished   bool      `json:"module_published" gorm:"column:module_published"`
	ModuleCreatedAt   time.Time `json:"module_created_at" gorm:"column:module_created_at"`
	ModuleUpdatedAt   time.Time `json:"module_updated_at" gorm:"column:module_updated_at"`
}

// BeforeCreate is a hook to set the created_at, updated_at fields and generate
// uuid random.
func (c *ModuleModel) BeforeCreate(tx *gorm.DB) error {
	c.ModuleID = uuid.NewString()
	c.ModuleCreatedAt = time.Now()
	c.ModuleUpdatedAt = time.Now()

	return nil
}

// BeforeUpdate is a hook to set the updated_at field.
func (c *ModuleModel) BeforeUpdate(tx *gorm.DB) error {
	c.ModuleUpdatedAt = time.Now()
	return nil
}
