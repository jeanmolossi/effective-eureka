package repository

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// StudentModel represents a student in the database.
type StudentModel struct {
	StudentID string    `gorm:"primary_key;column:student_id;unique"`
	Email     string    `gorm:"column:student_email"`
	Password  string    `gorm:"column:student_password"`
	CreatedAt time.Time `gorm:"column:student_created_at"`
	UpdatedAt time.Time `gorm:"column:student_updated_at"`
}

// BeforeCreate is a hook to set the created_at, updated_at fields and generate
// uuid random.
func (s *StudentModel) BeforeCreate(tx *gorm.DB) error {
	s.StudentID = uuid.NewString()
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()

	return nil
}

// BeforeUpdate is a hook to set the updated_at field.
func (s *StudentModel) BeforeUpdate(tx *gorm.DB) error {
	s.UpdatedAt = time.Now()
	return nil
}
