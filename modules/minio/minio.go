package minio

import min "github.com/minio/minio-go"

var (
	client *min.Client
)

func Init() {
	var (
		err error
	)

	client, err = min.New("localhost:9000", "sems_be", "sems_be", false)

	if err != nil {
		panic(err)
	}
}

func GetClient() *min.Client {
	return client
}
