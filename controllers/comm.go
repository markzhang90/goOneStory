package controllers

import (
	"github.com/astaxie/beego"
	"onestory/library"
	"onestory/models"
	"strings"
	"time"
	"onestory/services/request/third"
)

type (
	EmailConfirmController struct {
		beego.Controller
	}

	SendEmailAuthController struct {
		beego.Controller
	}
)

/**
not used
 */
func (c *EmailConfirmController) Get() {
	cookiekey := beego.AppConfig.String("passid")

	//get from cache
	passId, _ := c.GetSecureCookie(cookiekey, "passid")

	if len(passId) <= 0 {
		passId = c.GetString("passid", "")
		if len(passId) < 1 {
			output, _ := library.ReturnJsonWithError(library.GetUserFail, "ref", nil)
			c.Ctx.WriteString(output)
			c.StopRun()
			return
		}
	}

	cachedUser, err := models.GetUserFromCache(passId, false)
	if err != nil {
		output, _ := library.ReturnJsonWithError(library.GetUserFail, "ref", err.Error())
		c.Ctx.WriteString(output)
		return
	}

	email := cachedUser.Email
	subject := "激活账户通知"
	openUrl := "https://onestory.cn/user/activeuser?key=" + strings.ToLower(library.RandSeq(10)) + passId + strings.ToLower(library.RandSeq(10))
	message := "<html><body><a href='" + openUrl + "'>注册成功，请点击链接激活账户<a> <br> 或复制以下链接至浏览器 " + openUrl + " </body></html>"
	errEmail := third.SendToMail(email, subject, message, "html")
	if errEmail != nil {
		output, _ := library.ReturnJsonWithError(library.CodeErrApi, "发送激活邮件失败，请检查邮箱是否正确", "")
		c.Ctx.WriteString(output)
		c.StopRun()
		return
	}

	output, _ := library.ReturnJsonWithError(library.CodeSucc, "ref", "")
	c.Ctx.WriteString(output)
	c.StopRun()
	return
}


func (c *SendEmailAuthController) Get() {
	email := c.GetString("email", "")

	if len(email) < 5 {
		output, _ := library.ReturnJsonWithError(library.CodeErrCommen, "email 无效", nil)
		c.Ctx.WriteString(output)
		c.StopRun()
		return
	}
	randNum := library.RandMix(6)
	now := time.Now().Unix()
	newAuthConf := models.AuthCode{
		Email: email,
		Update_time:now,
		Create_time:now,
		Code:randNum,
	}

	authModel := models.NewAuthCode()
	resGetAll, errQuery := authModel.QueryGetAuthCodeByEmail(email)
	if errQuery == nil || len(resGetAll) > 2{
		output, _ := library.ReturnJsonWithError(library.CodeErrApi, "点的太快啦，请稍后再试", "")
		c.Ctx.WriteString(output)
		c.StopRun()
		return
	}
	_, err := authModel.AddNewAuthCode(newAuthConf)
	if err != nil {
		output, _ := library.ReturnJsonWithError(library.CodeErrApi, "验证码发送失败，请稍后再试", "")
		c.Ctx.WriteString(output)
		c.StopRun()
		return
	}
	subject := "一日记邮箱验证码"
	body := "<html><body>您的验证码为 : <br> <h1>"+randNum+"</h1></body></html>"
	mailErr := third.SendToMail(email, subject, body, "html")
	if mailErr != nil {
		output, _ := library.ReturnJsonWithError(library.CodeErrApi, "验证码发送失败,请确认邮箱正确", "")
		c.Ctx.WriteString(output)
		c.StopRun()
		return
	}
	output, _ := library.ReturnJsonWithError(library.CodeSucc, "succ", randNum)
	c.Ctx.WriteString(output)
	c.StopRun()
	return
}
