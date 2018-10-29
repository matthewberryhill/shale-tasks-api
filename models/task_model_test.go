package models

import (
	"flag"
	"runtime"
	"testing"

	"github.com/matthewberryhill/shale-tasks-api/config"

	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tkanos/gonfig"
	"gopkg.in/mgo.v2/bson"
)

var testMongoAddr = flag.String(
	"mongoAddr",
	"localhost:27017",
	"database service address",
)

var testId *bson.ObjectId

func TestTasks_CreateTask(t *testing.T) {
	Convey("If a test database exists", t, func() {
		configPath := "../config/dev.json"
		conf := config.Config{}
		err := gonfig.GetConf(configPath, &conf)
		if err != nil {
			panic(err)
		}

		ConfigureDB(conf.Mongo)

		Convey("When calling task.CreateTask(task)", func() {
			before := runtime.NumGoroutine()
			t := NewTask("test")
			So(t.Task, ShouldEqual, "test")
			t.CreateTask()

			testId = t.Id
			So(testId, ShouldNotBeNil)

			Convey("Then the task should be retrievable", func() {
				retrievedTask, _ := GetTaskById(t.Id.Hex())
				So(retrievedTask.Id.Hex(), ShouldEqual, t.Id.Hex())
			})

			Convey("Then number of goroutines after call should be the same", func() {
				after := runtime.NumGoroutine()
				So(before, ShouldBeLessThanOrEqualTo, after)
			})
		})
	})
}

func TestTasks_GetTasks(t *testing.T) {
	Convey("If a test database exists", t, func() {
		configPath := "../config/dev.json"
		conf := config.Config{}
		err := gonfig.GetConf(configPath, &conf)
		if err != nil {
			panic(err)
		}

		ConfigureDB(conf.Mongo)

		Convey("When calling GetTasks(...)", func() {
			before := runtime.NumGoroutine()
			tasks, _ := GetTasks()

			Convey("Then the retieved tasks [] should be greater than 0", func() {
				So(len(tasks), ShouldBeGreaterThan, 0)
			})

			Convey("Then number of goroutines after call should be the same", func() {
				after := runtime.NumGoroutine()
				So(before, ShouldBeLessThanOrEqualTo, after)
			})
		})
	})
}

func TestTasks_GetTaskById(t *testing.T) {
	Convey("If a test database exists", t, func() {
		configPath := "../config/dev.json"
		conf := config.Config{}
		err := gonfig.GetConf(configPath, &conf)
		if err != nil {
			panic(err)
		}

		ConfigureDB(conf.Mongo)

		Convey("When calling GetTaskById(testId)", func() {
			before := runtime.NumGoroutine()
			task, _ := GetTaskById(testId.Hex())

			Convey("Then the retieved task should have the same id has the testId", func() {
				So(task.Id.Hex(), ShouldEqual, testId.Hex())
			})

			Convey("Then number of goroutines after call should be the same", func() {
				after := runtime.NumGoroutine()
				So(before, ShouldBeLessThanOrEqualTo, after)
			})
		})
	})
}

func TestTasks_UpdateTask(t *testing.T) {
	Convey("If a test database exists", t, func() {
		configPath := "../config/dev.json"
		conf := config.Config{}
		err := gonfig.GetConf(configPath, &conf)
		if err != nil {
			panic(err)
		}

		ConfigureDB(conf.Mongo)

		Convey("When calling model.UpdateTask to be completed", func() {
			before := runtime.NumGoroutine()
			t, _ := GetTaskById(testId.Hex())
			t.Completed = true
			fmt.Println(t.Completed)
			t.UpdateTask()

			Convey("Then the task should be completed", func() {
				updatedTask, _ := GetTaskById(testId.Hex())
				So(updatedTask.Completed, ShouldEqual, true)
			})

			Convey("Then number of goroutines after call should be the same", func() {
				after := runtime.NumGoroutine()
				So(before, ShouldBeLessThanOrEqualTo, after)
			})
		})
	})
}

func TestTasks_DeleteTask(t *testing.T) {
	Convey("If a test database exists", t, func() {
		configPath := "../config/dev.json"
		conf := config.Config{}
		err := gonfig.GetConf(configPath, &conf)
		if err != nil {
			panic(err)
		}

		ConfigureDB(conf.Mongo)

		Convey("When calling DeleteTask(id)", func() {
			before := runtime.NumGoroutine()
			DeleteTask(testId.Hex())

			Convey("Then the task should no longer exist", func() {
				err := DeleteTask(testId.Hex())
				So(err.Error(), ShouldEqual, "not found")
			})

			Convey("Then number of goroutines after call should be the same", func() {
				after := runtime.NumGoroutine()
				So(before, ShouldBeLessThanOrEqualTo, after)
			})
		})
	})
}
