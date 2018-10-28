package main

import (
	"github.com/matthewberryhill/shale-tasks-api/server"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/config", server.GetConfig)

	e.Logger.Fatal(e.Start(":1323"))
}
