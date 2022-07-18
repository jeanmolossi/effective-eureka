package auth

import (
	"gorm.io/gorm"
)

type SessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) *SessionRepository {
	return &SessionRepository{db}
}

func (sr *SessionRepository) CreateSession(studentID string) (*Session, error) {
	session := &SessionModel{
		StudentID: studentID,
	}

	result := sr.db.Table("sess").Where("sess_student_id = ?", studentID).First(session)
	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return nil, result.Error
		}
	}

	if result.RowsAffected > 0 {
		return ModelToDomain(session), nil
	}

	result = sr.db.Table("sess").Create(session)
	if result.Error != nil {
		return nil, result.Error
	}

	return ModelToDomain(session), nil
}

func (sr *SessionRepository) GetSession(token string) (*SessionModel, error) {
	session := &SessionModel{}
	result := sr.db.Table("sess").Where("sess_id = ?", token).First(session)

	if result.Error != nil {
		return nil, result.Error
	}

	return session, nil
}

func (sr *SessionRepository) DeleteSession(token string) error {
	session := &SessionModel{}
	result := sr.db.Table("sess").Where("sess_id = ?", token).First(session)

	if result.Error != nil {
		return result.Error
	}

	result = sr.db.Table("sess").Delete(session)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
