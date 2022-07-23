package handler

import (
	"net/http"

	"github.com/jeanmolossi/effective-eureka/src/cmd/httputil"
	"github.com/jeanmolossi/effective-eureka/src/core/students/domain"
	"github.com/jeanmolossi/effective-eureka/src/core/students/factory"
	"github.com/jeanmolossi/effective-eureka/src/core/students/input"
	"github.com/jeanmolossi/effective-eureka/src/core/students/repository"
	"github.com/jeanmolossi/effective-eureka/src/core/students/usecase"
	"github.com/jeanmolossi/effective-eureka/src/pkg/auth"
	"github.com/jeanmolossi/effective-eureka/src/pkg/logger"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Handler struct {
	registerStudent domain.RegisterStudent
	getMe           domain.GetMe

	logger logger.Logger
}

func NewHandler(db *gorm.DB) *Handler {
	repo := repository.NewStudent(db)
	registerStudent := usecase.NewRegisterStudent(repo)
	getMeStudent := usecase.NewGetMe(repo,
		auth.NewSessionProvider(db),
	)

	return &Handler{
		registerStudent,
		getMeStudent,

		logger.NewLogger(),
	}
}

// RegisterStudent registers a new student.
// @Summary Register a new student.
// @Description Register a new student.
// @Tags students
// @Accept json
// @Produce json
// @Param student body input.StudentInfo true "Student information"
// @Success 201 {object} HttpStudentRegistered
// @Failure 400 {object} HttpCreateStudentBadRequestErr
// @Failure 500 {object} HttpStudentInternalServerErr
// @Router /students/register [post]
func (h *Handler) RegisterStudent(c echo.Context) error {
	var studentInfo *input.StudentInfo
	if err := c.Bind(&studentInfo); err != nil {
		h.logger.Errorln("error binding student info", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(studentInfo); err != nil {
		h.logger.Errorln("error validating student info", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	studentFactory := factory.NewStudent().CreateStudent(
		studentInfo.Email,
		studentInfo.Password,
	)

	registeredStudent, err := h.registerStudent.Run(studentFactory.Build())
	if err != nil {
		return c.JSON(http.StatusBadRequest, HttpCreateStudentBadRequestErr{
			Err: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, NewHttpStudentRegistered(registeredStudent))
}

// GetMe returns the auth user.
//
// @Summary Get auth student.
// @Description Get auth student.
// @Tags students
// @Accept json
// @Produce json
// @Success 200 {object} HttpStudentRegistered
// @Failure 400 {object} HttpCreateStudentBadRequestErr
// @Failure 403 {object} HttpStudentForbiddenErr
// @Failure 500 {object} HttpStudentInternalServerErr
// @Security access_token
// @Router /students/me [get]
func (h *Handler) GetMe(c echo.Context) error {
	student, err := h.getMe.Run(getAuthToken(c))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, httputil.HttpUnauthorizedErr{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, NewHttpStudent(student))
}

func getAuthToken(c echo.Context) string {
	cookieToken, err := c.Cookie("access_token")
	if err != nil {
		if authToken := c.Request().Header.Get("Authorization"); authToken != "" {
			return authToken
		}

		return ""
	}

	if cookieToken.Value != "" {
		return cookieToken.Value
	}

	return ""
}
