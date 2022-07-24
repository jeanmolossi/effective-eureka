package domain

import (
	"time"
)

type Lesson interface {
	GetSectionID() string
	GetLessonID() string
	GetTitle() string
	GetDescription() string
	GetThumbnail() string
	GetIndex() uint16
	IsPublished() bool
	GetTimestamps() (createdAt, updatedAt time.Time)

	SetSectionID(sectionID string)
	SetLessonID(lessonID string)
	SetTitle(title string)
	SetDescription(description string)
	SetThumbnail(thumbnail string)
	SetIndex(index uint16)

	Publish()
	Unpublish()
}

type AddLessonInSection interface {
	AddLesson(lesson Lesson) (Lesson, error)
}

type EditLessonInfo interface {
	EditLesson(lesson Lesson) (Lesson, error)
}

type GetLessonsInSection interface {
	Run(sectionID string) ([]Lesson, error)
}

type LessonUpdater func(lesson Lesson) (Lesson, error)

type LessonRepository interface {
	IssetSection(sectionID string) bool
	Create(lesson Lesson) (Lesson, error)
	GetLesson(lessonID string) (Lesson, error)
	GetLessonsFromSection(sectionID string) ([]Lesson, error)
	Edit(lessonID string, updater LessonUpdater) (Lesson, error)
}
