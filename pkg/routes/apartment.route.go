package routes

import (
	"saarm/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func ApartmentRoutes(g *echo.Group) {
	apartmentGroup := g.Group("/apartments")

	apartmentGroup.GET("", controllers.GetApartments)
	apartmentGroup.GET("/:id", controllers.GetApartmentByID)
	apartmentGroup.POST("", controllers.CreateApartments)
	apartmentGroup.PATCH("/:id", controllers.PatchApartmentByID)
	apartmentGroup.PUT("/:id", controllers.PutApartmentByID)
	apartmentGroup.DELETE("", controllers.DeleteApartmentByID)
	apartmentGroup.DELETE("/:id", controllers.DeleteApartmentByID)

	apartmentGroup.POST("/:id/link/users/:userId", controllers.LinkUserApartment)
}
