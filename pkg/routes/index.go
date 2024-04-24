package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Ping(c echo.Context) error {
	return c.String(http.StatusOK, "Pong!")
}

func Init(e *echo.Echo) {

	g := e.Group("/v1/api")
	g.GET("/ping", Ping)

	AuthGroupRoutes(g)

	UserGroupRoutes(g)

	RoomGroupRoutes(g)

	ConfigGroupRoutes(g)
}
