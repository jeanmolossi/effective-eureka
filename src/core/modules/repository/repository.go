// Package repository is a package to manage modules in Go.
package repository

import (
	"github.com/jeanmolossi/effective-eureka/src/core/modules/domain"
	"gorm.io/gorm"
)

type moduleRepository struct {
	db    *gorm.DB
	table string
}

func NewRepository(db *gorm.DB) *moduleRepository {
	return &moduleRepository{
		db:    db,
		table: "modules",
	}
}

// IssetCourseID returns true if module has a parent course ID.
func (r *moduleRepository) IssetCourseID(courseID string) bool {
	courseModel := struct {
		CourseID string `gorm:"column:course_id"`
	}{}

	courseResult := r.db.Table("courses").Select("course_id").Where(
		"course_id = ?", courseID,
	).First(&courseModel)

	return courseResult.RowsAffected != 0
}

// GetByID returns a module by ID.
func (r *moduleRepository) GetByID(moduleID string) (domain.Module, error) {
	model := &ModuleModel{}
	result := r.db.Table(r.table).Where("module_id = ?", moduleID).First(model)
	if result.Error != nil {
		return nil, result.Error
	}

	return ModelToDomain(model), nil
}

// GetByCourseID returns a list of modules by course ID.
func (r *moduleRepository) GetByCourseID(courseID string) ([]domain.Module, error) {
	if !r.IssetCourseID(courseID) {
		return nil, gorm.ErrRecordNotFound
	}

	models := []*ModuleModel{}
	result := r.db.Table(r.table).Where("course_id = ?", courseID).Find(&models)
	if result.Error != nil {
		return nil, result.Error
	}

	domainsModules := make([]domain.Module, len(models))
	for i, model := range models {
		domainsModules[i] = ModelToDomain(model)
	}

	return domainsModules, nil
}

// Create creates a new module.
func (r *moduleRepository) Create(module domain.Module) (domain.Module, error) {
	if !r.IssetCourseID(module.GetCourseID()) {
		return nil, gorm.ErrRecordNotFound
	}

	model := DomainToModel(module)
	result := r.db.Table(r.table).Create(model)
	if result.Error != nil {
		return nil, result.Error
	}

	return ModelToDomain(model), nil
}

// Edit updates a module.
func (r *moduleRepository) Edit(moduleID string, updater domain.ModuleUpdater) (domain.Module, error) {
	module, err := r.GetByID(moduleID)
	if err != nil {
		return nil, err
	}

	if module == nil {
		return nil, gorm.ErrRecordNotFound
	}

	updatedModule, err := updater(module)
	if err != nil {
		return nil, err
	}

	model := DomainToModel(updatedModule)
	result := r.db.Table(r.table).Save(model)
	if result.Error != nil {
		return nil, result.Error
	}

	return ModelToDomain(model), nil
}
