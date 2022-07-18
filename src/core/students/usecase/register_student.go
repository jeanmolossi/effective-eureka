package usecase

import (
	"errors"

	"github.com/jeanmolossi/effective-eureka/src/core/students/domain"
)

// registerStudent represents a register student usecase.
type registerStudent struct {
	studentRepository domain.StudentRepository
}

// NewRegisterStudent instantiate a new usecase.
func NewRegisterStudent(studentRepository domain.StudentRepository) domain.RegisterStudent {
	return &registerStudent{studentRepository}
}

// Run runs the usecase.
func (r *registerStudent) Run(student domain.Student) (domain.Student, error) {
	alreadySetUser, err := r.studentRepository.GetStudentByEmail(student.GetStudentEmail())
	if err != nil {
		if err.Error() != "record not found" {
			return nil, err
		}
	}

	if alreadySetUser != nil {
		return nil, errors.New("student already exists")
	}

	err = student.HashPassword()
	if err != nil {
		return nil, err
	}

	return r.studentRepository.CreateStudent(student)
}
