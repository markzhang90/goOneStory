package controllers

import (
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/logs"
	"onestory/models"
	"github.com/astaxie/beego/logs"
	"time"
	"onestory/library"
	"errors"
	"html/template"
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
	LogoutUserController struct {
		beego.Controller
	}
	UpdateUserProfileController struct {
		beego.Controller
	}
)

//新增用户
func (c *AddUserProfileController) Get() {

	email := c.GetString("email")
	phone, _ := c.GetInt64("phone")
	userData := models.UserProfile{
		Passid:      models.GetPid(phone, email),
		Email:       email,
		Phone:       phone,
		Password:    "123",
		Update_time: time.Now().Unix(),
		Nick_name:   "1234",
		Ext:         "",
	}
	c.EnableXSRF = false
	var newUserDb = models.NewUser()
	//var getUser = newUser.GetUserProfile()
	//logs.Warning(getUser)
	res, err := newUserDb.AddNewUserProfile(userData)

	var output string

	if err == nil {
		output, _ = library.ReturnJsonWithError(0, "ref", res)
	} else {
		output, _ = library.ReturnJsonWithError(1, "ref", err.Error())
	}
	c.Ctx.WriteString(output)
	return
}

//更新用户
func (c *UpdateUserProfileController) Get() {

	id, ok := c.GetInt("id")
	if ok != nil {
		output, _ := library.ReturnJsonWithError(1, "获取id失败", ok.Error())
		c.Ctx.WriteString(output)
		return
	}

	nickname := c.GetString("nickname")
	password := c.GetString("password")

	userData := models.UserProfile{
		Id:          id,
		//Passid:      models.GetPid(phone, email),
		//Email:       email,
		//Phone:       phone,
		Password:    password,
		Nick_name:   nickname,
		Ext:         "",
	}

	c.EnableXSRF = false
	var newUserDb = models.NewUser()
	//var getUser = newUser.GetUserProfile()
	//logs.Warning(getUser)
	resUser, errUpdate := newUserDb.UpdateNewUserProfile(userData)

	var output string

	if errUpdate == nil {

		cookiekey := beego.AppConfig.String("passid")
		cacheUserObj, cacheUserRes := models.SyncSetUserCache(resUser)

		if cacheUserRes {
			c.SetSecureCookie(cookiekey, "passid", resUser.Passid)
		}else{
			resBool, delError := models.CleanUserCache(resUser.Passid)
			if !resBool {
				logs.Warn(" update user cache fail " + delError.Error())
			}
		}
		output, _ := library.ReturnJsonWithError(0, "ref", cacheUserObj)
		c.Ctx.WriteString(output)
		return

	} else {
		output, _ = library.ReturnJsonWithError(1, errUpdate.Error(), errUpdate.Error())
	}
	c.Ctx.WriteString(output)
	return
}

//登录
func (c *LoginUserController) Post() {
	cookiekey := beego.AppConfig.String("passid")

	password := c.GetString("password")
	phone, _ := c.GetInt64("phone")

	c.EnableXSRF = false
	var newUserDb = models.NewUser()

	res, err := newUserDb.LoginUser(phone, password)

	var output string
	if err == nil {
		output, _ = library.ReturnJsonWithError(0, "ref", res)
		_, cacheUser := models.SyncSetUserCache(res)
		if cacheUser {
			//set redis fail
		}
		c.SetSecureCookie(cookiekey, "passid", res.Passid)


	} else {
		errCode := library.GetUserFail
		output, _ = library.ReturnJsonWithError(errCode, "ref", err.Error())
	}
	c.Ctx.WriteString(output)
	return
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

//获取用户信息
func (c *GetUserProfileController) Get() {

	cookiekey := beego.AppConfig.String("passid")

	var finalErr error

	//get from cache
	passId, resBool := c.GetSecureCookie(cookiekey, "passid")
	if resBool {
		cahchedUser, err := models.GetUserFromCache(passId)
		if err == nil {
			output, _ := library.ReturnJsonWithError(0, "ref", cahchedUser)
			c.Ctx.WriteString(output)
			return
		}
	}

	phone, errPhone := c.GetInt64("phone")
	email := c.GetString("email")
	logs.Warning(email)
	var newUserDb = models.NewUser()
	c.EnableXSRF = false

	if errPhone != nil {
		finalErr = errPhone
	} else {
		//var getUser = newUser.GetUserProfile()
		//logs.Warning(getUser)
		userProfile, errGetUser := newUserDb.GetUserProfileByPhone(phone)
		if errGetUser != nil {
			finalErr = errGetUser
		} else {
			cacheUserObj, cacheUserRes := models.SyncSetUserCache(userProfile)
			if cacheUserRes {
				c.SetSecureCookie(cookiekey, "passid", userProfile.Passid)
			}
			output, _ := library.ReturnJsonWithError(0, "ref", cacheUserObj)
			c.Ctx.WriteString(output)
			return
		}
	}

	if email == "" {
		finalErr = errors.New("获取用户失败")
	} else {
		//var getUser = newUser.GetUserProfile()
		//logs.Warning(getUser)
		userProfile, errGetUser := newUserDb.GetUserProfileByEmail(email)
		if errGetUser != nil {
			finalErr = errGetUser
		} else {
			cacheUserObj, cacheUserRes := models.SyncSetUserCache(userProfile)
			if cacheUserRes {
				c.SetSecureCookie(cookiekey, "passid", userProfile.Passid)
			}
			output, _ := library.ReturnJsonWithError(0, "ref", cacheUserObj)
			c.Ctx.WriteString(output)
			return
		}
	}

	if finalErr == nil {
		finalErr = errors.New("获取用户失败")
	}
	errCode := library.GetUserFail
	output, _ := library.ReturnJsonWithError(errCode, "ref", finalErr.Error())
	c.Ctx.WriteString(output)
	return
}
