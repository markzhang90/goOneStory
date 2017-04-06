package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(UserProfileData))
}
