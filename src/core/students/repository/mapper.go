package repository

import (
	"github.com/jeanmolossi/effective-eureka/src/core/students/domain"
	"github.com/jeanmolossi/effective-eureka/src/core/students/factory"
)

func ModelToDomain(model *StudentModel) domain.Student {
	studentFactory := factory.NewStudent().CreateStudent(model.Email, model.Password)
	if model.StudentID != "" {
		studentFactory.WithID(model.StudentID)
	}
	return studentFactory.Build()
}

func DomainToModel(domain domain.Student) *StudentModel {
	return &StudentModel{
		StudentID: domain.GetStudentID(),
		Email:     domain.GetStudentEmail(),
		Password:  domain.GetStudentPassword(),
	}
}
