package routes

import (
	"saarm/pkg/common"
	"saarm/pkg/controllers"
	"saarm/pkg/middlewares"

	"github.com/labstack/echo/v4"
)

func ApartmentRoutes(g *echo.Group) {
	aGroup := g.Group(common.APARTMENT_PATH, middlewares.HomeownerPermission)

	aGroup.GET("", controllers.GetApartments)
  aGroup.GET("/users/:id", controllers.GetApartmentsByUserID)
	aGroup.GET("/:id", controllers.GetApartmentByID)

	aGroup.POST("", controllers.CreateApartments)

	aGroup.PATCH("/:id", controllers.PatchApartmentByID)
	aGroup.PUT("/:id", controllers.PutApartmentByID)

	aGroup.DELETE("", controllers.DeleteApartmentByID)
	aGroup.DELETE("/:id", controllers.DeleteApartmentByID)
}
