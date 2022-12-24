package main

import (
	"log"

	"github.com/reyhanfikridz/backend-collection-go-1/app"
	"github.com/reyhanfikridz/backend-collection-go-1/internal/config"
)

// main function
func main() {
	// get app config for development
	conf, err := config.GetConfig("development")
	if err != nil {
		log.Fatal(err)
	}

	// get and run app
	app := app.App{}
	app.Run(conf)
}
