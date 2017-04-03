package routers

import (
	"onestory/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hello/?:id/?:test", &controllers.HelloController{})
}

