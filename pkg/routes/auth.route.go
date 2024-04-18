package routes

import (
	"saarm/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func AuthGroupRoutes(g *echo.Group) {
	authGroup := g.Group("/auth")

	authGroup.POST("/sign-in", controllers.SignIn)
	authGroup.POST("/sign-up", controllers.SignUp)
}
