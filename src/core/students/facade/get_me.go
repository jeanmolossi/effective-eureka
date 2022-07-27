package facade

import (
	"github.com/jeanmolossi/effective-eureka/src/core/students/domain"
	"github.com/jeanmolossi/effective-eureka/src/core/students/repository"
	"github.com/jeanmolossi/effective-eureka/src/core/students/usecase"
	"github.com/jeanmolossi/effective-eureka/src/pkg/auth"
	"gorm.io/gorm"
)

func NewGetMe(db *gorm.DB) domain.GetMe {
	getMe := usecase.NewGetMe(
		repository.NewStudent(db),
		auth.NewSessionProvider(db))

	return getMe
}
