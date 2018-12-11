package mongo

import (
	"sync"	

	"github.com/astaxie/beego"

	log "github.com/sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
)

type Database struct {
	Name 	string
	Session *mgo.Session
}

var once			sync.Once 
var MongoSession	*Database

func init() {
	MongoSession = GetSession()
	log.Infoln("Current database is ", MongoSession.Name)
}

func GetSession() *Database {
	once.Do(func() {
		MongoSession = create(
							beego.AppConfig.String("mongodb.url"),
							beego.AppConfig.String("mongodb.name"),
						)
	})

	return MongoSession
}

func create(mongodb_host string, mongodb_name string) *Database {
	session, err := mgo.Dial(mongodb_host)

	if err != nil {
		log.Errorln(err)
		log.Fatalln("Mongodb connection failed.")
	}

	log.Infoln("Mongodb connection successed.")

	return &Database{
		Name: mongodb_name, 
		Session: session}
}