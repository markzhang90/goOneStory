package controllers

import (
	"github.com/astaxie/beego"
	"onestory/library"
	"onestory/services/request/third"
	"github.com/astaxie/beego/logs"
	"onestory/models"
	"time"
	"strconv"
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

	clearRes := userDb.ClearProfileOut(userprofile)
	var uid = userprofile.Id
	userPost := models.NewPost()
	countAll, err := userPost.QueryCountUserPost(uid);

	if err != nil {
		countAll = -1
	}
	//var today = time.Now().Format("20060102");
	//
	//todayInt, _ := strconv.Atoi(today)
	//
	//todayArr := []int{todayInt}
	//
	//result, errGet := userPost.QueryUserPostByDate(userprofile.Id, todayArr, true, 1);
	//
	//clearRes["Today"] = false;
	//if errGet != nil {
	//	if len(result) > 0 {
	//		clearRes["Today"] = true;
	//	}
	//}

	clearRes["Post_count"] = countAll
	output, _ := library.ReturnJsonWithError(0, "", clearRes)
	c.Ctx.WriteString(output)
	return
}

