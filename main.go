package main

import (
	_ "onestory/routers"
	"github.com/astaxie/beego"
	"onestory/controllers"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"time"
	"github.com/astaxie/beego/logs"
)



func main() {
	//default
	orm.RegisterDataBase("default", "mysql", "root:@/test?charset=utf8")

	beego.ErrorController(&controllers.ErrorController{})


	//err log config
	projectName := "./logs/" + beego.AppConfig.String("appname") + "." + time.Now().Format("2006-01-02-15")
	beego.SetLogger(logs.AdapterFile, `{"filename":"`+projectName+`","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":3}`)

	//beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	//beego.BConfig.WebConfig.Session.SessionProviderConfig = "192.168.121.128:6379"
	beego.Run()
}
