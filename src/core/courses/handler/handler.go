// Package handler is a package to manage courses in Go.
package handler

import (
	"net/http"

	"github.com/jeanmolossi/effective-eureka/src/core/courses/domain"
	"github.com/jeanmolossi/effective-eureka/src/core/courses/input"
	"github.com/jeanmolossi/effective-eureka/src/core/courses/repository"
	"github.com/jeanmolossi/effective-eureka/src/core/courses/usecase"
	shared "github.com/jeanmolossi/effective-eureka/src/core/shared"
	"github.com/labstack/echo/v4"
)

// Handler is a struct to manage courses usecases.
type Handler struct {
	getCourseByID domain.GetCourseByID
	createCourse  domain.CreateCourse
}

// NewHandler is a factory method to create a Handler.
func NewHandler() (*Handler, error) {
	dbConn := shared.NewDbConnection()
	err := dbConn.Connect()
	if err != nil {
		return nil, err
	}

	repo := repository.NewRepository(dbConn.DB())
	getCourseByID := usecase.NewGetCourseByID(repo)
	createCourse := usecase.NewCreateCourse(repo)

	return &Handler{
		getCourseByID,
		createCourse,
	}, nil
}

// Create course endpoint
func (h *Handler) CreateCourse(c echo.Context) error {
	var input *input.CreateCourse

	// Bind input with input struct we expect
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Validate input with input struct we expect
	err = c.Validate(input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Create course using domain constructor.
	// [ ] - Should better use a factory method to create a course.
	course := domain.NewCourse(
		input.Title,
		input.Thumbnail,
		input.Description,
		input.Published,
	)

	newCourse, err := h.createCourse.Run(course)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, NewHttpCourseCreated(newCourse))
}

// GetCourseByID endpoint
func (h *Handler) GetCourseByID(c echo.Context) error {
	courseID := c.Param("courseID")

	course, err := h.getCourseByID.Run(courseID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, NewHttpCourseOk(course))
}
