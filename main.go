package main

import (
	_ "onestory/routers"
	"github.com/astaxie/beego"
	"onestory/controllers"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)



func main() {
	//default
	orm.RegisterDataBase("default", "mysql", "root:@/test?charset=utf8")

	beego.ErrorController(&controllers.ErrorController{})
	beego.SetLogger("file", `{"filename":"./logs/test.log"}`)

	//beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	//beego.BConfig.WebConfig.Session.SessionProviderConfig = "192.168.121.128:6379"
	beego.Run()
}
