package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/matthewberryhill/shale-tasks-api/models"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

type responseError struct {
	Error string `json:"error"`
}

func CreateTask(c echo.Context) error {
	type taskPayload struct {
		Task string `json:"task"`
	}

	tp := new(taskPayload)
	if err := c.Bind(tp); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if tp.Task == "" {
		er := new(responseError)
		er.Error = "Request payload requires the 'task' field"
		return c.JSON(http.StatusBadRequest, er)
	}

	ts, err := models.GetTasks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	for _, t := range ts {
		if tp.Task == t.Task {
			re := new(responseError)
			re.Error = "Task cannot share the same string as the 'task' field of another task"
			return c.JSON(http.StatusConflict, re)
		}
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

func UpdateTask(c echo.Context) error {
	type taskPayload struct {
		Task      string `json:"task,omitempty"`
		Completed bool   `json:"completed,omitempty"`
	}
	tp := new(taskPayload)
	if err := c.Bind(tp); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	var tpi map[string]interface{}
	tpiJson, _ := json.Marshal(tp)
	json.Unmarshal(tpiJson, &tpi)

	ts, err := models.GetTasks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	id := bson.ObjectIdHex(c.Param("id"))
	t, err := models.GetTaskById(id.Hex())
	if err != nil {
		if err.Error() == "not found" {
			return c.JSON(http.StatusNotFound, err)
		}
		return c.JSON(http.StatusInternalServerError, err)
	}

	for _, tt := range ts {
		if tp.Task == tt.Task && t.Task != tp.Task {
			re := new(responseError)
			re.Error = "Task cannot share the same string as the 'task' field of another task"
			return c.JSON(http.StatusConflict, re)
		}
	}

	for k, v := range tpi {
		if k == "task" {
			t.Task = v.(string)
		} else if k == "completed" {
			if v.(bool) {
				now := time.Now()
				t.DateCompleted = &now
				t.Completed = true
			}
		}
	}

	err = t.UpdateTask()
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
