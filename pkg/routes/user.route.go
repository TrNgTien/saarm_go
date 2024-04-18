package routes

import (
	"saarm/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func UserGroupRoutes(g *echo.Group) {
	g.GET("/users", controllers.GetUsers)
	g.GET("/users/:id", controllers.GetUserByID)
}
