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

  rGroup.POST("", controllers.CreateRoom, middlewares.HomeownerPermission)
  rGroup.POST("/:id/duplicate", controllers.DuplicateRoom, middlewares.HomeownerPermission)
	rGroup.POST("/:id/water-meters/detect", controllers.DetectWaterMeter)
	rGroup.POST("/:id/water-meters/submit", controllers.ConfirmWaterMeter)

	rGroup.PATCH("/:id", controllers.PatchRoomByID, middlewares.HomeownerPermission)
	rGroup.PUT("/:id", controllers.PutRoomByID, middlewares.HomeownerPermission)
	rGroup.DELETE("", controllers.DeleteRooms, middlewares.HomeownerPermission)
	rGroup.DELETE("/:id", controllers.DeleteRoomByID, middlewares.HomeownerPermission)
}
