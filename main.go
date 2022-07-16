package main

import (
	"github.com/jeanmolossi/effective-eureka/src/cmd"

	_ "github.com/jeanmolossi/effective-eureka/docs"
	_ "github.com/swaggo/files"
)

func main() {
	cmd.RunServer()
}
