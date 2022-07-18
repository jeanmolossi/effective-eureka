package auth

import (
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Middleware(db *gorm.DB) echo.MiddlewareFunc {
	sessionProvider := NewSessionProvider(db)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if shouldIgnorePath(c.Path()) {
				return next(c)
			}

			token, err := c.Cookie("session_token")
			if err != nil {
				return c.JSON(http.StatusForbidden, err)
			}

			if !sessionProvider.IsValidSession(token.Value) {
				return c.JSON(http.StatusForbidden, map[string]string{"error": "Missing authentication"})
			}

			return next(c)
		}
	}
}

func shouldIgnorePath(path string) bool {
	middlewareShouldIgnorePaths := []string{
		`/auth/(.*)$`,
		`/swagger/(.*)$`,
		`/students/(register)$`,
	}

	for _, p := range middlewareShouldIgnorePaths {
		match, err := regexp.MatchString(p, path)
		if err != nil {
			return false
		}

		if match {
			return true
		}
	}

	return false
}
