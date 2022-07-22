package auth

import (
	"net/http"
	"regexp"
	"time"

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

			var tokenStr string
			token, err := c.Cookie("access_token")
			if err != nil {
				authToken := c.Request().Header.Get("Authorization")
				if authToken == "" {
					return c.JSON(http.StatusForbidden, httputil.HttpForbiddenErr{
						Message: "Missing authentication",
					})
				}

				tokenStr = authToken
			} else {
				tokenStr = token.Value
			}

			if !sessionProvider.IsValidSession(tokenStr) {
				c.SetCookie(&http.Cookie{
					Name:    "access_token",
					Path:    "/",
					Expires: time.Unix(0, 0),
				})

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
		`/auth/(login)$`,
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
