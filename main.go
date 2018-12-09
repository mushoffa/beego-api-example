package main

import (
	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
	
	_ "beego-api-example/routers"

	mongodb "beego-api-example/pkg/database/mongo"
)

func init() {
	initMongoDb()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

func initMongoDb() {
	// Check mongodb connection
	mongodb_host := beego.AppConfig.String("mongodb.url")
	session, _ := mongodb.NewMongoDbSession(mongodb_host)
	log.Info(session)
	defer session.Close()
}
