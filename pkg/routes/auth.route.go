package routes

import (
	"saarm/pkg/common"
	"saarm/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func AuthGroupRoutes(g *echo.Group) {
	authGroup := g.Group(common.AUTH_PATH)

	authGroup.POST("/sign-in", controllers.SignIn)
	authGroup.POST("/sign-up", controllers.SignUp)
}
