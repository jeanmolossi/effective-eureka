package repository

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CourseModel is a model to manage courses. It reflects the database entity
type CourseModel struct {
	CourseID          string    `gorm:"primary_key;column:course_id;unique"`
	CourseTitle       string    `gorm:"column:course_title"`
	CourseThumb       string    `gorm:"column:course_thumb"`
	CourseDescription string    `gorm:"column:course_description"`
	CoursePublished   bool      `gorm:"column:course_published;index"`
	CourseCreatedAt   time.Time `gorm:"column:course_created_at"`
	CourseUpdatedAt   time.Time `gorm:"column:course_updated_at"`
}

// BeforeCreate is a hook to set the created_at, updated_at fields and generate
// uuid random.
func (c *CourseModel) BeforeCreate(tx *gorm.DB) error {
	c.CourseID = uuid.NewString()
	c.CourseCreatedAt = time.Now()
	c.CourseUpdatedAt = time.Now()

	return nil
}

// BeforeUpdate is a hook to set the updated_at field.
func (c *CourseModel) BeforeUpdate(tx *gorm.DB) error {
	c.CourseUpdatedAt = time.Now()
	return nil
}
