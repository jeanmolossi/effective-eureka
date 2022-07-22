package usecase

import (
	"github.com/jeanmolossi/effective-eureka/src/core/students/domain"
	"github.com/jeanmolossi/effective-eureka/src/pkg/auth"
)

type getMe struct {
	repo         domain.StudentRepository
	authProvider *auth.SessionProvider
}

// NewGetMe creates a new usecase to get a student by hash.
func NewGetMe(repo domain.StudentRepository, authProvider *auth.SessionProvider) domain.GetMe {
	return &getMe{repo, authProvider}
}

// GetMe gets a student by authentication hash.
func (g *getMe) Run(hash string) (domain.Student, error) {
	session, err := g.authProvider.GetSession(hash)
	if err != nil {
		return nil, err
	}

	if session.IsValid() && !session.IsRefreshExpired() {
		student, err := g.repo.GetStudentByID(session.StudentID)
		if err != nil {
			return nil, err
		}

		return student, nil
	}

	return nil, nil
}
