package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
	"onestory/services"
)

type (
	UserProfile struct {
		tableName string
		DbConnect *services.DbService
	}

	UserProfileData struct {
		Id int
		Passid int
		Email string
		Phone int
		Update_time int
		Nick_name string
		Ext string
	}
)

func (user *UserProfile) AddNewUserProfile(userprofileData UserProfileData)  {
	o := user.DbConnect.Orm
	o.Using(user.DbConnect.DbName)

	profile := new(UserProfileData)
	profile.Passid = userprofileData.Passid
	profile.Email = userprofileData.Email
	profile.Phone = userprofileData.Phone
	profile.Update_time = userprofileData.Update_time
	profile.Nick_name = userprofileData.Nick_name
	profile.Ext = userprofileData.Ext
	res, err := o.Insert(profile)
	logs.Warning(res)
	logs.Warning(err)

	if err != nil {
		logs.Warning(res)
	}

}


func NewUser() (*UserProfile) {

	dbService := services.NewService("onestory_main")
	logs.Warning(dbService)
	return &UserProfile{"onestory_userprofile",dbService}
}

func (user *UserProfile) GetUserProfile() (err error) {
	o := user.DbConnect.Orm
	o.Using(user.DbConnect.DbName)
	logs.Warning(user.DbConnect.DbName)

	var maps []orm.Params
	res, err := o.Raw("select * from onestory_userprofile where nick_name = ?", "oooook").Values(&maps)

	if err == nil && res > 0 {
		//data := maps[0]["email"]
		//logs.Warning(data)
		for key, v := range maps {
			logs.Warning(key)
			logs.Warning(v)
		}
	}

	return err
}
