package main

import (
	"flag"
	"os"

	"github.com/matthewberryhill/shale-tasks-api/config"
	"github.com/matthewberryhill/shale-tasks-api/models"
	"github.com/matthewberryhill/shale-tasks-api/server"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tkanos/gonfig"
)

var env = flag.String(
	"env",
	"dev",
	"config env",
)

func main() {
	flag.Parse()
	var configPath string

	env := os.Getenv("MONGO")
	if env == "dev" {
		configPath = "./config/dev.json"
	} else {
		configPath = "./config/prod.json"
	}
	conf := config.Config{}
	err := gonfig.GetConf(configPath, &conf)
	if err != nil {
		panic(err)
	}

	models.ConfigureDB(conf.Mongo)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/config", server.GetConfig)

	e.POST("/tasks", server.CreateTask)
	e.GET("/tasks", server.GetTasks)
	e.GET("/tasks/:id", server.GetTaskById)
	e.PUT("/tasks/:id", server.UpdateTask)
	e.DELETE("/tasks/:id", server.DeleteTask)

	e.Logger.Fatal(e.Start(":1323"))
}
