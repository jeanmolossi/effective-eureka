// Package repository is a package to manage courses in Go.
package repository

import (
	"fmt"

	"github.com/jeanmolossi/effective-eureka/src/core/courses/domain"
	"gorm.io/gorm"
)

// repo implements domain.CourseRepository
type repo struct {
	db *gorm.DB
}

// NewRepository is a factory method to create a repository to manage courses.
func NewRepository(db *gorm.DB) domain.CourseRepository {
	return &repo{db}
}

func (r *repo) GetByID(courseID string) (domain.Course, error) {
	return nil, fmt.Errorf("not implemented")
}

func (r *repo) GetByStudentID(studentID string) ([]domain.Course, error) {
	return nil, fmt.Errorf("not implemented")
}

// Create is the method who creates a course. It should returns the course created
// else it should return nil course and creation error
func (r *repo) Create(course domain.Course) (domain.Course, error) {
	modelCourse := DomainToModel(course, nil, nil)

	result := r.db.Table("courses").Create(modelCourse).Scan(&modelCourse)
	if result.Error != nil {
		return nil, fmt.Errorf("error creating course: %v", result.Error)
	}

	return ModelToDomain(modelCourse), nil
}

func (r *repo) Edit(courseID string, courseUpdater domain.CourseUpdater) (domain.Course, error) {
	return nil, fmt.Errorf("not implemented")
}
