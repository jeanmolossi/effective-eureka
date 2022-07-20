package auth

import (
	"log"
	"net/http"
	"regexp"

	"github.com/jeanmolossi/effective-eureka/src/cmd/httputil"
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
				log.Println(err)
				return c.JSON(http.StatusForbidden, httputil.HttpForbiddenErr{
					Message: err.Error(),
				})
			}

			if !sessionProvider.IsValidSession(token.Value) {
				return c.JSON(http.StatusForbidden, httputil.HttpForbiddenErr{
					Message: "Missing authentication",
				})
			}

			return next(c)
		}
	}
}

func shouldIgnorePath(path string) bool {
	middlewareShouldIgnorePaths := []string{
		`^/ping$`,
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
