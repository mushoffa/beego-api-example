package main

import (
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/plugins/cors"
	
	_ "beego-api-example/routers"
	_ "beego-api-example/pkg/database/mongo"
	"beego-api-example/pkg/middleware"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter, middleware.Cors())

	beego.Run()
}
