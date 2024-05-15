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
	rGroup.POST("/:id/water-meter", controllers.GetWaterMeter)

	rGroup.GET("", controllers.GetApartments)
	rGroup.GET("/:id", controllers.GetApartmentByID)

	rGroup.POST("", controllers.CreateRoom)
	rGroup.POST("/:id/duplicate", controllers.DuplicateRoom)

	rGroup.PATCH("/:id", controllers.PatchApartmentByID)
	rGroup.PUT("/:id", controllers.PutApartmentByID)
	rGroup.DELETE("", controllers.DeleteApartmentByID)
	rGroup.DELETE("/:id", controllers.DeleteApartmentByID)
}
