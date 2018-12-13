package middleware

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func Cors() beego.FilterFunc {
	return cors.Allow(&cors.Options {
		AllowAllOrigins	: true,
	// 		AllowCredentials: true,
	// 		AllowMethods	: []string{"*"},
	// 		AllowHeaders	: []string{"Accept", "Access-Control-Allow-Origin", "Authorization", "Content-Type", "Origin"},
	// 		ExposeHeaders	: []string{"Access-Control-Allow-Origin", "application/json", "Content-Length", "Content-Type" },
		})
}