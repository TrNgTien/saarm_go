package configs

import (
	"os"
	"saarm/pkg/utilities"
)

type Config struct {
	DB struct {
		PG struct {
			Host     string
			Port     int
			Username string
			Password string
			Database string
		}
	}
}

var env Config

func Init() {
	env.DB.PG.Host = os.Getenv("APP_ENV_POSTGRES_HOST")
	env.DB.PG.Port = utilities.GetIntValue(os.Getenv("APP_ENV_POSTGRES_PORT"))
	env.DB.PG.Username = os.Getenv("APP_ENV_POSTGRES_USERNAME")
	env.DB.PG.Password = os.Getenv("APP_ENV_POSTGRES_PASSWORD")
	env.DB.PG.Database = os.Getenv("APP_ENV_POSTGRES_DATABASE")
}

func GetENV() Config {
	return env
}
