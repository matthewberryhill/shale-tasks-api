package server

import (
	"net/http"

	"github.com/matthewberryhill/shale-tasks-api/models"

	"github.com/labstack/echo"
)

func GetConfig(c echo.Context) error {
	conf := models.GetConfig()

	return c.JSON(http.StatusOK, conf)
}
