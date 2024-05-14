package routes

import (
	"fmt"
	"net/http"
	"saarm/pkg/common"
	"saarm/pkg/utilities"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Ping(c echo.Context) error {
	return c.String(http.StatusOK, "Pong!")
}

func Init(e *echo.Echo) {

	g := e.Group("/v1/api")
	g.GET("/ping", Ping)

	AuthGroupRoutes(g)

	p := g.Group("")

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(common.JwtCustomClaims)
		},
		SigningKey: []byte(utilities.GetValueEnv("APP_ENV_SECRET_KEY", "secretKey")),
		SuccessHandler: func(c echo.Context) {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(*common.JwtCustomClaims)

			role, userID := claims.Role, claims.UserID

			if role == "" || userID == "" {
				fmt.Println("Unauthorized!")
				return
			}

      c.Set("role", role)
      c.Set("userID", userID)
		},
		ErrorHandler: func(c echo.Context, _ error) error {
			_, ok := c.Get("user").(*jwt.Token)

			if !ok {
				fmt.Println("JWT token missing or invalid")
				return utilities.R400(c, "JWT token missing or invalid")
			}

			return nil
		},
	}

	p.Use(echojwt.WithConfig(config))

	UserGroupRoutes(p)
	RoomGroupRoutes(p)
	ConfigGroupRoutes(p)
	ApartmentRoutes(p)
	RoleGroupRoutes(p)
}
