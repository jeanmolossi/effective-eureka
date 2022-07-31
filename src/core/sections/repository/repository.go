// Package repository is a package to manage sections in Go.
package repository

import (
	"errors"

	ldomain "github.com/jeanmolossi/effective-eureka/src/core/lessons/domain"
	"github.com/jeanmolossi/effective-eureka/src/core/sections/domain"
	"github.com/jeanmolossi/effective-eureka/src/core/shared"
	ormcondition "github.com/jeanmolossi/effective-eureka/src/pkg/orm_condition"
	"github.com/jeanmolossi/effective-eureka/src/pkg/paginator"
	"gorm.io/gorm"
)

type sectionRepository struct {
	db    *gorm.DB
	table string
}

func NewSectionRepository(db *gorm.DB) domain.SectionsRepository {
	return &sectionRepository{
		db:    db,
		table: "sections",
	}
}

// IssetModule returns true if the module exists
func (s *sectionRepository) IssetModule(moduleID string) (string, bool) {
	moduleModel := &LazyModuleModel{}

	result := s.db.Table("modules").Select("module_id, course_id").Where(
		"module_id = ?", moduleID,
	).First(moduleModel)
	if result.Error != nil {
		return "", false
	}

	return moduleModel.CourseID, result.Error == nil
}

// GetByModuleID returns the sections from a module
func (s *sectionRepository) GetByModuleID(filters ormcondition.FilterConditions, paginator paginator.Paginator) ([]domain.Section, error) {
	moduleIDinterface, hasCondition := filters.GetCondition("module_id")
	if !hasCondition {
		return nil, domain.NewBadRequestErr(
			errors.New("module_id is required"),
		)
	}

	moduleID := moduleIDinterface.(string)
	var courseID string
	var issetModule bool

	if courseID, issetModule = s.IssetModule(moduleID); !issetModule {
		return nil, shared.NewNotFoundErr(errors.New("module not found"))
	}

	var sections []*SectionModel
	result := s.db.Table(s.table)

	if filters.WithFields() {
		result = result.Select(filters.SelectFields(s.table))
	}

	if filters.HasConditions() {
		statement, values := filters.Conditions()
		result = result.Where(statement, values...)
	}

	if paginator.ShouldPaginate() {
		result = result.Scopes(paginator.Paginate)
	}

	result = result.Find(&sections)
	if result.Error != nil {
		return nil, result.Error
	}

	domainSections := make([]domain.Section, len(sections))
	for i, section := range sections {
		section.CourseID = courseID
		domainSections[i] = ModelToDomain(section)
	}

	return domainSections, nil
}

// GetByID returns the section from a module
func (s *sectionRepository) GetByID(sectionID string) (domain.Section, error) {
	section := &SectionModel{SectionID: sectionID}

	result := s.db.Table(s.table).First(&section)
	if result.Error != nil {
		return nil, result.Error
	}

	return ModelToDomain(section), nil
}

// Create creates a section in a module
func (s *sectionRepository) Create(section domain.Section) (domain.Section, error) {
	var courseID string
	var issetModule bool
	// can not create section is has no module
	if courseID, issetModule = s.IssetModule(section.GetModuleID()); !issetModule {
		return nil, domain.NewNotFoundErr(errors.New("module not found"))
	}

	// can not create section is has no parent module course
	if courseID == "" {
		return nil, domain.NewNotFoundErr(errors.New("course module parent not found"))
	}

	// auto set course ID by module
	section.SetCourseID(courseID)

	sectionModel := DomainToModel(section)
	result := s.db.Table(s.table).Create(sectionModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return ModelToDomain(sectionModel), nil
}

// Edit updates a section in a module
func (s *sectionRepository) Edit(section domain.Section, updater domain.SectionUpdater) (domain.Section, error) {
	var courseID string
	var issetModule bool
	// can not edit section is has no module
	if courseID, issetModule = s.IssetModule(section.GetModuleID()); !issetModule {
		return nil, domain.NewNotFoundErr(errors.New("module not found"))
	}

	// can not edit section is has no parent module course
	if courseID == "" {
		return nil, domain.NewNotFoundErr(errors.New("course module parent not found"))
	}

	// auto set course ID by module
	section.SetCourseID(courseID)

	currentSection, err := s.GetByID(section.GetSectionID())
	if err != nil {
		return nil, err
	}

	if currentSection == nil {
		return nil, domain.NewNotFoundErr(errors.New("section not found"))
	}

	if currentSection.GetCourseID() != section.GetCourseID() {
		return nil, domain.NewUnauthorizedErr(
			errors.New("can not change section between courses"),
		)
	}

	updatedSection, err := updater(currentSection)
	if err != nil {
		return nil, err
	}

	sectionModel := DomainToModel(updatedSection)
	result := s.db.Table(s.table).Save(sectionModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return ModelToDomain(sectionModel), nil
}

// GetLessons returns the lessons from a section
func (s *sectionRepository) GetLessons(sectionID string) ([]ldomain.Lesson, error) {
	return nil, errors.New("not implemented")
}
