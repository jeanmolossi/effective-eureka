package auth

import (
	"net/http"
	"time"

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

// Login handles the login request.
// It validates the credentials and creates a new session.
// It returns a JSON with the access token.
//
// @Summary Login
// @Description Login
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body LoginCredentials true "Login credentials"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/login [post]
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

// Logout handles the logout request.
// It deletes the session.
// It returns a JSON with the access token.
//
// @Summary Logout
// @Description Login
// @Tags auth
// @Accept json
// @Produce json
// @Param  cookies header string true "Access token cookie"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/logout [post]
func (h *Handler) Logout(c echo.Context) error {
	accessToken, err := c.Cookie("access_token")
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]string{"message": "No access token found"})
	}

	_, sessionID, err := Decode(accessToken.Value)
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]string{"message": err.Error()})
	}

	err = h.provider.DeleteSession(sessionID)
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]string{"message": err.Error()})
	}

	accessToken.Expires = time.Now().UTC().Local().Add(time.Hour * 24 * -7)
	c.SetCookie(&http.Cookie{
		Name:    "access_token",
		Value:   "null",
		Expires: accessToken.Expires,
	})

	return c.JSON(http.StatusAccepted, map[string]string{"message": "Logged out"})
}
