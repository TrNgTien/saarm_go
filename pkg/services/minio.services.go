package services

import (
	"fmt"
	"io"
	"math/rand"
	"path/filepath"
	"saarm/modules/minio"

	min "github.com/minio/minio-go"
)

func IsExistedBucket(bucketName string) (bool, error) {
	isExisted, err := minio.GetClient().BucketExists(bucketName)

	if err != nil {

		return false, err
	}

	return isExisted, nil
}

func getFileExtension(filename string) string {
	return filepath.Ext(filename)[1:]
}

func generateRandomString(n int) string {
	// Use a more secure random number generator in production
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func GetListBuckets() ([]min.BucketInfo, error) {
	var bucketList []min.BucketInfo

	buckets, err := minio.GetClient().ListBuckets()

	if err != nil {
		fmt.Println(err)
		return bucketList, err
	}

	bucketList = append(bucketList, buckets...)

	return bucketList, nil
}

func CreateBucket(bucketName string) (err error) {
	location := ""

	if err := minio.GetClient().MakeBucket(bucketName, location); err != nil {
		fmt.Println("[CreateBucket]: ", err)
	}

	return nil
}

func DeleteBucket(bucketName string) (err error) {
	if err := minio.GetClient().RemoveBucket(bucketName); err != nil {
		fmt.Println("[DeleteBucket]: ", err)
		return err
	}

	return nil
}

func UploadObject(bucketName string, objectName string, object io.Reader, size int64) error {

	// Get file extension for content type
	// ext := getFileExtension(file.Filename)
	// contentType := "application/multipart"
	// Create a unique object name to avoid conflicts
	// objectName := fmt.Sprintf("images/%s.%s", generateRandomString(10), ext)

	// _, err := minio.GetClient().PutObject(bucketName, objectName, object, size, minio.PutObjectOptions{ContentType: contentType})

	// if err != nil {
	// return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error uploading image to Minio: %v", err))
	// }

	return nil
}
