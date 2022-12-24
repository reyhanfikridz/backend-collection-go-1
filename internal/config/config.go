/*
Package config contain all configuration
*/
package config

import (
	"os"

	"github.com/joho/godotenv"
)

// MySQLConfig struct for database (MySQL) configuration
type MySQLConfig struct {
	DBUser     string
	DBPass     string
	DBName     string
	DBTestName string
	DBHost     string
	DBPort     string
}

// GinConfig struct for router (Gin) configuration
type GinConfig struct {
	Host string
	Port string
}

// Config struct for both database and router configuration
type Config struct {
	MySQLConfig
	GinConfig
}

// GetConfig get configuration
func GetConfig(env string) (Config, error) {
	// Load env data
	err := godotenv.Load(os.ExpandEnv(
		"$GOPATH/src/github.com/reyhanfikridz/backend-collection-go-1/.env"))
	if err != nil {
		return Config{}, err
	}

	// get db name
	DBName := os.Getenv("MYSQL_DBNAME")
	if env == "test" {
		DBName = os.Getenv("MYSQL_DBTESTNAME")
	}

	// get config
	conf := Config{
		MySQLConfig{
			DBUser: os.Getenv("MYSQL_DBUSER"),
			DBPass: os.Getenv("MYSQL_DBPASS"),
			DBName: DBName,
			DBHost: os.Getenv("MYSQL_DBHOST"),
			DBPort: os.Getenv("MYSQL_DBPORT"),
		},
		GinConfig{
			Host: os.Getenv("GIN_HOST"),
			Port: os.Getenv("GIN_PORT"),
		},
	}

	return conf, nil
}
