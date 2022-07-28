package domain

type GetCoursesParams struct {
	Fields       []string `query:"fields"`
	NotPublished bool     `query:"not_published"`
}

// GetCourses is a interface who provides methods to get courses.
type GetCourses interface {
	// Run is the method to get courses.
	Run(conditions *GetCoursesParams) ([]Course, error)
}
