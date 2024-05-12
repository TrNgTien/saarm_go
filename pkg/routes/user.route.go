package routes

import (
	"saarm/pkg/common"
	"saarm/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func UserGroupRoutes(g *echo.Group) {
	uGroup := g.Group(common.USER_PATH)

	uGroup.GET("/", controllers.GetUsers)
	uGroup.GET("/:id", controllers.GetUserByID)
	uGroup.POST("/", controllers.CreateUser)
	uGroup.PUT("/:id", controllers.UpdateUser)
	uGroup.PATCH("/:id", controllers.UpdateUser)
}
