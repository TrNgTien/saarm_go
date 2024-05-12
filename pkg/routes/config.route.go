package routes

import (
	"saarm/pkg/common"
	"saarm/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func ConfigGroupRoutes(g *echo.Group) {
	configGroup := g.Group(common.CONFIG_PATH)

	minioBucketGroup := configGroup.Group("/minio/buckets")
	minioBucketGroup.GET("", controllers.GetBuckets)
	minioBucketGroup.POST("/", controllers.CreateBucket)
	minioBucketGroup.DELETE("/", controllers.DeleteBucket)
	minioBucketGroup.POST("/:name", controllers.UploadObject)

}
