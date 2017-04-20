package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
	"onestory/services"
	"github.com/mitchellh/mapstructure"
)

type (
	PostDb struct {
		tableName string
		DbConnect *services.DbService
	}

	Posts struct {
		Id          int `orm:"auto"`
		Uid         int
		Title       string
		Create_date int64
		Rel         string
		Update_time int64
		Content     string
		Ext         string
	}
)

func NewPost() (*PostDb) {

	dbService, err := services.NewService("onestory")
	if err != nil{
		logs.Warn(err)
	}
	return &PostDb{"posts", dbService}

}

func (postDb *PostDb) GetUserAllRecentPosts(uid int, limit int) (postList []Posts, err error) {

	o := postDb.DbConnect.Orm
	o.Using(postDb.DbConnect.DbName)
	var maps []orm.Params

	if limit <= 0 {
		limit = 0
	}

	qs := o.QueryTable(postDb.tableName).Filter("uid", uid).OrderBy("-update_time").Limit(limit)
	_, err = qs.Values(&maps)
	var allResPosts []Posts
	if err == nil {
		for _, posts := range maps {
			eachPost, err := _assignMapToPost(posts)
			if err != nil{
				logs.Warning(err)
			}else{
				allResPosts = append(allResPosts, eachPost)
			}

		}
		return allResPosts, nil
	}

	return allResPosts, err
}

func _assignMapToPost(fromMap orm.Params) (eachPost Posts ,err error) {
	err = mapstructure.Decode(fromMap, &eachPost)
	if err != nil {
		return eachPost, err
		panic(err)
	}
	return eachPost, nil
	//logs.Warn(fromMap["Id"])
	//var emptyRes Posts
	//id, ok := fromMap["Id"].(int)
	//if !ok {
	//	return emptyRes, ok
	//}
	//uid, ok := fromMap["Uid"].(int)
	//if !ok {
	//	return emptyRes, ok
	//}
	//title, ok := fromMap["Title"].(string)
	//if !ok {
	//	return emptyRes, ok
	//}
	//createTime, ok := fromMap["Create_date"].(int64)
	//if !ok {
	//	return emptyRes, ok
	//}
	//
	//rel,ok := fromMap["Rel"].(string)
	//if !ok {
	//	return emptyRes, ok
	//}
	//updateTime,ok := fromMap["Update_time"].(int64)
	//if !ok {
	//	return emptyRes, ok
	//}
	//content, ok := fromMap["Content"].(string)
	//if !ok {
	//	return emptyRes, ok
	//}
	//
	//ext,ok := fromMap["Ext"].(string)
	//if !ok {
	//	return emptyRes, ok
	//}
	//
	//eachPost.Id = id
	//eachPost.Uid = uid
	//eachPost.Title = title
	//eachPost.Create_date = createTime
	//eachPost.Rel = rel
	//eachPost.Update_time = updateTime
	//eachPost.Content = content
	//eachPost.Ext = ext
}

func (postDb *PostDb) AddNewUserPost(NewPost Posts) (postId int64, err error) {
	o := postDb.DbConnect.Orm
	o.Using(postDb.DbConnect.DbName)
	postId = -1
	mypost := new(Posts)
	mypost.Uid = NewPost.Uid
	mypost.Title = NewPost.Title
	mypost.Content = NewPost.Content
	mypost.Update_time = NewPost.Update_time
	mypost.Create_date = NewPost.Create_date
	res, err := o.Insert(mypost)
	logs.Warning(res)
	logs.Warning("err", err)

	if err != nil {
		logs.Warning(postId)
	} else {
		postId = res
	}

	return postId, err
}
