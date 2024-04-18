package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Ping(c echo.Context) error {
	return c.String(http.StatusOK, "Pong!")
}

func Init(e *echo.Echo) {

	g := e.Group("/v1/api")
	g.GET("", Ping)

	AuthGroupRoutes(g)

	UserGroupRoutes(g)
}
