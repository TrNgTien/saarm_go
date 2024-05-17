package routes

import (
	"saarm/pkg/common"
	"saarm/pkg/controllers"
	"saarm/pkg/middlewares"

	"github.com/labstack/echo/v4"
)

func RoomGroupRoutes(g *echo.Group) {
	rGroup := g.Group(common.ROOM_PATH)

	// roomGroup.POST(":id", controllers.GetWaterMeter)
	rGroup.GET("/:id/bills", controllers.GetBills)
	rGroup.POST("/:id/water-meter", controllers.GetWaterMeter)

	rGroup.POST("", controllers.CreateRoom, middlewares.LandlordPermission)
	rGroup.POST("/:id/duplicate", controllers.DuplicateRoom, middlewares.LandlordPermission)

	rGroup.PATCH("/:id", controllers.PatchRoomByID, middlewares.LandlordPermission)
	rGroup.PUT("/:id", controllers.PutRoomByID, middlewares.LandlordPermission)
	rGroup.DELETE("", controllers.DeleteRooms, middlewares.LandlordPermission)
	rGroup.DELETE("/:id", controllers.DeleteRoomByID, middlewares.LandlordPermission)
}
