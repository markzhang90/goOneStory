package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}


func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type HelloController struct {
	beego.Controller
}

func (c *HelloController) Get() {
	varId := c.Ctx.Input.Param(":id")
	varTest := c.Ctx.Input.Param(":test")
	c.Ctx.WriteString("hello world" + varId + varTest)
}

