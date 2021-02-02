package main

import (
	"github.com/cerc-undip/tandera/config"
	"github.com/cerc-undip/tandera/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	e.GET("/users", controllers.GetUsers)
	e.POST("/users", controllers.CreateUser)

	e.Logger.Fatal(e.Start(config.PORT))
}
