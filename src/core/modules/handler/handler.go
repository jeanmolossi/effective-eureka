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
	createModule  domain.CreateModuleInCourse
	getModuleByID domain.GetModuleByID

	logger logger.Logger
}

// NewHandler is a factory method to create a Handler.
func NewHandler(db *gorm.DB) *Handler {
	repo := repository.NewRepository(db)
	createModule := usecase.NewCreateModuleInCourse(repo)
	getModuleByID := usecase.NewGetModuleByID(repo)

	return &Handler{
		createModule,
		getModuleByID,

		logger.NewLogger(),
	}
}

// CreateModule is a endpoint to create a module.
// @Summary Module creation
// @tags modules
// @description Create a module
// @accept json
// @produce json
// @param module body input.CreateModule true "Module object which will be created"
// @param courseID path string true "Course ID"
// @success 201 {object} HttpModuleCreated
// @failure 400 {object} HttpBadRequestErr
// @failure 403 {object} httputil.HttpMissingAuthenticationErr
// @failure 500 {object} httputil.HttpInternalServerErr
// @security access_token
// @router /course/{courseID}/module [post]
func (h *Handler) CreateModule(c echo.Context) error {
	input := new(input.CreateModule)

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
	)

	newModule, err := h.createModule.Run(module.Build())
	if err != nil {
		h.logger.Errorln("Error running usecase", err)
		return c.JSON(http.StatusInternalServerError, httputil.HttpInternalServerErr{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, NewHttpModuleCreated(newModule))
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
