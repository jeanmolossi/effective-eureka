package auth

import (
	"errors"
	"time"

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

	return sp.repo.CreateSession(student.GetStudentID())
}

func (sp *SessionProvider) GetSession(token string) (*SessionModel, error) {
	return sp.repo.GetSession(token)
}

func (sp *SessionProvider) DeleteSession(token string) error {
	return sp.repo.DeleteSession(token)
}

func (sp *SessionProvider) IsValidSession(token string) bool {
	session, err := sp.GetSession(token)
	if err != nil {
		return false
	}

	if session.Expiration.After(time.Now()) {
		return true
	}

	return true
}
