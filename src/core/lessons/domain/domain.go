package domain

import (
	"time"

	ormcondition "github.com/jeanmolossi/effective-eureka/src/pkg/orm_condition"
	"github.com/jeanmolossi/effective-eureka/src/pkg/paginator"
)

type Lesson interface {
	GetSectionID() string
	GetLessonID() string
	GetTitle() string
	GetDescription() string
	GetThumbnail() string
	GetVideoPreview() string
	GetVideo() string
	GetIndex() uint16
	IsPublished() bool
	GetTimestamps() (createdAt, updatedAt time.Time)

	SetSectionID(sectionID string)
	SetLessonID(lessonID string)
	SetTitle(title string)
	SetDescription(description string)
	SetThumbnail(thumbnail string)
	SetVideoPreview(previewUrl string)
	SetVideo(video string)
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

type LessonUpdater func(lesson Lesson) (Lesson, error)

type LessonRepository interface {
	IssetSection(sectionID string) bool
	Create(lesson Lesson) (Lesson, error)
	GetLesson(filters ormcondition.FilterConditions) (Lesson, error)
	GetLessonsFromSection(filters ormcondition.FilterConditions, pagination paginator.Paginator) ([]Lesson, error)
	Edit(lessonID string, updater LessonUpdater) (Lesson, error)
}
