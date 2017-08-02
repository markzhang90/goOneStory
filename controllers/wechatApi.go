package controllers

import (
	"github.com/astaxie/beego"
	"onestory/library"
	"onestory/services/request/third"
	"github.com/astaxie/beego/logs"
	"onestory/models"
)

type (
	LoginWehchatController struct {
		beego.Controller
	}

)

func (c *LoginWehchatController) Get()  {
	code := c.GetString("code", "")
	if len(code) < 1 {
		output, _ := library.ReturnJsonWithError(1, "获取code失败", "")
		c.Ctx.WriteString(output)
		return
	}

	wechat := third.NewWechatSmallApp("wechat-smallapp")
	if wechat == nil {
		output, _ := library.ReturnJsonWithError(1, "获取配置失败", "")
		c.Ctx.WriteString(output)
		return
	}
	//call wehchat api
	callRes, err := wechat.GetLoginOpenIdFronCode(code);

	if err != nil{
		output, _ := library.ReturnJsonWithError(library.CodeErrApi, "微信登录失败", callRes)
		c.Ctx.WriteString(output)
		return
	}

	var openid string
	if weChatBack , ok := callRes.(map[string]interface{}); ok{
		openid = weChatBack["openid"].(string)
		logs.Warn(openid)
	}else{
		output, _ := library.ReturnJsonWithError(library.CodeErrApi, "微信登录失败", callRes)
		c.Ctx.WriteString(output)
		return
	}

	//get userInfo by openId
	userDb := models.NewUser()
	userprofile, errGetDb := userDb.GetUserByOpenIdOrCreate(openid)
	if err != nil{
		output, _ := library.ReturnJsonWithError(library.CodeErrApi, "微信登录失败", errGetDb.Error())
		c.Ctx.WriteString(output)
		return
	}

	output, _ := library.ReturnJsonWithError(0, "", userprofile)
	c.Ctx.WriteString(output)
	return
}

