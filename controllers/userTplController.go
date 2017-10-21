package controllers

import (
	"github.com/astaxie/beego"
	"html/template"
)

type (
	RegisterController struct {
		beego.Controller
	}
	LoginUserController struct {
		beego.Controller
	}
	EditController struct {
		beego.Controller
	}
	ShowController struct {
		beego.Controller
	}
	ShowDetailController struct {
		beego.Controller
	}
	MainController struct {
		beego.Controller
	}
)

func (c *RegisterController) Get() {
	c.Data["xsrfdata"]= template.HTML(c.XSRFFormHTML())
	c.Layout = "onestory/base.html"
	c.TplName = "onestory/register.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Fixheader"] = "onestory/fixheader.html"
	c.LayoutSections["Footer"] = "onestory/footer.html"
}


//登录渲染页
func (c *LoginUserController) Get() {
	c.Data["xsrfdata"]= template.HTML(c.XSRFFormHTML())
	c.Layout = "onestory/base.html"
	c.TplName = "onestory/login.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Fixheader"] = "onestory/fixheader.html"
	c.LayoutSections["Footer"] = "onestory/footer.html"
	return
}


func (c *EditController) Get() {
	c.Data["xsrfdata"]= template.HTML(c.XSRFFormHTML())
	c.Layout = "onestory/base.html"
	c.TplName = "onestory/edit.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Fixheader"] = "onestory/fixheader.html"
	c.LayoutSections["Footer"] = "onestory/footer.html"
}

func (c *ShowController) Get() {
	c.Data["xsrfdata"]= template.HTML(c.XSRFFormHTML())
	c.Data["curId"]= ""
	c.Data["detail"]= false
	c.Layout = "onestory/base.html"
	c.TplName = "onestory/show.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Fixheader"] = "onestory/fixheader.html"
	c.LayoutSections["Footer"] = "onestory/footer.html"
}

func (c *ShowDetailController) Get() {
	getId := c.Ctx.Input.Param(":id")
	c.Data["xsrfdata"]= template.HTML(c.XSRFFormHTML())
	c.Data["curId"]= getId
	c.Data["detail"]= true
	c.Layout = "onestory/base.html"
	c.TplName = "onestory/show.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Fixheader"] = "onestory/fixheader.html"
	c.LayoutSections["Footer"] = "onestory/footer.html"
}


func (c *MainController) Get() {
	c.Data["xsrfdata"]= template.HTML(c.XSRFFormHTML())
	c.Layout = "onestory/base.html"
	c.TplName = "onestory/feed.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Fixheader"] = "onestory/fixheader.html"
	c.LayoutSections["Footer"] = "onestory/footer.html"
}
