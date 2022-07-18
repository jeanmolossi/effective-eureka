package auth

import (
	"net/http"

	"github.com/jeanmolossi/effective-eureka/src/pkg/logger"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Handler struct {
	db       *gorm.DB
	provider *SessionProvider

	logger logger.Logger
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{db, NewSessionProvider(db), logger.NewLogger()}
}

func (h *Handler) Login(c echo.Context) error {
	h.logger.Debugln("Login")

	var credentials *LoginCredentials
	invalidCredentials := InvalidCredentialsErr{Message: "Invalid credentials"}

	if err := c.Bind(&credentials); err != nil {
		h.logger.Errorln("Error binding credentials", err)
		return c.JSON(http.StatusBadRequest, invalidCredentials)
	}

	if err := c.Validate(credentials); err != nil {
		h.logger.Errorln("Error validating credentials", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	token, err := h.provider.CreateSession(credentials.Username, credentials.Password)
	if err != nil {
		h.logger.Errorln("error creating session", err)
		return c.JSON(http.StatusBadRequest, invalidCredentials)
	}

	c.SetCookie(&http.Cookie{
		Name:    "access_token",
		Value:   token.Hash(),
		Path:    "/",
		Expires: token.Expiration,
	})

	return c.JSON(http.StatusOK, map[string]string{"access_token": token.Hash()})
}

func (h *Handler) Logout(c echo.Context) error {
	return nil
}
