package common

import (
	"onestory/models"
	"onestory/library"
	"strings"
)

func SendRegisterEmail(userProfile models.UserProfile) error {

	email := userProfile.Email
	subject := "激活账户通知"
	openUrl :=  "https://onestory.cn/user/activeuser?key=" + strings.ToLower(library.RandSeq(10)) + userProfile.Passid + strings.ToLower(library.RandSeq(10))
	message := "<html><body><a href='"+openUrl+"'>注册成功，请点击链接激活账户<a> <br> 或复制以下链接至浏览器 " + openUrl +" </body></html>"
	errEmail := library.SendToMail(email, subject, message, "html")
	return errEmail
}
