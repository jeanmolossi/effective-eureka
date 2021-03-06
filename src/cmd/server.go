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

// @title			Effective Eureka API
// @version			0.0.1
// @description		This is a catalog video manager API.
// @contact.name	Jean Molossi
// @contact.url		https://jeanmolossi.com.br/
// @contact.email	jean.carlosmolossi@gmail.com
// @license.name	Apache 2.0
// @license.url		http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:8080
// @BasePath		/
// @securityDefinitions.basic  BasicAuth
func RunServer() {
	e := echo.New()
	e.Validator = shared.NewCustomValidator()

	// Middlewares
	e.Use(middleware.RequestID())
	e.Use(logger.Middleware())

	// Routes
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/ping", Ping)

	// Courses
	e.POST("/course", CreateCourse)
	e.GET("/course/:courseID", GetCourseByID)

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

// CreateCourse is a endpoint to create a course.
// @Summary Course creation
// @tags courses
// @description Create a course
// @accept json
// @produce json
// @param course body input.CreateCourse true "Course object which will be created"
// @success 201 {object} coursesHandler.HttpCourseCreated
// @failure 400 {object} coursesHandler.HttpCreateCourseBadRequestErr
// @failure 500 {object} httputil.HttpInternalServerErr
// @router /course [post]
func CreateCourse(c echo.Context) error {
	h, err := coursesHandler.NewHandler()
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			httputil.HttpInternalServerErr{Message: err.Error()},
		)
	}

	return h.CreateCourse(c)
}

// GetCourseByID is a endpoint to get a course by ID.
// @summary Course retrieval
// @tags courses
// @description Get a course by ID
// @accept json
// @produce json
// @param courseID path string true "Course ID"
// @success 200 {object} coursesHandler.HttpCourseOk
// @failure 400 {object} coursesHandler.HttpCourseByIDBadRequestErr
// @failure 404 {object} coursesHandler.HttpCourseNotFoundErr
// @failure 500 {object} httputil.HttpInternalServerErr
// @router /course/:courseID [get]
func GetCourseByID(c echo.Context) error {
	h, err := coursesHandler.NewHandler()
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			httputil.HttpInternalServerErr{Message: err.Error()},
		)
	}

	return h.GetCourseByID(c)
}
