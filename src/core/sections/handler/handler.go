package handler

import (
	"net/http"

	"github.com/jeanmolossi/effective-eureka/src/core/sections/domain"
	"github.com/jeanmolossi/effective-eureka/src/core/sections/factory"
	"github.com/jeanmolossi/effective-eureka/src/core/sections/input"
	"github.com/jeanmolossi/effective-eureka/src/core/sections/repository"
	"github.com/jeanmolossi/effective-eureka/src/core/sections/usecase"
	"github.com/jeanmolossi/effective-eureka/src/pkg/logger"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Handler struct {
	createSectionInModule domain.CreateSectionInModule
	editSectionInfo       domain.EditSectionInfo

	logger logger.Logger
}

func NewHandler(db *gorm.DB) *Handler {
	repo := repository.NewSectionRepository(db)
	createSectionInModule := usecase.NewCreateSectionInModule(repo)
	editSectionInfo := usecase.NewEditSectionInfo(repo)

	return &Handler{
		createSectionInModule,
		editSectionInfo,

		logger.NewLogger(),
	}
}

// CreateSectionInModule is a function to create a section in module.
//
// @summary Create a section in module
// @description Create a section in module
// @tags sections
// @accept json
// @produce json
// @param moduleID path string true "Module ID"
// @param section body input.CreateSection true "Section data"
// @success 201 {object} HttpSectionCreated
// @failure 400 {object} httputil.HttpBadRequestErr
// @failure 403 {object} httputil.HttpMissingAuthenticationErr
// @failure 404 {object} httputil.HttpNotFoundErr
// @failure 500 {object} httputil.HttpInternalServerErr
// @router /module/{moduleID}/section [post]
func (h *Handler) CreateSectionInModule(c echo.Context) error {
	input := new(input.CreateSection)

	if err := c.Bind(input); err != nil {
		return ErrorHandler(c, err)
	}

	if err := c.Validate(input); err != nil {
		return ErrorHandler(c, err)
	}

	section := factory.NewSection().CreateSection(
		input.ModuleID,
		input.Title,
		input.Index,
		input.Published,
		nil, nil,
	)

	createdSection, err := h.createSectionInModule.Run(section.Build())
	if err != nil {
		return ErrorHandler(c, err)
	}

	return c.JSON(http.StatusCreated, NewHttpSectionCreated(createdSection))
}

// EditSection is a function to edit a section.
//
// @summary Edit a section
// @description Edit a section
// @tags sections
// @accept json
// @produce json
// @param sectionID path string true "Section ID"
// @param section body input.EditSection true "Section data"
// @success 200 {object} HttpSectionOk
// @failure 400 {object} httputil.HttpBadRequestErr
// @failure 403 {object} httputil.HttpMissingAuthenticationErr
// @failure 404 {object} httputil.HttpNotFoundErr
// @failure 500 {object} httputil.HttpInternalServerErr
// @router /section/{sectionID} [put]
func (h *Handler) EditSectionInfo(c echo.Context) error {
	input := new(input.EditSection)

	if err := c.Bind(input); err != nil {
		return ErrorHandler(c, err)
	}

	if err := c.Validate(input); err != nil {
		return ErrorHandler(c, err)
	}

	section := factory.NewSection().CreateSection(
		input.ModuleID,
		input.Title,
		input.Index,
		input.Published,
		nil, nil,
	)

	editedSection, err := h.editSectionInfo.Run(section.Build())
	if err != nil {
		return ErrorHandler(c, err)
	}

	return c.JSON(http.StatusOK, NewHttpSectionOk(editedSection))
}
