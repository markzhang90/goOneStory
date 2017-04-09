package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
	"onestory/services"
	"crypto/md5"
	"encoding/hex"
)

type (
	UserProfileDb struct {
		tableName string
		DbConnect *services.DbService
	}

	UserProfile struct {
		Id          int `orm:"auto"`
		Passid      string
		Email       string
		Phone       int64
		Update_time int64
		Nick_name   string
		Ext         string
	}
)

func (userDb *UserProfileDb) GetUserProfileByPhone(phone int64) (targetUser UserProfile, errcode string) {
	o := userDb.DbConnect.Orm
	o.Using(userDb.DbConnect.DbName)
	targetUser = UserProfile{Phone: phone}
	err := o.Read(&targetUser, "phone")

	if err == orm.ErrNoRows {
		errcode = "查询不到"
	} else if err == orm.ErrMissPK {
		errcode = "找不到key"
	}else{
		errcode = ""
	}

	return targetUser, errcode
}

func GetPid(phone int64, email string) string {
	passIdEncode := md5.New()
	passIdEncode.Write([]byte(string(phone) + "_" + email))
	passId := hex.EncodeToString(passIdEncode.Sum(nil))
	return passId
}

func (userDb *UserProfileDb) AddNewUserProfile(userprofileData UserProfile) {
	o := userDb.DbConnect.Orm
	o.Using(userDb.DbConnect.DbName)

	profile := new(UserProfile)
	profile.Passid = userprofileData.Passid
	profile.Email = userprofileData.Email
	profile.Phone = userprofileData.Phone
	profile.Update_time = userprofileData.Update_time
	profile.Nick_name = userprofileData.Nick_name
	profile.Ext = userprofileData.Ext
	res, err := o.Insert(profile)
	logs.Warning(res)
	logs.Warning("err", err)

	if err != nil {
		logs.Warning(res)
	}

}

func NewUser() (*UserProfileDb) {

	dbService := services.NewService("onestory")
	logs.Warning(dbService)
	return &UserProfileDb{"user_profile", dbService}
}

func (userDb *UserProfileDb) GetUserProfile() (err error) {
	o := userDb.DbConnect.Orm
	o.Using(userDb.DbConnect.DbName)
	logs.Warning(userDb.DbConnect.DbName)

	var maps []orm.Params
	res, err := o.Raw("select * from user_profile where nick_name = ?", "oooook").Values(&maps)

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
