package models

import (
	"flag"
	"runtime"
	"testing"

	"github.com/matthewberryhill/shale-tasks-api/config"

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

		Convey("When calling models.CreateTask(...)", func() {
			before := runtime.NumGoroutine()
			t := NewTask("test")
			t.CreateTask()
			So(t.Task, ShouldEqual, "test")

			testId = t.Id
			So(testId, ShouldNotBeNil)

			Convey("Then the task should be retrievable", func() {
				//todo: implement
			})

			Convey("Then number of goroutines after call should be the same", func() {
				after := runtime.NumGoroutine()
				So(before, ShouldBeLessThanOrEqualTo, after)
			})
		})
	})
}
