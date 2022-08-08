package handler

import "github.com/jeanmolossi/effective-eureka/src/core/students/domain"

type HttpStudentRegistered struct {
	StudentID    string `json:"student_id" example:"05d4d9d3-01a3-4fd3-8d3e-e3178522f514"`
	StudentEmail string `json:"student_email" example:"john@doe.com"`
}

func NewHttpStudentRegistered(student domain.Student) *HttpStudentRegistered {
	return &HttpStudentRegistered{
		StudentID:    student.GetStudentID(),
		StudentEmail: student.GetStudentEmail(),
	}
}

type HttpStudent struct {
	StudentID    string `json:"student_id,omitempty" example:"05d4d9d3-01a3-4fd3-8d3e-e3178522f514"`
	StudentEmail string `json:"student_email,omitempty" example:"jean@email.com"`
}

func NewHttpStudent(student domain.Student) *HttpStudent {
	return &HttpStudent{
		StudentID:    student.GetStudentID(),
		StudentEmail: student.GetStudentEmail(),
	}
}
