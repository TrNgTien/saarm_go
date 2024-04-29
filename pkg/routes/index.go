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

	p := g.Group("")

	// p.Use(echojwt.WithConfig(echojwt.Config{
	// 	SigningKey: []byte(os.Getenv("APP_ENV_SECRET_KEY")),
	// 	BeforeFunc: func(c echo.Context) {
	//
	// 		token, ok := c.Get("user").(*jwt.Token) // by default token is stored under `user` key
	// 		fmt.Println("token", token, ok)
	//
	// 		if !ok {
	// 			fmt.Println("err")
	// 			return
	// 			// return utilities.R400(c, "JWT token missing or invalid")
	// 		}
	//
	// 		claims, ok := token.Claims.(jwt.MapClaims) // by default claims is of type `jwt.MapClaims`
	//
	// 		fmt.Println("[claims] ", claims)
	//
	// 		if !ok {
	// 			fmt.Println("failed to cast claims as jwt.MapClaims")
	// 			return
	// 		}
	// 	},
	// }))

	UserGroupRoutes(p)
	RoomGroupRoutes(p)
	ConfigGroupRoutes(p)
	ApartmentRoutes(p)
}
