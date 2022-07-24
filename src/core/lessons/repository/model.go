package repository

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LazeSectionModel struct {
	SectionID string `gorm:"primaryKey;column:section_id"`
}

type LessonModel struct {
	SectionID          string    `json:"section_id" gorm:"column:section_id"`
	LessonID           string    `json:"lesson_id" gorm:"primaryKey;column:lesson_id"`
	LessonIndex        uint16    `json:"lesson_index" gorm:"column:lesson_index"`
	LessonTitle        string    `json:"lesson_title" gorm:"column:lesson_title"`
	LessonDescription  string    `json:"lesson_description" gorm:"column:lesson_description"`
	LessonThumb        string    `json:"lesson_thumb" gorm:"column:lesson_thumb"`
	LessonVideo        string    `json:"lesson_video" gorm:"column:lesson_video"`
	LessonVideoPreview string    `json:"lesson_video_preview" gorm:"column:lesson_video_preview"`
	LessonPublished    bool      `json:"lesson_published" gorm:"column:lesson_published"`
	LessonCreatedAt    time.Time `json:"lesson_created_at" gorm:"column:lesson_created_at"`
	LessonUpdatedAt    time.Time `json:"lesson_updated_at" gorm:"column:lesson_updated_at"`
}

// BeforeCreate is a hook to set the created_at, updated_at fields and generate
// uuid random.
func (l *LessonModel) BeforeCreate(tx *gorm.DB) error {
	l.LessonID = uuid.NewString()
	l.LessonCreatedAt = time.Now()
	l.LessonUpdatedAt = time.Now()

	return nil
}

// BeforeUpdate is a hook to set the updated_at field.
func (l *LessonModel) BeforeUpdate(tx *gorm.DB) error {
	l.LessonUpdatedAt = time.Now()
	return nil
}
