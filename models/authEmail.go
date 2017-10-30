package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
	"onestory/services"
	"github.com/mitchellh/mapstructure"
	"encoding/json"
	"errors"
	"time"
)

type (
	authCodeDb struct {
		tableName string
		DbConnect *services.DbService
	}

	AuthCode struct {
		Id          int `orm:"auto"`
		Email       string
		Code      	string
		Create_time int64
		Update_time int64
		status      int
	}
)

func NewAuthCode() (*authCodeDb) {

	dbService, err := services.NewService("onestory")
	if err != nil {
		logs.Warn(err)
	}
	return &authCodeDb{"auth_code", dbService}

}

func (authCodeDb *authCodeDb)GetAuthCodeByPassidAndId(uid int, id int) (AuthCodeList []AuthCode, err error) {
	o := authCodeDb.DbConnect.Orm
	o.Using(authCodeDb.DbConnect.DbName)

	var maps []orm.Params
	var allResAuthCode []AuthCode

	qs := o.QueryTable(authCodeDb.tableName).Filter("uid", uid).Filter("id", id)
	_, err = qs.Values(&maps)
	if len(maps) < 1 {
		err = errors.New("no record")
	}
	if err == nil {
		for _, AuthCode := range maps {
			//logs.Warning(AuthCode)
			eachAuthCode, err := _assignMapToAuthCode(AuthCode)
			if err != nil {
				logs.Warning(err)
			} else {
				allResAuthCode = append(allResAuthCode, eachAuthCode)
			}

		}
		return allResAuthCode, nil
	}
	return allResAuthCode, err
}

/**
query count by uid
 */
func (authCodeDb *authCodeDb) QueryGetAuthCodeByEmail(email string) ([]AuthCode, error) {
	o := authCodeDb.DbConnect.Orm
	o.Using(authCodeDb.DbConnect.DbName)
	now := time.Now()
	duration, _ := time.ParseDuration("-2m")
	tenMago := now.Add(duration).Unix()
	var eachAuth AuthCode
	var eachAuthAll []AuthCode
	var maps []orm.Params
	qsGet := o.QueryTable(authCodeDb.tableName).Filter("email", email).Filter("Create_time__gte", tenMago).OrderBy("-id").Limit(1)
	_, err := qsGet.Values(&maps)
	if err == nil {
		if len(maps) < 1 {
			return eachAuthAll, orm.ErrNoRows
		}
		for _, authInfo := range maps {
			//logs.Warning(posts)
			eachAuth, err = _assignMapToAuthCode(authInfo)
			if err != nil {
				logs.Warning(err)
				continue
			}
			eachAuthAll = append(eachAuthAll, eachAuth)
		}
		return eachAuthAll, nil
	}
	logs.Warning("get user AuthCode count fail " + string(email) + " error : " + err.Error())
	return eachAuthAll, err
}

/**
add a new AuthCode
 */
func (authCodeDb *authCodeDb) AddNewAuthCode(NewAuthCode AuthCode) (authId int64, err error) {
	o := authCodeDb.DbConnect.Orm
	o.Using(authCodeDb.DbConnect.DbName)

	authId = -1
	mycode := new(AuthCode)
	mycode.Email = NewAuthCode.Email
	mycode.Code = NewAuthCode.Code
	mycode.Update_time = NewAuthCode.Update_time
	mycode.Create_time = NewAuthCode.Create_time
	res, err := o.Insert(mycode)
	
	if err != nil {
		jsonData, _ := json.Marshal(mycode)
		logs.Warning("AuthCode fail %s error: %s ", jsonData, err.Error())
	} else {
		authId = res
	}

	return authId, err
}


func _assignMapToAuthCode(fromMap orm.Params) (eachAuthCode AuthCode, err error) {
	err = mapstructure.Decode(fromMap, &eachAuthCode)
	if err != nil {
		return eachAuthCode, err
		panic(err)
	}
	return eachAuthCode, nil
}
