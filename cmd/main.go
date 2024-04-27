package main

import (
	"net/http"
	"saarm/pkg/server"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${remote_ip} | ${method} ${uri} - ${status} - ${latency_human}\n",
	}))

	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {}))

	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "",
		ContentTypeNosniff:    "",
		XFrameOptions:         "",
		HSTSMaxAge:            3600,
		ContentSecurityPolicy: "default-src 'self'",
	}))

	e.Use(middleware.Gzip())

	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowCredentials: false,
	}))

	server.Bootstrap(e)

	e.GET("/explorer/*", echoSwagger.WrapHandler)
	e.GET("/ping", func(c echo.Context) error { return c.JSON(200, echo.Map{"success": true}) })

	e.Logger.Fatal(e.Start(":8000"))
}
