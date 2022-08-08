// Package cmd is a package to manage the application commands.
//
// It documents and handle the server endpoints.
// Generate the swagger documentation
package cmd

import (
	"net/http"

	coursesHandler "github.com/jeanmolossi/effective-eureka/src/core/courses/handler"
	lessonsHandler "github.com/jeanmolossi/effective-eureka/src/core/lessons/handler"
	modulesHandler "github.com/jeanmolossi/effective-eureka/src/core/modules/handler"
	purchasesHandler "github.com/jeanmolossi/effective-eureka/src/core/purchases/handler"
	sectionsHandler "github.com/jeanmolossi/effective-eureka/src/core/sections/handler"
	shared "github.com/jeanmolossi/effective-eureka/src/core/shared"
	studentsHandler "github.com/jeanmolossi/effective-eureka/src/core/students/handler"

	"github.com/jeanmolossi/effective-eureka/src/cmd/httputil"
	"github.com/jeanmolossi/effective-eureka/src/pkg/auth"
	"github.com/jeanmolossi/effective-eureka/src/pkg/logger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func RunServer() {
	e := echo.New()
	e.Validator = shared.NewCustomValidator()

	dbConn := shared.NewDbConnection()
	err := dbConn.Connect()
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Middlewares
	e.Use(Cors())
	e.Use(middleware.RequestID())
	e.Use(logger.Middleware())
	e.Use(auth.Middleware(dbConn.DB()))

	// Courses
	ch := coursesHandler.NewHandler(dbConn.DB())
	e.POST("/course", ch.CreateCourse)
	e.GET("/course", ch.GetCourses)
	e.GET("/course/:courseID", ch.GetCourseByID)
	e.PUT("/course/:courseID", ch.EditCourseInfo)
	e.POST("/course/:courseID/module", ch.CreateModule)

	// Modules
	mh := modulesHandler.NewHandler(dbConn.DB())
	e.GET("/module/:moduleID", mh.GetModule)
	e.GET("/course/:courseID/modules", mh.GetModulesFromCourse)
	e.PUT("/module/:moduleID", mh.EditModuleInfo)

	// Sections
	sech := sectionsHandler.NewHandler(dbConn.DB())
	e.POST("/module/:moduleID/section", sech.CreateSectionInModule)
	e.PUT("/section/:sectionID", sech.EditSectionInfo)
	e.GET("/module/:moduleID/sections", sech.GetSectionsFromModule)
	e.GET("/section/:sectionID/lessons", sech.GetSectionLessons)

	// Lessons
	lh := lessonsHandler.NewHandler(dbConn.DB())
	e.POST("/section/:sectionID/lesson", lh.CreateLesson)
	e.PUT("/lesson/:lessonID", lh.EditLessonInfo)
	e.GET("/lesson/:lessonID", lh.GetLessonByID)

	// Authentication
	ah := auth.NewHandler(dbConn.DB())
	e.POST("/auth/login", ah.Login)
	e.POST("/auth/logout", ah.Logout)

	// Students
	sh := studentsHandler.NewHandler(dbConn.DB())
	e.POST("/students/register", sh.RegisterStudent)
	e.GET("/students/me", sh.GetMe)

	// Purchases
	ph := purchasesHandler.NewHandler(dbConn.DB())
	e.GET("/purchases", ph.GetPurchases)

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
