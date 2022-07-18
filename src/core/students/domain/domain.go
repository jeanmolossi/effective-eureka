package domain

// Student represents a student in the system.
type Student interface {
	// GetStudentID returns the student's ID.
	GetStudentID() string
	// GetStudentEmail returns the student's email.
	GetStudentEmail() string
	// GetStudentPassword returns the student's password.
	GetStudentPassword() string

	// SetStudentID sets the student's ID.
	SetStudentID(string)
	// SetStudentEmail sets the student's email.
	SetStudentEmail(string)
	// SetStudentPassword sets the student's password.
	SetPassword(string)

	// IsValidPassword returns true if received password is valid.
	IsValidPassword(string) bool
	// HashPassword hashes the student's password.
	HashPassword() error
}

// RegisterStudent registers a new student.
type RegisterStudent interface {
	// Run execute registration a new student.
	Run(student Student) (Student, error)
}

type StudentRepository interface {
	// GetStudentByID returns a student by ID.
	GetStudentByID(studentID string) (Student, error)
	// GetStudentByEmail returns a student by email.
	GetStudentByEmail(studentEmail string) (Student, error)

	// CreateStudent creates a new student.
	CreateStudent(student Student) (Student, error)
	// UpdateStudent updates a student.
	UpdateStudent(student Student) (Student, error)
	// DeleteStudent deletes a student.
	DeleteStudent(studentID string) error
}
