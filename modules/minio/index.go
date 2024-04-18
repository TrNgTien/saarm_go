package minio

import (
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	minioClient *minio.Client
)

func Init() {
	var (
		err error
	)

	endpoint := "localhost:9000"
	accessKeyID := os.Getenv("APP_ENV_MINIO_SECRET_KEY")
	secretAccessKey := os.Getenv("APP_ENV_MINIO_ACCESS_KEY")
	useSSL := false

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		panic(err)
	}

	log.Printf("%#v\n", minioClient) // minioClient is now setup
}

func GetClient() *minio.Client {
	return minioClient
}
