package routes

import (
	"saarm/pkg/common"
	"saarm/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func RoleGroupRoutes(g *echo.Group) {
	rGroup := g.Group(common.ROLE_PATH)

  rGroup.POST("/", controllers.CreateRole)

	rGroup.GET("", controllers.GetApartments)
	rGroup.GET(":id", controllers.GetApartmentByID)

	rGroup.PATCH(":id", controllers.PatchApartmentByID)
	rGroup.PUT(":id", controllers.PutApartmentByID)

	rGroup.DELETE("", controllers.DeleteApartmentByID)
	rGroup.DELETE(":id", controllers.DeleteApartmentByID)
}
