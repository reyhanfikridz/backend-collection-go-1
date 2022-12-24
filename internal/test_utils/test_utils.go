package test_utils

import (
	"encoding/json"
	"io"

	"github.com/reyhanfikridz/backend-collection-go-1/app"
	"github.com/reyhanfikridz/backend-collection-go-1/internal/config"
	"github.com/steinfletcher/apitest"
)

// GetTestApp get testing application
func GetTestApp() (app.App, error) {
	conf, err := config.GetConfig("test")
	if err != nil {
		return app.App{}, err
	}

	a := app.App{}
	a.InitializeDB(conf.MySQLConfig)
	if err != nil {
		return app.App{}, err
	}

	a.InitializeRouter()
	return a, nil
}

// ParseDataFromApitestResult parse data from apitest call api result
func ParseDataFromApitestResult(result apitest.Result, data any) error {
	body, err := io.ReadAll(result.Response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	return nil
}
