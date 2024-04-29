package routes

import (
	"saarm/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func RoomGroupRoutes(g *echo.Group) {
	roomGroup := g.Group("/rooms/")

  // roomGroup.POST("/:id", controllers.GetWaterMeter)
  // roomGroup.GET("/:id/bills", controllers.GetWaterMeter)
  roomGroup.POST("water-meter", controllers.GetWaterMeter)
}
