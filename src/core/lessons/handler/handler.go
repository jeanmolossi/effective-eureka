package handler

import (
	"net/http"

	"github.com/jeanmolossi/effective-eureka/src/core/lessons/domain"
	"github.com/jeanmolossi/effective-eureka/src/core/lessons/factory"
	"github.com/jeanmolossi/effective-eureka/src/core/lessons/input"
	"github.com/jeanmolossi/effective-eureka/src/core/lessons/repository"
	"github.com/jeanmolossi/effective-eureka/src/core/lessons/usecase"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Handler struct {
	addLessonInSection domain.AddLessonInSection
	editLesson         domain.EditLessonInfo
}

func NewHandler(db *gorm.DB) *Handler {
	repo := repository.NewRepository(db)
	addLessonInSection := usecase.NewAddLessonInSection(repo)
	editLesson := usecase.NewEditLessonInfo(repo)

	return &Handler{
		addLessonInSection,
		editLesson,
	}
}

// CreateLesson is a endpoint to create a lesson.
//
// @summary Create a lesson
// @description Create a lesson
// @tags lessons
// @accept json
// @produce json
// @param lesson body input.CreateLesson true "Lesson"
// @param sectionID path string true "Section ID"
// @success 201 {object} HttpLessonCreated
// @failure 400 {object} httputil.HttpBadRequestErr
// @failure 403 {object} httputil.HttpMissingAuthenticationErr
// @failure 404 {object} httputil.HttpNotFoundErr
// @failure 500 {object} httputil.HttpInternalServerErr
// @security access_token
// @router /section/{sectionID}/lesson [post]
func (h *Handler) CreateLesson(c echo.Context) error {
	input := new(input.CreateLesson)

	if err := c.Bind(input); err != nil {
		return ErrorHandler(c, err)
	}

	if err := c.Validate(input); err != nil {
		return ErrorHandler(c, err)
	}

	lesson := factory.NewLesson().CreateLesson(
		input.Title,
		input.Description,
		input.Thumbnail,
		input.Index,
		input.Published,
		nil, nil,
	).WithSectionID(input.SectionID)

	newLesson, err := h.addLessonInSection.AddLesson(lesson.Build())
	if err != nil {
		return ErrorHandler(c, err)
	}

	return c.JSON(http.StatusCreated, NewHttpLessonCreated(newLesson))
}

// EditLessonInfo is a endpoint to edit a lesson.
//
// @summary Edit a lesson
// @description Edit a lesson
// @tags lessons
// @accept json
// @produce json
// @param lesson body input.EditLessonInfo true "Lesson"
// @param lessonID path string true "Lesson ID"
// @success 200 {object} HttpLessonOk
// @failure 400 {object} httputil.HttpBadRequestErr
// @failure 403 {object} httputil.HttpMissingAuthenticationErr
// @failure 404 {object} httputil.HttpNotFoundErr
// @failure 500 {object} httputil.HttpInternalServerErr
// @security access_token
// @router /lesson/{lessonID} [put]
func (h *Handler) EditLessonInfo(c echo.Context) error {
	input := new(input.EditLessonInfo)

	if err := c.Bind(input); err != nil {
		return ErrorHandler(c, err)
	}

	if err := c.Validate(input); err != nil {
		return ErrorHandler(c, err)
	}

	lesson := factory.NewLesson().CreateLesson(
		input.Title,
		input.Description,
		input.Thumbnail,
		input.Index,
		input.Published,
		nil, nil,
	).WithLessonID(input.LessonID).WithSectionID(input.SectionID)

	updatedLesson, err := h.editLesson.EditLesson(lesson.Build())
	if err != nil {
		return ErrorHandler(c, err)
	}

	return c.JSON(http.StatusOK, NewHttpLessonOk(updatedLesson))
}
