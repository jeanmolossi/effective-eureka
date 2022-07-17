// Package main is the entry point of the application.
//
// Handles the rest api and the event handlers.
package main

import (
	"github.com/jeanmolossi/effective-eureka/src/cmd"

	_ "github.com/jeanmolossi/effective-eureka/docs"
	_ "github.com/swaggo/files"
)

func main() {
	cmd.RunServer()
}
