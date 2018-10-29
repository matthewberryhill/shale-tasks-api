package models

import (
	"log"

	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var mongoAddr *mgo.DialInfo

func ConfigureDB(config []string) {
	mongoAddr = &mgo.DialInfo{
		Addrs: config,
	}
}

type Task struct {
	Id          *bson.ObjectId `json:"id" bson:"_id"`
	Task        string         `json:"task"`
	DateCreated *time.Time     `json:"date_created"`
	Completed   bool           `json:"completed"`
}

func NewTask(task string) *Task {
	t := new(Task)
	id := bson.NewObjectId()
	t.Id = &id
	t.Task = task
	now := time.Now()
	t.DateCreated = &now
	t.Completed = false

	return t
}

func (u *Task) CreateTask() error {
	session, err := mgo.DialWithInfo(mongoAddr)
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
		return err
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("shale").C("tasks")
	_, err = c.UpsertId(u.Id, u)
	if err != nil {
		log.Println("Error creating Task: ", err.Error())
		return err
	}

	return nil
}

func GetTasks() ([]*Task, error) {
	session, err := mgo.DialWithInfo(mongoAddr)
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
		return nil, err
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("shale").C("tasks")
	var Tasks []*Task
	err = c.Find(bson.M{}).All(&Tasks)
	if err != nil {
		log.Println("Error retrieving Tasks: ", err.Error())
		return nil, err
	}

	return Tasks, nil
}

func DeleteTask(id string) error {
	session, err := mgo.DialWithInfo(mongoAddr)
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
		return nil
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("shale").C("tasks")
	err = c.RemoveId(bson.ObjectIdHex(id))
	if err != nil {
		log.Println("Error deleteing Property: ", err.Error())
		return err
	}

	return nil
}
