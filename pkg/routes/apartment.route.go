package routes

import (
	"saarm/pkg/common"
	"saarm/pkg/controllers"
	"saarm/pkg/middlewares"

	"github.com/labstack/echo/v4"
)

func ApartmentRoutes(g *echo.Group) {
	aGroup := g.Group(common.APARTMENT_PATH, middlewares.LandlordPermission)

	aGroup.GET("", controllers.GetApartments)
	aGroup.GET(":id", controllers.GetApartmentByID)

	aGroup.POST("", controllers.CreateApartments)

	aGroup.PATCH(":id", controllers.PatchApartmentByID)
	aGroup.PUT(":id", controllers.PutApartmentByID)

	aGroup.DELETE("", controllers.DeleteApartmentByID)
	aGroup.DELETE(":id", controllers.DeleteApartmentByID)

	// Remove
	// aGroup.POST(":id/link/users/:userId", controllers.LinkUserApartment)
}
