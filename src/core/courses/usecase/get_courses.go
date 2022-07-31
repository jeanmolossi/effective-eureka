package usecase

import (
	"github.com/jeanmolossi/effective-eureka/src/core/courses/domain"
	ormcondition "github.com/jeanmolossi/effective-eureka/src/pkg/orm_condition"
	"github.com/jeanmolossi/effective-eureka/src/pkg/paginator"
)

type getCourses struct {
	repo domain.CourseRepository
}

func NewGetCourses(repo domain.CourseRepository) domain.GetCourses {
	return &getCourses{repo}
}

// Run is the method to get courses.
func (g *getCourses) Run(params *domain.GetCoursesParams) ([]domain.Course, error) {
	paginator := paginator.NewPaginator()
	filters := ormcondition.NewFilterConditions()
	filters.AddCondition("course_published", true)

	if params != nil {
		filters.AddFields(params.Fields)

		if params.NotPublished {
			filters.RemoveCondition("course_published")
		}
	}

	return g.repo.GetCourses(filters, paginator)
}
