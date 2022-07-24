package usecase

import "github.com/jeanmolossi/effective-eureka/src/core/lessons/domain"

type editLessonInfo struct {
	repo domain.LessonRepository
}

func NewEditLessonInfo(repo domain.LessonRepository) domain.EditLessonInfo {
	return &editLessonInfo{
		repo: repo,
	}
}

func (e *editLessonInfo) EditLesson(lesson domain.Lesson) (domain.Lesson, error) {
	return e.repo.Edit(lesson.GetLessonID(), e.updater(lesson))
}

func (e *editLessonInfo) updater(newLesson domain.Lesson) domain.LessonUpdater {
	return func(currentLesson domain.Lesson) (domain.Lesson, error) {
		if newLesson.GetTitle() != "" {
			currentLesson.SetTitle(newLesson.GetTitle())
		}

		if newLesson.GetDescription() != "" {
			currentLesson.SetDescription(newLesson.GetDescription())
		}

		if newLesson.GetThumbnail() != "" {
			currentLesson.SetThumbnail(newLesson.GetThumbnail())
		}

		if newLesson.GetIndex() != 0 {
			currentLesson.SetIndex(newLesson.GetIndex())
		}

		if newLesson.IsPublished() {
			currentLesson.Publish()
		} else {
			currentLesson.Unpublish()
		}

		return currentLesson, nil
	}
}
