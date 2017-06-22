package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
	"onestory/services"
	"github.com/mitchellh/mapstructure"
	"encoding/json"
)

type (
	PostDb struct {
		tableName string
		DbConnect *services.DbService
	}

	Posts struct {
		Id          int `orm:"auto"`
		Uid         int
		Header      string
		Rel         string
		Create_date int64
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

/**
get most recent post
 */
func (postDb *PostDb) GetUserClosestPost(uid int, givenDate int, isNext bool) (postList Posts, err error) {
	o := postDb.DbConnect.Orm
	o.Using(postDb.DbConnect.DbName)
	var queryGet string
	if isNext {
		queryGet = "select * from "+postDb.tableName+" where uid = ? and create_date > ? order by create_date asc limit 1";
	}else{
		queryGet = "select * from "+postDb.tableName+" where uid = ? and create_date < ? order by create_date desc limit 1";
	}
	var posts Posts
	logs.Warning(o.Raw(queryGet, uid, givenDate))
	err = o.Raw(queryGet, uid, givenDate).QueryRow(&posts)

	return posts, err
}

/**
get all user posts
 */
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
}

/**
add a new post
 */
func (postDb *PostDb) AddNewUserPost(NewPost Posts) (postId int64, err error) {
	o := postDb.DbConnect.Orm
	o.Using(postDb.DbConnect.DbName)

	postId = -1
	mypost := new(Posts)
	mypost.Uid = NewPost.Uid
	mypost.Header = NewPost.Header
	mypost.Rel = NewPost.Rel
	mypost.Content = NewPost.Content
	mypost.Update_time = NewPost.Update_time
	mypost.Create_date = NewPost.Create_date
	res, err := o.Insert(mypost)

	if err != nil {
		jsonData,_ := json.Marshal(mypost)
		logs.Warning("post fail %s error: %s ", jsonData, err.Error())
	} else {
		postId = res
	}

	return postId, err
}
