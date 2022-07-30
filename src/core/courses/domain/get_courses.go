package domain

type GetCoursesParams struct {
	Fields       []string `query:"fields"`
	NotPublished bool     `query:"not_published"`
	Page         uint16   `query:"page"`
	ItemsPerPage int      `query:"items_per_page"`
}

// GetCourses is a interface who provides methods to get courses.
type GetCourses interface {
	// Run is the method to get courses.
	Run(conditions *GetCoursesParams) ([]Course, error)
}
