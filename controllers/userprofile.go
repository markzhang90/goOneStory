package controllers

import (
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/logs"
	"onestory/models"
	"github.com/astaxie/beego/logs"
	"time"
	"encoding/json"
)

type UserProfileController struct {
	beego.Controller
}

func (c *UserProfileController) Get() {

	email := c.GetString("email")
	phone, _ := c.GetInt64("phone")
	userData := models.UserProfile{
		Passid: models.GetPid(phone, email),
		Email:email,
		Phone:phone,
		Update_time:time.Now().Unix(),
		Nick_name:"1234",
		Ext:"",
	}
	logs.Warn(userData)
	c.EnableXSRF = false
	var newUserDb = models.NewUser()
	//var getUser = newUser.GetUserProfile()
	//logs.Warning(getUser)
	newUserDb.AddNewUserProfile(userData)
	//redirect
	c.Abort("404")
	//c.Ctx.Redirect(302, "/hello/123/5")
}

type GetUserProfileController struct {
	beego.Controller
}

func (c *GetUserProfileController) Get() {

	phone, _ := c.GetInt64("phone")

	c.EnableXSRF = false
	var newUserDb = models.NewUser()
	//var getUser = newUser.GetUserProfile()
	//logs.Warning(getUser)
	userProfile, err := newUserDb.GetUserProfileByPhone(phone)

	if err != "" {
		c.Ctx.WriteString("hello world" + string(err))
	}

	userJson, _ := json.Marshal(userProfile)
	c.Ctx.WriteString("hello world" + string(userJson))

}

