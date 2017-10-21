package controllers

import (
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/logs"
	"onestory/models"
	"github.com/astaxie/beego/logs"
	"time"
	"onestory/library"
	"errors"
	"onestory/services/common"
)

type (
	AddUserProfileController struct {
		beego.Controller
	}
	GetUserProfileController struct {
		beego.Controller
	}
	LogoutUserController struct {
		beego.Controller
	}
	UpdateUserProfileController struct {
		beego.Controller
	}
	GetUserProfileInfoController struct {
		beego.Controller
	}
)


//新增用户
func (c *AddUserProfileController) Post() {
	c.EnableXSRF = false
	email := c.GetString("email")
	phone, _ := c.GetInt64("phone", 0)
	avatar := c.GetString("avatar", "")
	nickname := c.GetString("nickname", "")
	password := c.GetString("password", "")

	var output string

	if len(password) < 6 {
		output, _ = library.ReturnJsonWithError(1, "密码不可少于6位", "")
		c.Ctx.WriteString(output)
		return
	}
	if len(nickname) <= 0 {
		output, _ = library.ReturnJsonWithError(1, "昵称不能为空", "")
		c.Ctx.WriteString(output)
		return
	}
	if len(email) <= 0 {
		output, _ = library.ReturnJsonWithError(1, "邮箱不能为空", "")
		c.Ctx.WriteString(output)
		return
	}

	userData := models.UserProfile{
		Passid:      models.GetPid(phone, email),
		Email:       email,
		Phone:       phone,
		Openid:       "0",
		Password:    password,
		Update_time: time.Now().Unix(),
		Nick_name:   nickname,
		Avatar:   	 avatar,
		Ext:         "",
		Active:      0,
	}
	var newUserDb = models.NewUser()
	//var getUser = newUser.GetUserProfile()
	//logs.Warning(getUser)
	res, err := newUserDb.AddNewUserProfile(userData)

	if err == nil {
		targetUser, errGet := newUserDb.GetUserProfileById(int(res))
		if errGet != nil{
			output, _ = library.ReturnJsonWithError(1, "ref", err.Error())
		}
		cookiekey := beego.AppConfig.String("passid")
		models.SyncSetUserCache(targetUser, false)
		logs.Warning(targetUser.Passid)
		c.SetSecureCookie(cookiekey, "passid", "")
		c.SetSecureCookie(cookiekey, "passid", targetUser.Passid)
		errEmail := common.SendRegisterEmail(targetUser)
		if errEmail != nil{
			output, _ = library.ReturnJsonWithError(1, "发送激活邮件失败，请检查邮箱是否填写正确", "")
		}else{
			output, _ = library.ReturnJsonWithError(0, "ref", true)
		}
	} else {
		if err.Error() == "user exist" {
			output, _ = library.ReturnJsonWithError(1, "用户信息已被注册", err.Error())
		}else{
			output, _ = library.ReturnJsonWithError(1, "ref", err.Error())
		}
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
	avatar := c.GetString("avatar")

	userData := models.UserProfile{
		Id:          id,
		//Passid:      models.GetPid(phone, email),
		//Email:       email,
		//Phone:       phone,
		Password:    password,
		Nick_name:   nickname,
		Avatar:   	 avatar,
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
		cacheUserObj, cacheUserRes := models.SyncSetUserCache(resUser, false)

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

	password := c.GetString("password", "")
	email := c.GetString("email", "")

	c.EnableXSRF = false
	var newUserDb = models.NewUser()

	res, err := newUserDb.LoginUserByEmail(email, password)
	logs.Warn(res)
	var output string
	if err == nil {
		output, _ = library.ReturnJsonWithError(0, "ref", res)
		_, cacheUser := models.SyncSetUserCache(res, false)
		if cacheUser {
			//set redis fail
		}
		cookiekey := beego.AppConfig.String("passid")
		c.SetSecureCookie(cookiekey, "passid", res.Passid)
	} else {
		errCode := library.GetUserFail
		output, _ = library.ReturnJsonWithError(errCode, "ref", err.Error())
	}
	c.Ctx.WriteString(output)
	return
}

//获取用户信息
func (c *GetUserProfileController) Get() {

	c.EnableXSRF = false

	cookiekey := beego.AppConfig.String("passid")

	var finalErr error

	//get from cache
	passId, resBool := c.GetSecureCookie(cookiekey, "passid")
	if resBool {
		cahchedUser, err := models.GetUserFromCache(passId, true)
		if err == nil {
			output, _ := library.ReturnJsonWithError(0, "ref", cahchedUser)
			c.Ctx.WriteString(output)
			return
		}
	}

	phone, errPhone := c.GetInt64("phone")
	email := c.GetString("email")
	var newUserDb = models.NewUser()

	if errPhone != nil {
		finalErr = errPhone
	} else {
		//var getUser = newUser.GetUserProfile()
		//logs.Warning(getUser)
		userProfile, errGetUser := newUserDb.GetUserProfileByPhone(phone)
		if errGetUser != nil {
			finalErr = errGetUser
		} else {
			cacheUserObj, cacheUserRes := models.SyncSetUserCache(userProfile, false)
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
			cacheUserObj, cacheUserRes := models.SyncSetUserCache(userProfile, false)
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
