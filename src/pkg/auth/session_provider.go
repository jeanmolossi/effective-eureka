package auth

import (
	"errors"
	"log"

	"github.com/jeanmolossi/effective-eureka/src/core/students/domain"
	"github.com/jeanmolossi/effective-eureka/src/core/students/repository"
	"gorm.io/gorm"
)

type SessionProvider struct {
	repo        *SessionRepository
	studentRepo domain.StudentRepository
}

func NewSessionProvider(db *gorm.DB) *SessionProvider {
	repo := NewSessionRepository(db)
	studentRepo := repository.NewStudent(db)
	return &SessionProvider{repo, studentRepo}
}

func (sp *SessionProvider) CreateSession(username, password string) (*Session, error) {
	student, err := sp.studentRepo.GetStudentByEmail(username)
	if err != nil {
		return nil, err
	}

	if !student.IsValidPassword(password) {
		return nil, errors.New("invalid password")
	}

	accessToken := NewAccessToken(student.GetStudentID())
	session := NewSession("",
		student.GetStudentID(),
		accessToken,
		NewRefreshToken(student.GetStudentID()),
		accessToken.Expiration,
	)

	return sp.repo.CreateSession(session)
}

func (sp *SessionProvider) GetSession(hash string) (*Session, error) {
	studentID, sessionID, err := Decode(hash)
	if err != nil {
		return nil, err
	}

	model, err := sp.repo.GetSession(sessionID)
	if err != nil {
		return nil, err
	}

	if model.StudentID != studentID {
		err = sp.repo.DeleteSession(sessionID)
		if err != nil {
			return nil, err
		}

		return nil, errors.New("invalid session")
	}

	return ModelToDomain(model), nil
}

func (sp *SessionProvider) DeleteSession(sessionID string) error {
	return sp.repo.DeleteSession(sessionID)
}

func (sp *SessionProvider) IsValidSession(hash string) bool {
	session, err := sp.GetSession(hash)
	if err != nil {
		return false
	}

	if !session.IsRefreshExpired() && session.IsExpired() {
		err = sp.RefreshSession(session)
		if err == nil {
			return true
		}
	}

	if session.IsRefreshExpired() && session.IsValid() {
		err = sp.RenewRefreshToken(session)
		if err == nil {
			return true
		}
	}

	if session.IsRefreshExpired() && session.IsExpired() {
		err = sp.DeleteSession(session.SessID)
		if err != nil {
			log.Println(err)
		}

		return false
	}

	return true
}

func (sp *SessionProvider) RefreshSession(s *Session) error {
	err := sp.repo.UpdateSession(s.SessID, func(session *Session) (*Session, error) {
		session.AccessToken = NewAccessToken(s.StudentID)
		session.Expiration = session.AccessToken.Expiration
		return session, nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (sp *SessionProvider) RenewRefreshToken(s *Session) error {
	err := sp.repo.UpdateSession(s.SessID, func(session *Session) (*Session, error) {
		session.RefreshToken = NewRefreshToken(s.StudentID)
		return session, nil
	})

	if err != nil {
		return err
	}

	return nil
}
