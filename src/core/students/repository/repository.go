package repository

import (
	"errors"

	"github.com/jeanmolossi/effective-eureka/src/core/students/domain"
	"gorm.io/gorm"
)

// studentRepository represents a student repository.
type studentRepository struct {
	db *gorm.DB
}

// NewStudent instantiate a new student repository.
func NewStudent(db *gorm.DB) domain.StudentRepository {
	return &studentRepository{db}
}

// GetStudentByID returns a student by ID.
func (s *studentRepository) GetStudentByID(studentID string) (domain.Student, error) {
	model := &StudentModel{}
	result := s.db.Table("students").Where("student_id = ?", studentID).First(model)

	if result.Error != nil {
		return nil, result.Error
	}

	return ModelToDomain(model), nil
}

// GetStudentByEmail returns a student by email.
func (s *studentRepository) GetStudentByEmail(studentEmail string) (domain.Student, error) {
	model := &StudentModel{}
	result := s.db.Table("students").Where("student_email = ?", studentEmail).First(model)

	if result.Error != nil {
		return nil, result.Error
	}

	return ModelToDomain(model), nil
}

// CreateStudent creates a new student.
func (s *studentRepository) CreateStudent(student domain.Student) (domain.Student, error) {
	model := DomainToModel(student)
	result := s.db.Table("students").Create(model).Scan(model)

	if result.Error != nil {
		return nil, result.Error
	}

	return ModelToDomain(model), nil
}

// UpdateStudent updates a student.
func (s *studentRepository) UpdateStudent(student domain.Student) (domain.Student, error) {
	return nil, errors.New("not implemented")
}

// DeleteStudent deletes a student.
func (s *studentRepository) DeleteStudent(studentID string) error {
	return errors.New("not implemented")
}