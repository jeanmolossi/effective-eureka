package usecase_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jeanmolossi/effective-eureka/mocks"
	"github.com/jeanmolossi/effective-eureka/src/core/courses/domain"
	"github.com/jeanmolossi/effective-eureka/src/core/courses/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCourse(t *testing.T) {
	t.Run("should create a course", func(t *testing.T) {
		want := domain.NewCourse("title", "thumb", "desc", false)

		repo := new(mocks.CourseRepository)
		repo.On("Create", mock.Anything).Return(want, nil).Run(
			func(args mock.Arguments) {
				course := args.Get(0).(domain.Course)
				// Fake generated ID
				course.SetCourseID(uuid.NewString())
			})

		service := usecase.NewCreateCourse(repo)

		course := domain.NewCourse("title", "thumb", "desc", false)
		courseCreated, err := service.Run(course)

		assert.Nil(t, err)
		assert.NotEqual(t, course.GetCourseID(), courseCreated.GetCourseID())
		assert.Equal(t, course.GetCourseTitle(), courseCreated.GetCourseTitle())
		assert.Equal(t, course.GetCourseThumb(), courseCreated.GetCourseThumb())
		assert.Equal(t, course.GetCourseDesc(), courseCreated.GetCourseDesc())
	})
}
