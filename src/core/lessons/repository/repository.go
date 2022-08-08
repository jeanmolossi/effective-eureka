package repository

import (
	"errors"

	"github.com/jeanmolossi/effective-eureka/src/core/lessons/domain"
	ormcondition "github.com/jeanmolossi/effective-eureka/src/pkg/orm_condition"
	"github.com/jeanmolossi/effective-eureka/src/pkg/paginator"
	"gorm.io/gorm"
)

type lessonRepository struct {
	db    *gorm.DB
	table string
}

func NewRepository(db *gorm.DB) domain.LessonRepository {
	return &lessonRepository{db, "lessons"}
}

func (l *lessonRepository) IssetSection(sectionID string) bool {
	sectionModel := &LazeSectionModel{}
	result := l.db.Table("sections").Select("section_id").Where(
		"section_id = ?", sectionID,
	).First(sectionModel)

	return result.RowsAffected > 0
}

func (l *lessonRepository) Create(lesson domain.Lesson) (domain.Lesson, error) {
	if !l.IssetSection(lesson.GetSectionID()) {
		return nil, domain.NewNotFoundErr(
			errors.New("section not found"),
		)
	}

	model := DomainToModel(lesson)
	result := l.db.Table(l.table).Create(model)
	if result.Error != nil {
		return nil, result.Error
	}

	return ModelToDomain(model), nil
}

func (l *lessonRepository) GetLesson(filters ormcondition.FilterConditions) (domain.Lesson, error) {
	model := &LessonModel{}
	result := l.db.Table(l.table)

	if filters.WithFields() {
		result = result.Select(filters.SelectFields(l.table))
	}

	if filters.HasConditions() {
		result = result.Where(filters.Conditions())
	} else {
		return nil, domain.NewBadRequestErr(
			errors.New("lesson id is required"),
		)
	}

	result = result.First(model)

	if result.Error != nil {
		return nil, result.Error
	}

	return ModelToDomain(model), nil
}

func (l *lessonRepository) GetLessonsFromSection(filters ormcondition.FilterConditions, pagination paginator.Paginator) ([]domain.Lesson, error) {
	sectionIDCondition, exists := filters.GetCondition("section_id")
	if !exists {
		return nil, domain.NewBadRequestErr(
			errors.New("section_id is required"),
		)
	}

	sectionID := sectionIDCondition.(string)

	if !l.IssetSection(sectionID) {
		return nil, domain.NewNotFoundErr(
			errors.New("section not found"),
		)
	}

	models := []*LessonModel{}
	result := l.db.Table(l.table)

	if filters.WithFields() {
		result = result.Select(filters.SelectFields(l.table))
	}

	if filters.HasConditions() {
		statement, values := filters.Conditions()
		result = result.Where(statement, values...)
	}

	if pagination.ShouldPaginate() {
		result = result.Scopes(pagination.Paginate)
	}

	result = result.Find(&models)

	if result.Error != nil {
		return nil, result.Error
	}

	lessons := make([]domain.Lesson, len(models))
	for i, model := range models {
		lessons[i] = ModelToDomain(model)
	}

	return lessons, nil
}

func (l *lessonRepository) Edit(lessonID string, updater domain.LessonUpdater) (domain.Lesson, error) {
	filters := ormcondition.NewFilterConditions()
	filters.AddCondition("lesson_id", lessonID)

	currentLesson, err := l.GetLesson(filters)
	if err != nil {
		return nil, err
	}

	updatedLesson, err := updater(currentLesson)
	if err != nil {
		return nil, err
	}

	model := DomainToModel(updatedLesson)
	result := l.db.Table(l.table).Save(model)
	if result.Error != nil {
		return nil, result.Error
	}

	return ModelToDomain(model), nil
}
