package server

import (
	"fmt"
	"saarm/modules/minio"
	"saarm/modules/pg"
	"saarm/pkg/configs"
	"saarm/pkg/helpers"
	"saarm/pkg/routes"

	"github.com/labstack/echo/v4"
)

func Bootstrap(e *echo.Echo) {
  fmt.Println("[Bootstrap] START Configs things................")
  configs.Init()
  helpers.InitCron()

  fmt.Println("[Bootstrap] START Postgresql................")
	pg.InitPg()


  fmt.Println("[Bootstrap] START Minio................")
	minio.Init()

  fmt.Println("[Bootstrap] START Routes................")
	routes.Init(e)
}
