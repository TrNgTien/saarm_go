package minio

import (
	"os"
	"saarm/pkg/utilities"

	min "github.com/minio/minio-go"
)

var (
	client *min.Client
)

func Init() {
	var (
		err error
	)

	client, err = min.New(utilities.GetValueEnv(os.Getenv("APP_ENV_MINIO_ENDPOINT"), "localhost:9000"),
		utilities.GetValueEnv(os.Getenv("APP_ENV_MINIO_ACCESS_KEY"), "minio-root"),
		utilities.GetValueEnv(os.Getenv("APP_ENV_MINIO_SECRET_KEY"), "tien19217"),
		false)

	if err != nil {
		panic(err)
	}
}

func GetClient() *min.Client {
	return client
}
