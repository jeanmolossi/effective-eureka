package domain

import "golang.org/x/crypto/bcrypt"

type student struct {
	studentID       string
	studentEmail    string
	studentPassword string
}

func NewStudent(email, password string) Student {
	return &student{
		studentID:       "",
		studentEmail:    email,
		studentPassword: password,
	}
}

// GetStudentID returns the student's ID.
func (s *student) GetStudentID() string {
	return s.studentID
}

// GetStudentEmail returns the student's email.
func (s *student) GetStudentEmail() string {
	return s.studentEmail
}

// GetStudentPassword returns the student's email.
func (s *student) GetStudentPassword() string {
	return s.studentPassword
}

// SetStudentID sets the student's ID.
func (s *student) SetStudentID(studentID string) {
	s.studentID = studentID
}

// SetStudentEmail sets the student's email.
func (s *student) SetStudentEmail(email string) {
	s.studentEmail = email
}

// SetStudentPassword sets the student's password.
func (s *student) SetPassword(password string) {
	s.studentPassword = password
}

// IsValidPassword returns true if received password is valid.
func (s *student) IsValidPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(s.studentPassword), []byte(password)) == nil
}

// HashPassword hashes the student's password.
func (s *student) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(s.studentPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	s.studentPassword = string(hash)
	return nil
}
