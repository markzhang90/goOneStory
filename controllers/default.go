package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"onestory/services/rediscli"
	"html/template"
)

type (
	MainController struct {
		beego.Controller
	}
	EditController struct {
		beego.Controller
	}
	TestController struct {
		beego.Controller
	}
)


func (c *MainController) Get() {
	c.Data["xsrfdata"]= template.HTML(c.XSRFFormHTML())
	c.Layout = "onestory/base.html"
	c.TplName = "onestory/feed.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Fixheader"] = "onestory/fixheader.html"
	c.LayoutSections["Footer"] = "onestory/footer.html"
}

func (c *EditController) Get() {
	c.Data["xsrfdata"]= template.HTML(c.XSRFFormHTML())
	c.Layout = "onestory/base.html"
	c.TplName = "onestory/edit.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Fixheader"] = "onestory/fixheader.html"
	c.LayoutSections["Footer"] = "onestory/footer.html"
}


func (c *TestController) Get() {

	defer func(){ // 必须要先声明defer，否则不能捕获到panic异常
		logs.Warning("begin defer")

		if err:=recover();err!=nil{
			logs.Warning(err)
		}
		logs.Warning("ends defer")
	}()

	varId := c.Ctx.Input.Param(":id")
	varTest := c.Ctx.Input.Param(":test")

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

	defer conn.Close()

	c.Ctx.WriteString("hello world" + varId + varTest)
}

