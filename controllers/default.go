package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"onestory/models"
	"onestory/services/rediscli"
	"onestory/library"
)

type MainController struct {
	beego.Controller
}


func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
	c.EnableXSRF = false
	var newUser = models.NewUser()
	var getUser = newUser.GetUserProfile()
	logs.Warning(getUser)

	//redirect
	c.Abort("404")
	//c.Ctx.Redirect(302, "/hello/123/5")
}

type HelloController struct {
	beego.Controller
}

func (c *HelloController) Get() {
	varId := c.Ctx.Input.Param(":id")
	varTest := c.Ctx.Input.Param(":test")
	logs.Warning(c.GetString("get"))


	v := c.GetSession("asta")
	if v == nil {
		c.SetSession("asta", int(1))
		c.Data["num"] = 0
	} else {
		c.SetSession("asta", v.(int)+1)
		c.Data["num"] = v.(int)
	}
	logs.Warning("coooool")

	conn := rediscli.RedisClient.Get()
	_, err2 := conn.Do("SET", "hello", "world")
	if err2!=nil{
		panic(err2)
	}
	logs.Warning(11111)

	library.RecoverPanic()
	defer conn.Close()
	logs.Warning(err2)

	c.Ctx.WriteString("hello world" + varId + varTest)
}

