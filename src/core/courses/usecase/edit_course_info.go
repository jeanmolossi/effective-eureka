package usecase

import "github.com/jeanmolossi/effective-eureka/src/core/courses/domain"

// editCourseInfo is a usecase to edit a course info.
// It implements the domain.EditCourseInfo interface.
type editCourseInfo struct {
	repo domain.CourseRepository
}

func NewEditCourseInfo(repo domain.CourseRepository) domain.EditCourseInfo {
	return &editCourseInfo{repo}
}

func (e *editCourseInfo) Run(course domain.Course) (domain.Course, error) {
	return e.repo.Edit(course.GetCourseID(), e.updater(course))
}

func (e *editCourseInfo) updater(newCourseInfo domain.Course) domain.CourseUpdater {
	return func(currentCourseInfo domain.Course) (domain.Course, error) {
		if newCourseInfo.GetCourseTitle() != "" {
			currentCourseInfo.SetCourseTitle(newCourseInfo.GetCourseTitle())
		}

		if newCourseInfo.GetCourseThumb() != "" {
			currentCourseInfo.SetCourseThumb(newCourseInfo.GetCourseThumb())
		}

		if newCourseInfo.GetCourseDesc() != "" {
			currentCourseInfo.SetCourseDesc(newCourseInfo.GetCourseDesc())
		}

		return currentCourseInfo, nil
	}
}
