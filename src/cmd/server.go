// Package cmd is a package to manage the application commands.
//
// It documents and handle the server endpoints.
// Generate the swagger documentation
package cmd

import (
	"net/http"

	coursesHandler "github.com/jeanmolossi/effective-eureka/src/core/courses/handler"
	shared "github.com/jeanmolossi/effective-eureka/src/core/shared"

	"github.com/jeanmolossi/effective-eureka/src/cmd/httputil"
	"github.com/jeanmolossi/effective-eureka/src/pkg/logger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func RunServer() {
	e := echo.New()
	e.Validator = shared.NewCustomValidator()

	// Middlewares
	e.Use(middleware.RequestID())
	e.Use(logger.Middleware())

	// Courses
	ch, err := coursesHandler.NewHandler()
	if err != nil {
		e.Logger.Fatal(err)
		return
	}
	e.POST("/course", ch.CreateCourse)
	e.GET("/course/:courseID", ch.GetCourseByID)
	e.PUT("/course/:courseID", ch.EditCourseInfo)

	// Routes
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/ping", Ping)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Ping is a health check endpoint.
//
// @summary Ping the server.
// @description A simple health check.
// @tags healthcheck
// @accept json
// @produce json
// @success 200 {object} httputil.PingOk
// @failure 500 {object} httputil.PingInternalServerErr
// @failure 502 {object} httputil.PingInternalServerErr
// @failure 503 {object} httputil.PingInternalServerErr
// @router /ping [get]
func Ping(c echo.Context) error {
	if c.Request().Method != "GET" {
		return c.JSON(http.StatusNotAcceptable, nil)
	}

	return c.JSON(http.StatusOK, httputil.PingOk{Message: "pong"})
}
