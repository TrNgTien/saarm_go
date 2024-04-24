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

func Init() error {
	port, err := utilities.GetIntValue(os.Getenv("APP_ENV_POSTGRES_PORT"))

	if err != nil {
		return err
	}

	env.DB.PG.Host = os.Getenv("APP_ENV_POSTGRES_HOST")
	env.DB.PG.Port = port
	env.DB.PG.Username = os.Getenv("APP_ENV_POSTGRES_USERNAME")
	env.DB.PG.Password = os.Getenv("APP_ENV_POSTGRES_PASSWORD")
	env.DB.PG.Database = os.Getenv("APP_ENV_POSTGRES_DATABASE")
	return nil
}

func GetENV() Config {
	return env
}
