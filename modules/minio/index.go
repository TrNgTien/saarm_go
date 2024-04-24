package minio

import (
	min "github.com/minio/minio-go"
)

var (
	client *min.Client
)

func Init() {
	var (
		err error
	)

	client, err = min.New("localhost:9000", "minio-root", "tien19217", false)
	if err != nil {
		panic(err)
	}
}

func GetClient() *min.Client {
	return client
}
