package main

import (
	"log"

	"github.com/reyhanfikridz/backend-collection-go-1/app"
	"github.com/reyhanfikridz/backend-collection-go-1/internal/config"
	"github.com/reyhanfikridz/backend-collection-go-1/internal/student"
)

// main function
func main() {
	// migrate development database
	err := MigrateDB("development")
	if err != nil {
		log.Print("Migrate database development failed! Error:", err)
	} else {
		log.Print("Migrate database development success!")
	}

	// migrate test database
	err = MigrateDB("test")
	if err != nil {
		log.Print("Migrate database test failed! Error:", err)
	} else {
		log.Print("Migrate database test success!")
	}
}

// MigrateDB migrate database
//
// env == "development" for development database
//
// env == "test" for test database
func MigrateDB(env string) error {
	// get config
	conf, err := config.GetConfig(env)
	if err != nil {
		return err
	}

	// get app and initialize database only
	app := app.App{}
	app.InitializeDB(conf.MySQLConfig)

	// migrate database
	err = app.DB.AutoMigrate(student.Student{})
	return err
}
