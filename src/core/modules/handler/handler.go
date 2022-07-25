// Package handler is a package to manage modules in Go.
package handler

import (
	"net/http"

	"github.com/jeanmolossi/effective-eureka/src/cmd/httputil"
	"github.com/jeanmolossi/effective-eureka/src/core/modules/domain"
	"github.com/jeanmolossi/effective-eureka/src/core/modules/factory"
	"github.com/jeanmolossi/effective-eureka/src/core/modules/input"
	"github.com/jeanmolossi/effective-eureka/src/core/modules/repository"
	"github.com/jeanmolossi/effective-eureka/src/core/modules/usecase"
	"github.com/jeanmolossi/effective-eureka/src/pkg/logger"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Handler is a struct to manage courses usecases.
type Handler struct {
	createModule         domain.CreateModuleInCourse
	getModuleByID        domain.GetModuleByID
	editModuleInfo       domain.EditModuleInfo
	getModulesFromCourse domain.GetModuleFromCourse

	logger logger.Logger
}

// NewHandler is a factory method to create a Handler.
func NewHandler(db *gorm.DB) *Handler {
	repo := repository.NewRepository(db)
	createModule := usecase.NewCreateModuleInCourse(repo)
	getModuleByID := usecase.NewGetModuleByID(repo)
	editModuleInfo := usecase.NewEditModuleInfo(repo)
	getModulesFromCourse := usecase.NewGetModuleFromCourse(repo)

	return &Handler{
		createModule,
		getModuleByID,
		editModuleInfo,
		getModulesFromCourse,

		logger.NewLogger(),
	}
}

// GetModule is a endpoint to get a module.
//
// @Summary Module retrieval
// @tags modules
// @description Get a module
// @accept json
// @produce json
// @param moduleID path string true "Module ID"
// @success 200 {object} HttpModuleOk
// @failure 400 {object} HttpBadRequestErr
// @failure 403 {object} httputil.HttpMissingAuthenticationErr
// @failure 404 {object} httputil.HttpNotFoundErr
// @failure 500 {object} httputil.HttpInternalServerErr
// @security access_token
// @router /module/{moduleID} [get]
func (h *Handler) GetModule(c echo.Context) error {
	moduleID := c.Param("moduleID")

	if moduleID == "" {
		return c.JSON(http.StatusBadRequest, httputil.HttpBadRequestErr{
			Message: "module id is required",
		})
	}

	module, err := h.getModuleByID.Run(moduleID)
	if err != nil {
		h.logger.Errorln("Error running usecase", err)
		return c.JSON(http.StatusNotFound, httputil.HttpNotFoundErr{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, NewHttpModuleOk(module))
}

// EditModuleInfo is a endpoint to edit a module.
//
// @Summary Module retrieval
// @tags modules
// @description Edit a module
// @accept json
// @produce json
// @param moduleID path string true "Module ID"
// @param module body input.EditModuleInfo true "Module object which will be updated"
// @success 200 {object} HttpModuleOk
// @failure 400 {object} HttpBadRequestErr
// @failure 403 {object} httputil.HttpMissingAuthenticationErr
// @failure 404 {object} httputil.HttpNotFoundErr
// @failure 500 {object} httputil.HttpInternalServerErr
// @security access_token
// @router /module/{moduleID} [put]
func (h *Handler) EditModuleInfo(c echo.Context) error {
	input := new(input.EditModuleInfo)

	// Bind input with input struct we expect
	err := c.Bind(input)
	if err != nil {
		h.logger.Errorln("Error binding input", err)
		return c.JSON(http.StatusInternalServerError, httputil.HttpInternalServerErr{
			Message: err.Error(),
		})
	}

	// Validate input with input struct we expect
	err = c.Validate(input)
	if err != nil {
		h.logger.Errorln("Error validating input", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	// Create module using factory.
	module := factory.NewModule().CreateModule(
		input.CourseID,
		input.Title,
		input.Thumbnail,
		input.Description,
		input.Published,
	).WithID(input.ModuleID)

	newModule, err := h.editModuleInfo.Run(module.Build())
	if err != nil {
		h.logger.Errorln("Error running usecase", err)
		return ErrorHandler(c, err)
	}

	return c.JSON(http.StatusOK, NewHttpModuleOk(newModule))
}

// GetModulesFromCourse is a endpoint to get all modules from a course.
//
// @Summary Module retrieval
// @tags modules
// @description Get all modules from a course
// @accept json
// @produce json
// @param courseID path string true "Course ID"
// @success 200 {array} HttpModuleOk
// @failure 400 {object} HttpBadRequestErr
// @failure 403 {object} httputil.HttpMissingAuthenticationErr
// @failure 404 {object} httputil.HttpNotFoundErr
// @failure 500 {object} httputil.HttpInternalServerErr
// @security access_token
// @router /course/{courseID}/modules [get]
func (h *Handler) GetModulesFromCourse(c echo.Context) error {
	courseID := c.Param("courseID")

	if courseID == "" {
		return c.JSON(http.StatusBadRequest, httputil.HttpBadRequestErr{
			Message: "course id is required",
		})
	}

	modules, err := h.getModulesFromCourse.Run(courseID)
	if err != nil {
		h.logger.Errorln("Error running usecase", err)
		return ErrorHandler(c, err)
	}

	httpModules := make([]*HttpModuleOk, len(modules))
	for i, module := range modules {
		httpModules[i] = NewHttpModuleOk(module)
	}

	return c.JSON(http.StatusOK, httpModules)
}
