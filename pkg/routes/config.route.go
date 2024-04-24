package routes

import (
	"saarm/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func ConfigGroupRoutes(g *echo.Group) {
	configGroup := g.Group("/config")

	configGroup.GET("/minio/buckets", controllers.GetBuckets)

	configGroup.POST("/minio/buckets", controllers.CreateBucket)

	configGroup.DELETE("/minio/buckets", controllers.DeleteBucket)

  configGroup.POST("/minio/buckets/:name", controllers.UploadObject)
}
