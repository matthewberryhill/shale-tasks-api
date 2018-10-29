package server

import (
	"net/http"

	"github.com/matthewberryhill/shale-tasks-api/models"

	"github.com/labstack/echo"
)

func CreateTask(c echo.Context) error {
	type taskPayload struct {
		Task string `json:"task"`
	}

	tp := new(taskPayload)
	if err := c.Bind(tp); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	t := models.NewTask(tp.Task)
	if err := t.CreateTask(); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, t)
}

func GetTasks(c echo.Context) error {
	ts, err := models.GetTasks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, ts)
}

func GetTaskById(c echo.Context) error {
	id := c.Param("id")

	t, err := models.GetTaskById(id)
	if err != nil {
		if err.Error() == "not found" {
			return c.JSON(http.StatusNotFound, err)
		}
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, t)
}

func DeleteTask(c echo.Context) error {
	id := c.Param("id")

	err := models.DeleteTask(id)
	if err != nil {
		if err.Error() == "not found" {
			return c.JSON(http.StatusNotFound, err)
		}
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusNoContent)
}
