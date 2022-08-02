package handler

import (
	"errors"
	"net/http"

	"github.com/jeanmolossi/effective-eureka/src/core/lessons/facade"
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
	getSectionsFromModule domain.GetSectionsFromModule

	getLessonsInSection facade.GetLessonsInSection

	logger logger.Logger
}

func NewHandler(db *gorm.DB) *Handler {
	repo := repository.NewSectionRepository(db)
	createSectionInModule := usecase.NewCreateSectionInModule(repo)
	editSectionInfo := usecase.NewEditSectionInfo(repo)
	getSectionsFromModule := usecase.NewGetSectionsFromModule(repo)

	getLessonsInSection := facade.NewGetLessonsInSection(db)

	return &Handler{
		createSectionInModule,
		editSectionInfo,
		getSectionsFromModule,

		getLessonsInSection,

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
// @security access_token
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
// @security access_token
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
	).WithID(input.SectionID)

	editedSection, err := h.editSectionInfo.Run(section.Build())
	if err != nil {
		return ErrorHandler(c, err)
	}

	return c.JSON(http.StatusOK, NewHttpSectionOk(editedSection))
}

// GetSectionsFromModule is a function to get sections from module.
//
// @summary Get sections from module
// @description Get sections from module
// @tags sections
// @accept json
// @produce json
// @param moduleID path string true "Module ID"
// @param not_published query bool false "List not published courses too"
// @param fields query []string false "Only get that fields"
// @param page query uint16 false "Page"
// @param items_per_page query int false "Only get that fields"
// @success 200 {array} HttpSectionWithMeta
// @failure 400 {object} httputil.HttpBadRequestErr
// @failure 403 {object} httputil.HttpMissingAuthenticationErr
// @failure 404 {object} httputil.HttpNotFoundErr
// @failure 500 {object} httputil.HttpInternalServerErr
// @security access_token
// @router /module/{moduleID}/sections [get]
func (h *Handler) GetSectionsFromModule(c echo.Context) error {
	params := new(domain.GetSectionsParams)

	if err := c.Bind(params); err != nil {
		return ErrorHandler(c, err)
	}

	sections, err := h.getSectionsFromModule.Run(params)
	if err != nil {
		return ErrorHandler(c, err)
	}

	httpSections := make([]*HttpSectionOk, len(sections))
	for i, section := range sections {
		httpSections[i] = NewHttpSectionOk(section)
	}

	return c.JSON(http.StatusOK, NewHttpSectionWithMeta(httpSections, params))
}

// GetSectionLessons is a function to get section lessons.
//
// @summary Get lessons from section
// @description Get lessons from section
// @tags sections
// @accept json
// @produce json
// @param sectionID path string true "Section ID"
// @success 200 {array} HttpLessonOk
// @failure 400 {object} httputil.HttpBadRequestErr
// @failure 403 {object} httputil.HttpMissingAuthenticationErr
// @failure 404 {object} httputil.HttpNotFoundErr
// @failure 500 {object} httputil.HttpInternalServerErr
// @security access_token
// @router /section/{sectionID}/lessons [get]
func (h *Handler) GetSectionLessons(c echo.Context) error {
	sectionID := c.Param("sectionID")

	if sectionID == "" {
		return ErrorHandler(c, domain.NewBadRequestErr(
			errors.New("sectionID is required"),
		))
	}

	lessons, err := h.getLessonsInSection.Run(sectionID)
	if err != nil {
		return ErrorHandler(c, err)
	}

	return c.JSON(http.StatusOK, lessons)
}
