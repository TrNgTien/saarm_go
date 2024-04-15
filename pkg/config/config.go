package config

import (
	"os"
)

type Config struct {
	DB struct {
		PG struct {
			Host     string
			Port     string
			Username string
			Password string
			Database string
		}
	}
}

var env Config

func Init() {
	env.DB.PG.Host = os.Getenv("APP_ENV_POSTGRES_HOST")
	env.DB.PG.Port = os.Getenv("APP_ENV_POSTGRES_PORT")
	env.DB.PG.Username = os.Getenv("APP_ENV_POSTGRES_USERNAME")
	env.DB.PG.Password = os.Getenv("APP_ENV_POSTGRES_PASSWORD")
	env.DB.PG.Database = os.Getenv("APP_ENV_POSTGRES_DATABASE")
}

func GetENV(key string) string {
	return os.Getenv(key)
}
