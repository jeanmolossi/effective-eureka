package repository

import "github.com/jeanmolossi/effective-eureka/src/core/lessons/domain"

func ModelToDomain(model *LessonModel) domain.Lesson {
	if model == nil {
		return domain.NewLesson("", "", "", "", "", 0, false, nil, nil)
	}

	return domain.NewLesson(
		model.SectionID,
		model.LessonID,
		model.LessonTitle,
		model.LessonDescription,
		model.LessonThumb,
		model.LessonIndex,
		model.LessonPublished,
		&model.LessonCreatedAt,
		&model.LessonUpdatedAt,
	)
}

func DomainToModel(lesson domain.Lesson) *LessonModel {
	createdAt, updatedAt := lesson.GetTimestamps()
	return &LessonModel{
		SectionID:         lesson.GetSectionID(),
		LessonID:          lesson.GetLessonID(),
		LessonTitle:       lesson.GetTitle(),
		LessonDescription: lesson.GetDescription(),
		LessonThumb:       lesson.GetThumbnail(),
		LessonIndex:       lesson.GetIndex(),
		LessonPublished:   lesson.IsPublished(),
		LessonCreatedAt:   createdAt,
		LessonUpdatedAt:   updatedAt,
	}
}
