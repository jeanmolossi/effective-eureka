package factory

import "github.com/jeanmolossi/effective-eureka/src/core/students/domain"

type StudentFactory interface {
	CreateStudent(email, password string) StudentFactory
	WithID(id string) StudentFactory
	Build() domain.Student
}

type student struct {
	domain.Student
}

func NewStudent() StudentFactory {
	return &student{
		Student: domain.NewStudent("", ""),
	}
}

func (s *student) CreateStudent(email, password string) StudentFactory {
	s.Student.SetStudentEmail(email)
	s.Student.SetPassword(password)
	return s
}

func (s *student) WithID(id string) StudentFactory {
	s.Student.SetStudentID(id)
	return s
}

func (s *student) Build() domain.Student {
	return s.Student
}
