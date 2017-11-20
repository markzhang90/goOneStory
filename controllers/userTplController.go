package controllers

import (
	"github.com/astaxie/beego"
	"html/template"
	"onestory/library"
	"onestory/models"
)

type (
	RegisterController struct {
		beego.Controller
	}
	ProfileController struct {
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
	ActiveUserProfileController struct {
		beego.Controller
	}
	
	IndexController struct {
		beego.Controller
	}
)

func (c *ActiveUserProfileController) Get() {
	key := c.GetString("key")
	c.Data["xsrfdata"]= template.HTML(c.XSRFFormHTML())
	c.Layout = "onestory/base.html"
	c.TplName = "onestory/msgalert.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Fixheader"] = "onestory/fixheader.html"
	c.LayoutSections["Footer"] = "onestory/footer.html"

	if len(key) < 42 {
		c.Data["msg"] = "激活码无效";
	}

	realPassId := library.Substr(key, 10, 42)
	var newUserDb = models.NewUser()
	userProFile, err := newUserDb.GetUserProfileByPassId(realPassId)
	if err != nil {
		c.Data["msg"] = "激活码无效";
	}
	_, errUp := newUserDb.ActiveUserProfile(userProFile)

	if errUp != nil {
		c.Data["msg"] = "激活码失败";
	}
	if _, ok := c.Data["msg"]; ok {
		//存在
		c.Data["header"] = "oops！";
		c.StopRun();
		return;
	}
	c.Data["msg"] = "激活成功!!";
	c.Data["header"] = "";
}

func (c *RegisterController) Get() {
	c.Data["xsrfdata"]= template.HTML(c.XSRFFormHTML())
	c.Layout = "onestory/base.html"
	c.TplName = "onestory/register.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Fixheader"] = "onestory/fixheader.html"
	c.LayoutSections["Footer"] = "onestory/footer.html"
}

func (c *ProfileController) Get() {
	c.Data["xsrfdata"]= template.HTML(c.XSRFFormHTML())
	c.Layout = "onestory/base.html"
	c.TplName = "onestory/profile.html"
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

func (c *IndexController) Get() {
	c.Data["xsrfdata"]= template.HTML(c.XSRFFormHTML())
	c.TplName = "onestory/index.html"
}
