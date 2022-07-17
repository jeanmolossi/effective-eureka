package repository

import (
	"time"

	"github.com/jeanmolossi/effective-eureka/src/core/courses/domain"
)

// ModelToDomain converts a CourseModel to a CourseDomain.
func ModelToDomain(model *CourseModel) domain.Course {
	if model == nil {
		return domain.NewCourse("", "", "", false)
	}

	domainCourse := domain.NewCourse(
		model.CourseTitle,
		model.CourseThumb,
		model.CourseDescription,
		model.CoursePublished)

	domainCourse.SetCourseID(model.CourseID)

	return domainCourse
}

// DomainToModel converts a CourseDomain to a CourseModel.
func DomainToModel(course domain.Course, createdAt, updatedAt *time.Time) *CourseModel {
	var createdAtVal, updatedAtVal time.Time

	// If createdAt is not nil, set it to the current time.
	if createdAt != nil {
		createdAtVal = *createdAt
	}

	// If updatedAt is not nil, set it to the current time.
	if updatedAt != nil {
		updatedAtVal = *updatedAt
	}

	return &CourseModel{
		CourseID:          course.GetCourseID(),
		CourseTitle:       course.GetCourseTitle(),
		CourseThumb:       course.GetCourseThumb(),
		CourseDescription: course.GetCourseDesc(),
		CoursePublished:   course.IsCoursePublished(),
		CourseCreatedAt:   createdAtVal,
		CourseUpdatedAt:   updatedAtVal,
	}
}
