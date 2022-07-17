package domain_test

import (
	"testing"

	"github.com/jeanmolossi/effective-eureka/src/core/courses/domain"
	"github.com/stretchr/testify/assert"
)

func TestCourse(t *testing.T) {
	t.Run("should get correct values from course", func(t *testing.T) {
		course := domain.NewCourse("course-title", "course-thumb", "course-desc", false)
		course.SetCourseID("course-id")

		assert.Equal(t, "course-id", course.GetCourseID())
		assert.Equal(t, "course-title", course.GetCourseTitle())
		assert.Equal(t, "course-thumb", course.GetCourseThumb())
		assert.Equal(t, "course-desc", course.GetCourseDesc())
		assert.Equal(t, false, course.IsCoursePublished())
	})

	t.Run("should update original values", func(t *testing.T) {
		course := domain.NewCourse("course-title", "course-thumb", "course-desc", false)

		course.SetCourseTitle("new-title")
		course.SetCourseThumb("new-thumb")
		course.SetCourseDesc("new-desc")
		course.PublishCourse()

		assert.Equal(t, "new-title", course.GetCourseTitle())
		assert.Equal(t, "new-thumb", course.GetCourseThumb())
		assert.Equal(t, "new-desc", course.GetCourseDesc())
		assert.Equal(t, true, course.IsCoursePublished())
		course.UnpublishCourse()
		assert.Equal(t, false, course.IsCoursePublished())
	})
}
