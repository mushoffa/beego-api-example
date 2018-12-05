package database

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	log "github.com/sirupsen/logrus"
)

var MgoSession *mgo.Session

func init() {
	mongodb_host := beego.AppConfig.String("mongodb.url")
	session, err := mgo.Dial(mongodb_host)

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Info("Successfully connected to mongodb at ", mongodb_host)

	MgoSession =  session
}