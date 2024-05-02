package server

import (
	"saarm/modules/minio"
	"saarm/modules/pg"
	"saarm/pkg/configs"
	"saarm/pkg/helpers"
	"saarm/pkg/routes"

	"github.com/labstack/echo/v4"
)

func Bootstrap(e *echo.Echo) {
  configs.Init()
	pg.InitPg()
	minio.Init()
	routes.Init(e)
  helpers.InitCron()
}
