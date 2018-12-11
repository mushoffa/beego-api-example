package main

import (
	"github.com/astaxie/beego"
	
	_ "beego-api-example/routers"
	_ "beego-api-example/pkg/database/mongo"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
