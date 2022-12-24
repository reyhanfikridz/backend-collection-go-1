/*
Package app contain main application of backend-collection-go-1
*/
package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/reyhanfikridz/backend-collection-go-1/internal/config"
	"github.com/reyhanfikridz/backend-collection-go-1/internal/student"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// App struct of application
type App struct {
	DB     *gorm.DB
	Router *gin.Engine
}

// Run shortcut for running the application
func (app *App) Run(conf config.Config) {
	app.InitializeDB(conf.MySQLConfig)
	app.InitializeRouter()
	app.Router.Run(fmt.Sprintf("%s:%s", conf.GinConfig.Host, conf.GinConfig.Port))
}

// InitializeDB initialize database application
func (app *App) InitializeDB(conf config.MySQLConfig) error {
	var err error

	// set DSN
	DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DBUser, conf.DBPass, conf.DBHost, conf.DBPort, conf.DBName)

	// open the database with silent logger
	app.DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return fmt.Errorf("cannot connect to database! Error: %s", err)
	}

	return nil
}

// InitializeRouter initialize router application
func (app *App) InitializeRouter() {
	// init gin router
	app.Router = gin.Default()

	// set student route handler
	studentHandler := student.Handler{DB: app.DB}
	app.Router.POST("api/students", studentHandler.AddStudent)
	app.Router.GET("api/students/:id", studentHandler.GetStudent)
	app.Router.GET("api/students", studentHandler.GetStudents)
	app.Router.PUT("api/students/:id", studentHandler.ReplaceStudent)
	app.Router.PATCH("api/students/:id", studentHandler.UpdateStudent)
	app.Router.DELETE("api/students/:id", studentHandler.DeleteStudent)
}
