package handler

import (
	"errors"
	"net/http"

	"github.com/jeanmolossi/effective-eureka/src/core/purchases/data"
	"github.com/jeanmolossi/effective-eureka/src/core/shared"
	studentsDomain "github.com/jeanmolossi/effective-eureka/src/core/students/domain"
	studentsFacade "github.com/jeanmolossi/effective-eureka/src/core/students/facade"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Handler struct {
	// facade
	getMe studentsDomain.GetMe

	// data
	getPurchases data.GetPurchases
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		getMe:        studentsFacade.NewGetMe(db),
		getPurchases: data.NewGetPurchases(),
	}
}

// GetPurchases
//
// @summary Get sells
// @description Get sells
// @tags purchases
// @accept json
// @produce json
// @param transaction query string true "Transaction"
// @success 200 {object} HttpPurchasesOk
// @failure 400 {object} map[string]string
// @failure 404 {object} map[string]string
// @failure 500 {object} map[string]string
// @security access_token
// @router /purchases [get]
func (h *Handler) GetPurchases(c echo.Context) error {
	transaction := c.QueryParam("transaction")
	if transaction == "" {
		return shared.ErrorHandler(c, shared.NewBadRequestErr(
			errors.New("transaction is required"),
		))
	}

	input := new(studentsDomain.GetMeParams)
	input.StudentID = c.Get("studentID").(string)

	student, err := h.getMe.Run(input)
	if err != nil {
		return shared.ErrorHandler(c, err)
	}

	responseItems, err := h.getPurchases.Get(
		student.GetStudentEmail(),
		transaction,
	)
	if err != nil {
		return shared.ErrorHandler(c, err)
	}

	return c.JSON(http.StatusOK, NewHttpPurchaseOk(responseItems))
}
