// Package repository is a package to manage courses in Go.
package repository

import (
	"errors"
	"fmt"

	"github.com/jeanmolossi/effective-eureka/src/core/courses/domain"
	"github.com/jeanmolossi/effective-eureka/src/core/shared"
	"gorm.io/gorm"
)

// repo implements domain.CourseRepository
type repo struct {
	db    *gorm.DB
	table string
}

// NewRepository is a factory method to create a repository to manage courses.
func NewRepository(db *gorm.DB) domain.CourseRepository {
	return &repo{db, "courses"}
}

// GetByID receives and course ID and query that on properly database and returns
// the course found or nil and error if not found.
func (r *repo) GetByID(courseID string) (domain.Course, error) {
	model := &CourseModel{}
	result := r.db.Table(r.table).Where("course_id = ?", courseID).First(&CourseModel{}).Scan(model)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, shared.NewNotFoundErr(
				errors.New("course not found"),
			)
		}

		return nil, fmt.Errorf("error getting course: %v", result.Error)
	}

	return ModelToDomain(model), nil
}

func (r *repo) GetByStudentID(studentID string) ([]domain.Course, error) {
	return nil, fmt.Errorf("not implemented")
}

// GetCourses returns all courses.
func (r *repo) GetCourses(filters shared.FilterConditions) ([]domain.Course, error) {
	models := []*CourseModel{}

	result := r.db.Table(r.table)
	if filters.WithFields() {
		result = result.Select(filters.OnlyFields(r.table))
	}

	if filters.HasConditions() {
		statement, values := filters.Conditions()
		result = result.Where(statement, values...)
	}

	result = result.Find(&models)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, shared.NewNotFoundErr(
				errors.New("courses not found"),
			)
		}

		return nil, fmt.Errorf("error getting courses: %v", result.Error)
	}

	courses := make([]domain.Course, len(models))
	for i, model := range models {
		courses[i] = ModelToDomain(model)
	}

	return courses, nil
}

// Create is the method who creates a course. It should returns the course created
// else it should return nil course and creation error
func (r *repo) Create(course domain.Course) (domain.Course, error) {
	modelCourse := DomainToModel(course, nil, nil)

	result := r.db.Table(r.table).Create(modelCourse).Scan(&modelCourse)
	if result.Error != nil {
		return nil, fmt.Errorf("error creating course: %v", result.Error)
	}

	return ModelToDomain(modelCourse), nil
}

func (r *repo) Edit(courseID string, courseUpdater domain.CourseUpdater) (domain.Course, error) {
	// Get course by ID
	course, err := r.GetByID(courseID)
	if err != nil {
		return nil, err
	}

	if course == nil {
		return nil, fmt.Errorf("course not found")
	}

	// execute updater
	updatedCourse, err := courseUpdater(course)
	if err != nil {
		return nil, fmt.Errorf("[courseUpdater] error updating course: %v", err)
	}

	model := DomainToModel(updatedCourse, nil, nil)
	// Save course: https://gorm.io/docs/update.html#Save-All-Fields
	result := r.db.Table(r.table).Where("course_id = ?", courseID).Save(model).Scan(model)

	if result.Error != nil {
		return nil, fmt.Errorf("error updating course: %v", result.Error)
	}

	return ModelToDomain(model), nil
}
