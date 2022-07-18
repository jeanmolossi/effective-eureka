package auth

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SessionModel struct {
	SessID       string    `gorm:"column:sess_id;primary_key"`
	StudentID    string    `gorm:"column:sess_student_id"`
	Expiration   time.Time `gorm:"column:sess_expiration"`
	AccessToken  string    `gorm:"column:sess_access_token"`
	RefreshToken string    `gorm:"column:sess_refresh_token"`
}

func (sm *SessionModel) BeforeCreate(*gorm.DB) error {
	sm.SessID = uuid.NewString()
	if sm.Expiration.IsZero() {
		sm.Expiration = time.Now().UTC().Local().Add(time.Minute * 10)
	}

	return nil
}
