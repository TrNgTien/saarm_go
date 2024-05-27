package routes

import (
	"saarm/pkg/common"
	"saarm/pkg/controllers"
	"saarm/pkg/middlewares"

	"github.com/labstack/echo/v4"
)

func RoomGroupRoutes(g *echo.Group) {
	rGroup := g.Group(common.ROOM_PATH)

	rGroup.GET("/:id", controllers.GetRoomByID)
	rGroup.GET("/bills", controllers.GetBills)
	rGroup.GET("/:id/bills", controllers.GetBillByRoom)
	rGroup.GET("/:id/water-meters/histories", controllers.GetHistorySubmitted)
	rGroup.GET("/:id/water-meters/is-submitted", controllers.CheckSubmittedWaterMeter)
	rGroup.POST("/:id/water-meters/detect", controllers.DetectWaterMeter)
	rGroup.POST("/:id/water-meters/submit", controllers.ConfirmWaterMeter)

	rGroup.POST("", controllers.CreateRoom, middlewares.LandlordPermission)
	rGroup.POST("/:id/duplicate", controllers.DuplicateRoom, middlewares.LandlordPermission)

	rGroup.PATCH("/:id", controllers.PatchRoomByID, middlewares.LandlordPermission)
	rGroup.PUT("/:id", controllers.PutRoomByID, middlewares.LandlordPermission)
	rGroup.DELETE("", controllers.DeleteRooms, middlewares.LandlordPermission)
	rGroup.DELETE("/:id", controllers.DeleteRoomByID, middlewares.LandlordPermission)
}
