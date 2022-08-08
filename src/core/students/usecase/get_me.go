package usecase

import (
	"github.com/jeanmolossi/effective-eureka/src/core/students/domain"
	"github.com/jeanmolossi/effective-eureka/src/pkg/auth"
	ormcondition "github.com/jeanmolossi/effective-eureka/src/pkg/orm_condition"
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
func (g *getMe) Run(input *domain.GetMeParams) (domain.Student, error) {
	filters := ormcondition.NewFilterConditions()
	filters.AddFields(input.Fields)
	filters.AddCondition("student_id", input.StudentID)

	student, err := g.repo.GetStudentByID(filters)
	if err != nil {
		return nil, err
	}

	return student, nil
}
