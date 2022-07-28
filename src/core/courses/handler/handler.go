// Package handler is a package to manage courses in Go.
package handler

import (
	"net/http"

	"github.com/jeanmolossi/effective-eureka/src/cmd/httputil"
	"github.com/jeanmolossi/effective-eureka/src/core/courses/domain"
	"github.com/jeanmolossi/effective-eureka/src/core/courses/factory"
	"github.com/jeanmolossi/effective-eureka/src/core/courses/input"
	"github.com/jeanmolossi/effective-eureka/src/core/courses/repository"
	"github.com/jeanmolossi/effective-eureka/src/core/courses/usecase"
	"github.com/jeanmolossi/effective-eureka/src/core/modules/facade"
	"github.com/jeanmolossi/effective-eureka/src/core/shared"
	"github.com/jeanmolossi/effective-eureka/src/pkg/logger"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Handler is a struct to manage courses usecases.
type Handler struct {
	getCourseByID  domain.GetCourseByID
	createCourse   domain.CreateCourse
	editCourseInfo domain.EditCourseInfo
	getCourses     domain.GetCourses

	createModule facade.CreateModule

	logger logger.Logger
}

// NewHandler is a factory method to create a Handler.
func NewHandler(db *gorm.DB) *Handler {
	repo := repository.NewRepository(db)
	getCourseByID := usecase.NewGetCourseByID(repo)
	createCourse := usecase.NewCreateCourse(repo)
	editCourseInfo := usecase.NewEditCourseInfo(repo)
	getCourses := usecase.NewGetCourses(repo)

	createModule := facade.NewCreateModule(db)

	return &Handler{
		getCourseByID,
		createCourse,
		editCourseInfo,
		getCourses,

		createModule,

		logger.NewLogger(),
	}
}

// CreateCourse is a endpoint to create a course.
// @Summary Course creation
// @tags courses
// @description Create a course
// @accept json
// @produce json
// @param course body input.CreateCourse true "Course object which will be created"
// @success 201 {object} HttpCourseCreated
// @failure 400 {object} HttpCreateCourseBadRequestErr
// @failure 500 {object} httputil.HttpInternalServerErr
// @security access_token
// @router /course [post]
func (h *Handler) CreateCourse(c echo.Context) error {
	var input *input.CreateCourse

	// Bind input with input struct we expect
	err := c.Bind(&input)
	if err != nil {
		h.logger.Errorln("Error binding input", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Validate input with input struct we expect
	err = c.Validate(input)
	if err != nil {
		h.logger.Errorln("Error validating input", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	// Create course using factory.
	course := factory.NewCourse().CreateCourse(
		input.Title,
		input.Thumbnail,
		input.Description,
		input.Published,
	)

	newCourse, err := h.createCourse.Run(course.Build())
	if err != nil {
		h.logger.Errorln("Error running usecase", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, NewHttpCourseCreated(newCourse))
}

// GetCourseByID is a endpoint to get a course by ID.
// @summary Course retrieval
// @tags courses
// @description Get a course by ID
// @accept json
// @produce json
// @param courseID path string true "Course ID"
// @success 200 {object} HttpCourseOk
// @failure 400 {object} HttpCourseByIDBadRequestErr
// @failure 404 {object} HttpCourseNotFoundErr
// @failure 500 {object} httputil.HttpInternalServerErr
// @security access_token
// @router /course/{courseID} [get]
func (h *Handler) GetCourseByID(c echo.Context) error {
	courseID := c.Param("courseID")

	course, err := h.getCourseByID.Run(courseID)
	if err != nil {
		h.logger.Errorln("Error running usecase", err)
		return shared.ErrorHandler(c, err)
	}

	return c.JSON(http.StatusOK, NewHttpCourseOk(course))
}

// EditCourseInfo is a endpoint to edit a course.
// @Summary Course edition
// @tags courses
// @description Edit a course basic information
// @accept json
// @produce json
// @param courseID path string true "Course ID"
// @param course body input.EditCourseInfo true "Course object which will be edited"
// @success 200 {object} HttpCourseOk
// @failure 400 {object} HttpEditCourseInfoBadRequestErr
// @failure 404 {object} HttpCourseNotFoundErr
// @failure 500 {object} httputil.HttpInternalServerErr
// @security access_token
// @router /course/{courseID} [put]
func (h *Handler) EditCourseInfo(c echo.Context) error {
	courseID := c.Param("courseID")
	if courseID == "" {
		h.logger.Errorln("courseID is empty")
		return c.JSON(http.StatusBadRequest, HttpCourseByIDBadRequestErr{"Missing course_id param"})
	}

	var input *input.EditCourseInfo
	err := c.Bind(&input)
	if err != nil {
		h.logger.Errorln("Error binding input", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	err = c.Validate(input)
	if err != nil {
		h.logger.Errorln("validation error", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	// Create course using factory constructor.
	course := factory.NewCourse().CreateCourse(
		input.Title,
		input.Thumbnail,
		input.Description,
		false,
	).WithID(courseID)

	updatedCourse, err := h.editCourseInfo.Run(course.Build())
	if err != nil {
		h.logger.Errorln("Error running usecase", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, NewHttpCourseOk(updatedCourse))
}

// CreateModule is a endpoint to create a module.
//
// @Summary Module creation
// @tags courses
// @description Create a module
// @accept json
// @produce json
// @param module body input.CreateModule true "Module object which will be created"
// @param courseID path string true "Course ID"
// @success 201 {object} HttpModuleCreated
// @failure 400 {object} httputil.HttpBadRequestErr
// @failure 403 {object} httputil.HttpMissingAuthenticationErr
// @failure 500 {object} httputil.HttpInternalServerErr
// @security access_token
// @router /course/{courseID}/module [post]
func (h *Handler) CreateModule(c echo.Context) error {
	// Bind input with input struct we expect
	err := c.Bind(h.createModule.Input())
	if err != nil {
		h.logger.Errorln("Error binding input", err)
		return c.JSON(http.StatusInternalServerError, httputil.HttpInternalServerErr{
			Message: err.Error(),
		})
	}

	// Validate input with input struct we expect
	err = c.Validate(h.createModule.Input())
	if err != nil {
		h.logger.Errorln("Error validating input", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	newModule, err := h.createModule.Run()
	if err != nil {
		h.logger.Errorln("Error running usecase", err)
		return c.JSON(http.StatusInternalServerError, httputil.HttpInternalServerErr{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, newModule)
}

// GetCourses is a endpoint to get a course list.
//
// @summary Course retrieval
// @tags courses
// @description Get a course list
// @accept json
// @produce json
// @param not_published query bool false "List not published courses too"
// @param fields query []string false "Only get that fields"
// @success 200 {object} []HttpCourseOk
// @failure 400 {object} HttpCourseByIDBadRequestErr
// @failure 404 {object} HttpCourseNotFoundErr
// @failure 500 {object} httputil.HttpInternalServerErr
// @security access_token
// @router /course [get]
func (h *Handler) GetCourses(c echo.Context) error {
	input := new(domain.GetCoursesParams)

	if err := c.Bind(input); err != nil {
		h.logger.Errorln("Error binding input", err)
		return c.JSON(http.StatusInternalServerError, httputil.HttpInternalServerErr{
			Message: err.Error(),
		})
	}

	courses, err := h.getCourses.Run(input)
	if err != nil {
		h.logger.Errorln("Error running usecase", err)
		return shared.ErrorHandler(c, err)
	}

	httpCourses := make([]*HttpCourseOk, len(courses))
	for i, course := range courses {
		httpCourses[i] = NewHttpCourseOk(course)
	}

	return c.JSON(http.StatusOK, httpCourses)
}
