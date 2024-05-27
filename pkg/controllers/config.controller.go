package controllers

import (
	"fmt"
	"net/http"

	// "saarm/pkg/common"
	"saarm/pkg/services"
	"saarm/pkg/utilities"

	// "github.com/golang-jwt/jwt/v5"

	"github.com/labstack/echo/v4"
)

type BucketRequest struct {
	BucketName string `json:"bucketName" validate:"required"`
}

type ObjectRequest struct {
	ObjectName string `json:"objectName" validate:"required"`
}

type UploadObjectResponse struct {
	FileName string `json:"fileName" validate:"required"`
	FileSize int64  `json:"fileSize" validate:"required"`
}

func GetBuckets(c echo.Context) error {
	buckets, err := services.GetListBuckets()

	if err != nil {
		return utilities.R400(c, "[GetBuckets]: GetBuckets Error")
	}

	return utilities.R200(c, buckets)
}

func CreateBucket(c echo.Context) error {
	b := new(BucketRequest)

	if err := c.Bind(b); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	bucket := BucketRequest{
		BucketName: b.BucketName,
	}

	if err := services.CreateBucket(bucket.BucketName); err != nil {
		return utilities.R400(c, "[CreateBucket] : create bucket failed!")
	}

	return utilities.R204(c)
}

func DeleteBucket(c echo.Context) error {
	b := new(BucketRequest)

	if err := c.Bind(b); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	bucket := BucketRequest{
		BucketName: b.BucketName,
	}

	if err := services.DeleteBucket(bucket.BucketName); err != nil {
		return utilities.R400(c, "[DeleteBucket] : delete bucket failed!")
	}

	return utilities.R204(c)
}

func UploadObject(c echo.Context) error {
	name := c.Param("name")
	formValues, err := c.FormFile("image")

	fmt.Println("[name]", name)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	fileName := formValues.Filename
	fileSize := formValues.Size

	return utilities.R200(c, UploadObjectResponse{FileName: fileName, FileSize: fileSize})
}
