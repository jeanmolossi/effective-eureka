package auth

import (
	"log"

	"gorm.io/gorm"
)

type SessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) *SessionRepository {
	return &SessionRepository{db}
}

func (sr *SessionRepository) CreateSession(session *Session) (*Session, error) {
	sessionModel := DomainToModel(session)

	result := sr.db.Table("sess").Where("sess_student_id = ?", session.StudentID).First(sessionModel)
	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return nil, result.Error
		}
	}

	if result.RowsAffected > 0 {
		domainSess := ModelToDomain(sessionModel)
		if domainSess.AccessToken != nil {
			return domainSess, nil
		} else {
			sr.DeleteSession(domainSess.SessID)
			sessionModel = DomainToModel(session)
		}
	}

	result = sr.db.Table("sess").Create(sessionModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return ModelToDomain(sessionModel), nil
}

func (sr *SessionRepository) GetSession(token string) (*SessionModel, error) {
	session := &SessionModel{}
	result := sr.db.Table("sess").Where("sess_id = ?", token).First(session)

	if result.Error != nil {
		return nil, result.Error
	}

	return session, nil
}

func (sr *SessionRepository) UpdateSession(sessionID string, updater func(session *Session) (*Session, error)) error {
	model := &SessionModel{}
	result := sr.db.Table("sess").Where("sess_id = ?", sessionID).First(model)
	if result.Error != nil {
		return result.Error
	}

	session := ModelToDomain(model)
	updatedSession, err := updater(session)
	if err != nil {
		return err
	}

	updatedModel := DomainToModel(updatedSession)
	result = sr.db.Table("sess").Save(updatedModel)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (sr *SessionRepository) DeleteSession(sessionID string) error {
	session := &SessionModel{}
	result := sr.db.Table("sess").Where("sess_id = ?", sessionID).First(session)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			log.Println("session not found")
			return nil
		}

		return result.Error
	}

	result = sr.db.Table("sess").Delete(session)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
