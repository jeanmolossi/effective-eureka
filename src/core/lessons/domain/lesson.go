package domain

import "time"

type lesson struct {
	sectionID   string
	lessonID    string
	index       uint16
	title       string
	description string
	thumb       string
	published   bool
	createdAt   time.Time
	updatedAt   time.Time
}

func NewLesson(sectionID, lessonID, title, description, thumb string, index uint16, published bool, createdAt, updatedAt *time.Time) Lesson {
	less := &lesson{
		sectionID:   sectionID,
		lessonID:    lessonID,
		index:       index,
		title:       title,
		description: description,
		thumb:       thumb,
		published:   published,
	}

	if createdAt != nil {
		less.createdAt = *createdAt
	}

	if updatedAt != nil {
		less.updatedAt = *updatedAt
	}

	return less
}

func (l *lesson) GetSectionID() string {
	return l.sectionID
}

func (l *lesson) GetLessonID() string {
	return l.lessonID
}

func (l *lesson) GetTitle() string {
	return l.title
}

func (l *lesson) GetDescription() string {
	return l.description
}

func (l *lesson) GetThumbnail() string {
	return l.thumb
}

func (l *lesson) GetIndex() uint16 {
	return l.index
}

func (l *lesson) IsPublished() bool {
	return l.published
}

func (l *lesson) GetTimestamps() (createdAt, updatedAt time.Time) {
	return l.createdAt, l.updatedAt
}

func (l *lesson) SetSectionID(sectionID string) {
	l.sectionID = sectionID
}

func (l *lesson) SetLessonID(lessonID string) {
	l.lessonID = lessonID
}

func (l *lesson) SetTitle(title string) {
	l.title = title
}

func (l *lesson) SetDescription(description string) {
	l.description = description
}

func (l *lesson) SetThumbnail(thumbnail string) {
	l.thumb = thumbnail
}

func (l *lesson) SetIndex(index uint16) {
	l.index = index
}

func (l *lesson) Publish() {
	l.published = true
}

func (l *lesson) Unpublish() {
	l.published = false
}
