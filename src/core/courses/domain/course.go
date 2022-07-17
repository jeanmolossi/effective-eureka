package domain

// course is a struct who implements Course interface.
type course struct {
	courseID        string
	courseTitle     string
	courseThumb     string
	courseDesc      string
	coursePublished bool
}

// NewCourse returns a new instance of Course.
func NewCourse(courseTitle, courseThumb, courseDesc string, published bool) Course {
	return &course{
		courseID:        "",
		courseTitle:     courseTitle,
		courseThumb:     courseThumb,
		courseDesc:      courseDesc,
		coursePublished: published,
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

// GetCourseDesc returns the course description.
func (c *course) GetCourseDesc() string {
	return c.courseDesc
}

// IsCoursePublished returns true if the course is published.
func (c *course) IsCoursePublished() bool {
	return c.coursePublished
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

// SetCourseDesc sets the course description.
func (c *course) SetCourseDesc(courseDesc string) {
	c.courseDesc = courseDesc
}

// PublishCourse publishes the course.
func (c *course) PublishCourse() {
	c.coursePublished = true
}

// UnpublishCourse unpublishes the course.
func (c *course) UnpublishCourse() {
	c.coursePublished = false
}
