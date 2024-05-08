package routes

import (
	"saarm/pkg/common"
	"saarm/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func ApartmentAccountRoutes(g *echo.Group) {
	accountGroup := g.Group(common.APARTMENT_ACCOUNT_PATH)

	accountGroup.GET("", controllers.GetApartments)
	accountGroup.GET(":id", controllers.GetApartmentByID)
	accountGroup.POST("", controllers.CreateApartments)
	accountGroup.PATCH(":id", controllers.PatchApartmentByID)
	accountGroup.PUT(":id", controllers.PutApartmentByID)
	accountGroup.DELETE("", controllers.DeleteApartmentByID)
	accountGroup.DELETE(":id", controllers.DeleteApartmentByID)

	accountGroup.POST(":id/link/accounts/:accountId", controllers.LinkUserApartment)
}
