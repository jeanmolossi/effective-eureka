package cmd

import (
	_ "github.com/jeanmolossi/effective-eureka/docs"
	_ "github.com/swaggo/files"

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

	// Middlewares
	e.Use(middleware.RequestID())
	e.Use(logger.Middleware())

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
	return c.JSON(200, httputil.PingOk{Message: "pong"})
}
