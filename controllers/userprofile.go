package controllers

import (
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/logs"
	"onestory/models"
	"github.com/astaxie/beego/logs"
)

type UserProfileController struct {
	beego.Controller
}


func (c *UserProfileController) Get() {

	email := c.GetString("emial")
	phone, _ := c.GetInt("phone")

	userData := models.UserProfileData{
		Passid:12345,
		Email:email,
		Phone:phone,
		Update_time:1234561,
		Nick_name:"1234",
		Ext:"",
	}
	logs.Warn(userData)
	c.EnableXSRF = false
	var newUser = models.NewUser()
	//var getUser = newUser.GetUserProfile()
	//logs.Warning(getUser)
	newUser.AddNewUserProfile(userData)
	//redirect
	c.Abort("404")
	//c.Ctx.Redirect(302, "/hello/123/5")
}
