package domain

// Course is a interface who provides methods to manage courses.
type Course interface {
	// GetCourseID returns the course ID.
	GetCourseID() string
	// GetCourseTitle returns the course title.
	GetCourseTitle() string
	// GetCourseThumb returns the course thumb.
	GetCourseThumb() string
	// GetCourseDesc returns the course description.
	GetCourseDesc() string
	// IsCoursePublished returns true if the course is published.
	IsCoursePublished() bool

	// SetCourseID sets the course ID.
	SetCourseID(courseID string)
	// SetCourseTitle sets the course title.
	SetCourseTitle(courseTitle string)
	// SetCourseThumb sets the course thumb.
	SetCourseThumb(courseThumb string)
	// SetCourseDesc sets the course description.
	SetCourseDesc(courseDesc string)
	// PublishCourse publishes the course.
	PublishCourse()
	// UnpublishCourse unpublishes the course.
	UnpublishCourse()
}

type GetCourseByID interface {
	Run(courseID string) (Course, error)
}

type GetCoursesFromStudent interface {
	Run(studentID string) ([]Course, error)
}

// CreateCourse is interface segregation to create a course.
type CreateCourse interface {
	// Run is the method with handles application to create a course.
	Run(course Course) (Course, error)
}

type EditCourseInfo interface {
	Run(course Course) (Course, error)
}

type PublishCourse interface {
	Run(courseID string) error
}

type UnpublishCourse interface {
	Run(courseID string) error
}

type CourseRepository interface {
	GetByID(courseID string) (Course, error)
	GetByStudentID(studentID string) ([]Course, error)
	Create(course Course) (Course, error)
	Edit(course Course) (Course, error)
}
