package usecase

import (
	"github.com/jeanmolossi/effective-eureka/src/core/courses/domain"
	"github.com/jeanmolossi/effective-eureka/src/core/courses/repository"
)

type getCourses struct {
	repo domain.CourseRepository
}

func NewGetCourses(repo domain.CourseRepository) domain.GetCourses {
	return &getCourses{repo}
}

// Run is the method to get courses.
func (g *getCourses) Run(params *domain.GetCoursesParams) ([]domain.Course, error) {
	filters := repository.Filters{
		ConditionMap: map[string]interface{}{
			"course_published": true,
		},
	}

	if params != nil {
		filters.Fields = params.Fields

		if params.NotPublished {
			filters.ConditionMap = nil
		}
	}

	paginator := repository.PagesConfig{
		Page:         1,
		ItemsPerPage: 10,
	}

	if params.Page > 0 {
		paginator.Page = params.Page
	}

	if params.ItemsPerPage > 0 {
		paginator.ItemsPerPage = params.ItemsPerPage
	}

	return g.repo.GetCourses(&filters, &paginator)
}
