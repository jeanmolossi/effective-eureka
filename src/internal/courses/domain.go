package courses

// Course is a interface who provides methods to manage courses.
type Course interface {
	// GetCourseID returns the course ID.
	GetCourseID() string
	// GetCourseTitle returns the course title.
	GetCourseTitle() string
	// GetCourseThumb returns the course thumb.
	GetCourseThumb() string

	// SetCourseID sets the course ID.
	SetCourseID(courseID string)
	// SetCourseTitle sets the course title.
	SetCourseTitle(courseTitle string)
	// SetCourseThumb sets the course thumb.
	SetCourseThumb(courseThumb string)
}

// course is a struct who implements Course interface.
type course struct {
	courseID    string
	courseTitle string
	courseThumb string
}

// NewCourse returns a new instance of Course.
func NewCourse() Course {
	return &course{
		courseID:    "",
		courseTitle: "",
		courseThumb: "",
	}
}

// GetCourseID returns the course ID.
func (c *course) GetCourseID() string {
	return c.courseID
}

// GetCourseTitle returns the course title.
func (c *course) GetCourseTitle() string {
	return c.courseTitle
}

// GetCourseThumb returns the course thumb.
func (c *course) GetCourseThumb() string {
	return c.courseThumb
}

// SetCourseID sets the course ID.
func (c *course) SetCourseID(courseID string) {
	c.courseID = courseID
}

// SetCourseTitle sets the course title.
func (c *course) SetCourseTitle(courseTitle string) {
	c.courseTitle = courseTitle
}

// SetCourseThumb sets the course thumb.
func (c *course) SetCourseThumb(courseThumb string) {
	c.courseThumb = courseThumb
}
