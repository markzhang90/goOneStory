package controllers

import (
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/logs"
	"onestory/models"
	"github.com/astaxie/beego/logs"
	"time"
	"onestory/library"
	"errors"
)

type (
	AddUserProfileController struct {
		beego.Controller
	}
	GetUserProfileController struct {
		beego.Controller
	}
	LoginUserController struct {
		beego.Controller
	}
)

//新增用户
func (c *AddUserProfileController) Get() {

	email := c.GetString("email")
	phone, _ := c.GetInt64("phone")
	userData := models.UserProfile{
		Passid: models.GetPid(phone, email),
		Email:email,
		Phone:phone,
		Password:"123",
		Update_time:time.Now().Unix(),
		Nick_name:"1234",
		Ext:"",
	}
	logs.Warn(userData)
	c.EnableXSRF = false
	var newUserDb = models.NewUser()
	//var getUser = newUser.GetUserProfile()
	//logs.Warning(getUser)
	res, err := newUserDb.AddNewUserProfile(userData)

	var output string

	if err == nil{
		output , _ = library.ReturnJsonWithError(0, "ref", res)
	}else {
		output , _ = library.ReturnJsonWithError(1, "ref", err.Error())
	}
	c.Ctx.WriteString(output)
	return
}


//登录
func (c *LoginUserController) Get()  {
	cookiekey := beego.AppConfig.String("passid")

	password := c.GetString("password")
	phone, _ := c.GetInt64("phone")

	c.EnableXSRF = false
	var newUserDb = models.NewUser()

	res, err := newUserDb.LoginUser(phone, password)

	var output string
	if err == nil{
		output , _ = library.ReturnJsonWithError(0, "ref", res)
		_, cacheUser := models.SyncSetUserCache(res)
		if cacheUser{
			c.SetSecureCookie(cookiekey, "passid", res.Passid)
		}

	}else {
		errCode := library.GetUserFail
		output , _ = library.ReturnJsonWithError(errCode, "ref", err.Error())
	}
	c.Ctx.WriteString(output)
	return
}

//获取用户信息
func (c *GetUserProfileController) Get() {

	cookiekey := beego.AppConfig.String("passid")

	var finalErr error
	passId, resBool := c.GetSecureCookie(cookiekey, "passid")
	if resBool {
		cahchedUser, err := models.GetUserFromCache(passId)
		if err == nil{
			output , _ := library.ReturnJsonWithError(0, "ref", cahchedUser)
			c.Ctx.WriteString(output)
			return
		}
	}else{
		phone, errPhone := c.GetInt64("phone")
		email := c.GetString("email")
		logs.Warning(email)
		var newUserDb = models.NewUser()
		c.EnableXSRF = false

		if errPhone != nil {
			finalErr = errPhone
		}else{
			//var getUser = newUser.GetUserProfile()
			//logs.Warning(getUser)
			userProfile, errGetUser := newUserDb.GetUserProfileByPhone(phone)
			if errGetUser != nil {
				finalErr = errGetUser
			}else{
				cacheUserObj, cacheUserRes := models.SyncSetUserCache(userProfile)
				if cacheUserRes{
					c.SetSecureCookie(cookiekey, "passid", userProfile.Passid)
				}
				output , _ := library.ReturnJsonWithError(0, "ref", cacheUserObj)
				c.Ctx.WriteString(output)
				return
			}
		}

		if email == "" {
			finalErr = errors.New("获取用户失败")
		}else{
			//var getUser = newUser.GetUserProfile()
			//logs.Warning(getUser)
			userProfile, errGetUser := newUserDb.GetUserProfileByEmail(email)
			if errGetUser != nil {
				finalErr = errGetUser
			}else{
				cacheUserObj, cacheUserRes := models.SyncSetUserCache(userProfile)
				if cacheUserRes{
					c.SetSecureCookie(cookiekey, "passid", userProfile.Passid)
				}
				output , _ := library.ReturnJsonWithError(0, "ref", cacheUserObj)
				c.Ctx.WriteString(output)
				return
			}
		}



	}
	if finalErr == nil{
		finalErr = errors.New("获取用户失败")
	}
	errCode := library.GetUserFail
	output , _ := library.ReturnJsonWithError(errCode, "ref", finalErr.Error())
	c.Ctx.WriteString(output)
	return
}

