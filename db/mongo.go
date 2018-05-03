package db

import (
	"gitlab.smartm2m.co.kr/dev/stove/config"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	UserCollection     string = "user"
	ResumeCollection   string = "resume"
	QuestionCollection string = "question"
)

var session *mgo.Session
var database *mgo.Database

func Connect(conf config.MongoDB) error {
	var err error
	session, err = mgo.Dial(conf.Host)
	if err != nil {
		return err
	}

	database = session.DB(conf.DBName)

	return nil
}

func Insert(col string, doc interface{}) error {
	c := database.C(col)
	return c.Insert(doc)
}

func Update(col string, selector interface{}, doc interface{}) error {
	c := database.C(col)
	return c.Update(selector, doc)
}

func FindOne(col string, selector bson.M, result interface{}) error {
	c := database.C(col)
	return c.Find(selector).One(result)
}

func Find(col string, selector bson.M, result interface{}) error {
	c := database.C(col)
	return c.Find(selector).All(result)
}

func FindAll(col string, results interface{}) error {
	c := database.C(col)
	q := c.Find(nil)

	return q.All(results)
}
