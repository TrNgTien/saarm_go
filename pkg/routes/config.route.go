package routes

import (
	"saarm/pkg/common"
	"saarm/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func ConfigGroupRoutes(g *echo.Group) {
	configGroup := g.Group(common.CONFIG_PATH)
	minioBucketPath := "minio/buckets"

	configGroup.GET(minioBucketPath, controllers.GetBuckets)
	configGroup.POST(minioBucketPath, controllers.CreateBucket)
	configGroup.DELETE(minioBucketPath, controllers.DeleteBucket)
	configGroup.POST(minioBucketPath+"/:name", controllers.UploadObject)
}
