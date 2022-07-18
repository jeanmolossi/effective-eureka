package handler

import (
	"net/http"

	"github.com/jeanmolossi/effective-eureka/src/core/students/domain"
	"github.com/jeanmolossi/effective-eureka/src/core/students/factory"
	"github.com/jeanmolossi/effective-eureka/src/core/students/input"
	"github.com/jeanmolossi/effective-eureka/src/core/students/repository"
	"github.com/jeanmolossi/effective-eureka/src/core/students/usecase"
	"github.com/jeanmolossi/effective-eureka/src/pkg/logger"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Handler struct {
	registerStudent domain.RegisterStudent

	logger logger.Logger
}

func NewHandler(db *gorm.DB) *Handler {
	repo := repository.NewStudent(db)
	registerStudent := usecase.NewRegisterStudent(repo)

	return &Handler{
		registerStudent,

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
