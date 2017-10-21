package controllers

import (
	"github.com/astaxie/beego"
	"onestory/library"
	"onestory/models"
	"strings"
	"github.com/astaxie/beego/logs"
)

type (
	EmailConfirmController struct {
		beego.Controller
	}
)


func (c *EmailConfirmController) Get() {
	cookiekey := beego.AppConfig.String("passid")

	//get from cache
	passId, _ := c.GetSecureCookie(cookiekey, "passid")

	if len(passId) <= 0 {
		passId = c.GetString("passid", "")
		if len(passId) < 1{
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
	openUrl :=  "https://onestory.cn/user/activeuser?key=" + strings.ToLower(library.RandSeq(10)) + passId + strings.ToLower(library.RandSeq(10))
	message := "<html><body><a href='"+openUrl+"'>注册成功，请点击链接激活账户<a> <br> 或复制以下链接至浏览器 " + openUrl +" </body></html>"
	errEmail := library.SendToMail(email, subject, message, "html")
	if errEmail != nil{
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