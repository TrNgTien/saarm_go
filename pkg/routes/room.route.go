package routes

import (
	"saarm/pkg/common"
	"saarm/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func RoomGroupRoutes(g *echo.Group) {
	rGroup := g.Group(common.ROOM_PATH)

  // roomGroup.POST(":id", controllers.GetWaterMeter)
  // roomGroup.GET(":id/bills", controllers.GetWaterMeter)
  rGroup.POST(":id/water-meter", controllers.GetWaterMeter)
}
