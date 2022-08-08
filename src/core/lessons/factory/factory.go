package factory

import (
	"time"

	"github.com/jeanmolossi/effective-eureka/src/core/lessons/domain"
)

type Lesson interface {
	CreateLesson(title string, description string, thumb, videoPreview, video string, index uint16, published bool, createdAt *time.Time, updatedAt *time.Time) Lesson
	WithLessonID(lessonID string) Lesson
	WithSectionID(sectionID string) Lesson
	Build() domain.Lesson
}

type lesson struct {
	domain.Lesson
}

func NewLesson() Lesson {
	return &lesson{domain.NewLesson(
		"",    // sectionID
		"",    // lessonID
		"",    // title
		"",    // description
		"",    // thumb
		"",    // videoPreview
		"",    // video
		0,     // index
		false, // published
		nil,   // createdAt
		nil),  // updatedAt
	}
}

func (l *lesson) CreateLesson(title, description, thumb, videoPreview, video string, index uint16, published bool, createdAt *time.Time, updatedAt *time.Time) Lesson {
	if createdAt != nil && updatedAt != nil {
		l.Lesson = domain.NewLesson("", "",
			title,
			description,
			thumb,
			videoPreview,
			video,
			index,
			published,
			createdAt, updatedAt)
	} else {
		l.Lesson.SetTitle(title)
		l.Lesson.SetDescription(description)
		l.Lesson.SetThumbnail(thumb)
		l.Lesson.SetVideoPreview(videoPreview)
		l.Lesson.SetVideo(video)
		l.Lesson.SetIndex(index)

		if published {
			l.Lesson.Publish()
		} else {
			l.Lesson.Unpublish()
		}
	}

	return l
}

func (l *lesson) WithLessonID(lessonID string) Lesson {
	l.Lesson.SetLessonID(lessonID)
	return l
}

func (l *lesson) WithSectionID(sectionID string) Lesson {
	l.Lesson.SetSectionID(sectionID)
	return l
}

func (l *lesson) Build() domain.Lesson {
	return l.Lesson
}
