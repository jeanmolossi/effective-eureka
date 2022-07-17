// Package main is the entry point of the application.
//
// Handles the rest api and the event handlers.
package main

import (
	"github.com/jeanmolossi/effective-eureka/docs"
	"github.com/jeanmolossi/effective-eureka/src/cmd"

	_ "github.com/jeanmolossi/effective-eureka/docs"
	_ "github.com/swaggo/files"
)

// @termsOfService  github.com/jeanmolossi/effective-eureka/terms/
// @contact.name	Jean Molossi
// @contact.url		https://jeanmolossi.com.br/
// @contact.email	jean.carlosmolossi@gmail.com
// @license.name	Apache 2.0
// @license.url		github.com/jeanmolossi/effective-eureka/LICENSE
// @securityDefinitions.basic  BasicAuth
func main() {
	docs.SwaggerInfo.Title = "Effective Eureka API"
	docs.SwaggerInfo.Description = "This is a catalog video manager API."
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"

	cmd.RunServer()
}
