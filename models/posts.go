package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
	"onestory/services"
	"github.com/mitchellh/mapstructure"
	"encoding/json"
	"onestory/library"
	"errors"
)

type (
	PostDb struct {
		tableName string
		DbConnect *services.DbService
	}

	PubPost struct {
		Id          int `orm:"auto"`
		Story_id    int64
		Header      string
		Rel         string
		Update_time int64
		Content     string
		Ext         string
	}
	Posts struct {
		Id          int `orm:"auto"`
		Story_id    int64
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
	if err != nil {
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
		queryGet = "select * from " + postDb.tableName + " where uid = ? and create_date > ? order by create_date asc limit 1";
	} else {
		queryGet = "select * from " + postDb.tableName + " where uid = ? and create_date < ? order by create_date desc limit 1";
	}
	var posts Posts

	err = o.Raw(queryGet, uid, givenDate).QueryRow(&posts)

	return posts, err
}

func (postDb *PostDb)GetPostByPassidAndId(uid int, id int) (postList []Posts, err error) {
	o := postDb.DbConnect.Orm
	o.Using(postDb.DbConnect.DbName)

	var maps []orm.Params
	var allResPosts []Posts

	qs := o.QueryTable(postDb.tableName).Filter("uid", uid).Filter("id", id)
	_, err = qs.Values(&maps)
	if len(maps) < 1 {
		err = errors.New("no record")
	}
	if err == nil {
		for _, posts := range maps {
			//logs.Warning(posts)
			eachPost, err := _assignMapToPost(posts)
			if err != nil {
				logs.Warning(err)
			} else {
				allResPosts = append(allResPosts, eachPost)
			}

		}
		return allResPosts, nil
	}
	return allResPosts, err
}

/**
post date list
 */
func (postDb *PostDb) QueryUserPostByDate(uid int, queryDateArr []int, orderByDesc bool, limit int) (postList []Posts, err error) {

	o := postDb.DbConnect.Orm
	o.Using(postDb.DbConnect.DbName)
	//var queryGet string
	var orderby string
	if orderByDesc {
		orderby = "-create_date"
	} else {
		orderby = "create_date"
	}
	var maps []orm.Params
	var allResPosts []Posts
	qs := o.QueryTable(postDb.tableName).Filter("uid", uid).Filter("create_date__in", queryDateArr).OrderBy(orderby).Limit(limit)
	_, err = qs.Values(&maps)

	if err == nil {
		for _, posts := range maps {
			//logs.Warning(posts)
			eachPost, err := _assignMapToPost(posts)
			if err != nil {
				logs.Warning(err)
			} else {
				allResPosts = append(allResPosts, eachPost)
			}

		}

		return allResPosts, nil
	}
	return allResPosts, err
}

/**
query count
 */
func (postDb *PostDb) QueryCountUserPostByDateRange(uid int, startDate int, endDate int) (num int64, err error) {
	o := postDb.DbConnect.Orm
	o.Using(postDb.DbConnect.DbName)
	qsNum, err := o.QueryTable(postDb.tableName).Filter("uid", uid).Filter("create_date__gte", startDate).Filter("create_date__lte", endDate).Count()
	if err == nil {
		return qsNum, nil
	}
	return -1, err
}

/**
query count by uid
 */
func (postDb *PostDb) QueryCountUserPost(uid int) (num int64, err error) {
	o := postDb.DbConnect.Orm
	o.Using(postDb.DbConnect.DbName)
	qsNum, err := o.QueryTable(postDb.tableName).Filter("uid", uid).Count()
	if err == nil {
		return qsNum, nil
	}
	logs.Warning("get user post count fail " + string(uid) + " error : " + err.Error())
	return -1, err
}

/**
post date list
 */
func (postDb *PostDb) QueryUserPostByDateRange(uid int, startDate int, endDate int, orderByDesc bool, limit int) (postList []Posts, err error) {

	o := postDb.DbConnect.Orm
	o.Using(postDb.DbConnect.DbName)
	//var queryGet string
	var orderby string
	if orderByDesc {
		orderby = "-create_date"
	} else {
		orderby = "create_date"
	}
	var maps []orm.Params
	var allResPosts []Posts
	qs := o.QueryTable(postDb.tableName).Filter("uid", uid).Filter("create_date__gte", startDate).Filter("create_date__lte", endDate).OrderBy(orderby).Limit(limit)
	_, err = qs.Values(&maps)
	//queryGet = "select * from " + postDb.tableName + " where uid = ? and create_date > ? and create_date < ? order by create_date " + orderby + " limit ?";
	//searchVal  := []int{uid, startDate, endDate, limit}
	//_, err = o.Raw(queryGet, searchVal).Values(&maps)

	if err == nil {
		for _, posts := range maps {
			//logs.Warning(posts)
			eachPost, err := _assignMapToPost(posts)
			if err != nil {
				logs.Warning(err)
			} else {
				allResPosts = append(allResPosts, eachPost)
			}

		}

		return allResPosts, nil
	}
	return allResPosts, err
}

func (postDb *PostDb) QueryUserPostByStoryId(storyId int64, uid int, limit int, page int, orderByDesc bool, countChannel chan int64) (postList []Posts, err error) {

	o := postDb.DbConnect.Orm
	o.Using(postDb.DbConnect.DbName)
	//var queryGet string
	var orderby string
	if orderByDesc {
		orderby = "-create_date"
	} else {
		orderby = "create_date"
	}
	offset := page * limit
	var maps []orm.Params
	var allResPosts []Posts

	qs := o.QueryTable(postDb.tableName).Filter("uid", uid).Filter("story_id", storyId).OrderBy(orderby).Limit(limit, offset)

	if countChannel != nil{
		go func(intChannel chan int64) {
			defer close(intChannel)
			counter, errCount := o.QueryTable(postDb.tableName).Filter("uid", uid).Filter("story_id", storyId).OrderBy(orderby).Count()
			if errCount == nil {
				countChannel <- counter
			}
			countChannel<-0
		}(countChannel)
	}

	_, err = qs.Values(&maps)
	if err == nil {
		for _, posts := range maps {
			//logs.Warning(posts)
			eachPost, err := _assignMapToPost(posts)
			if err != nil {
				logs.Warning(err)
			} else {
				allResPosts = append(allResPosts, eachPost)
			}
		}
		return allResPosts, nil
	}
	return allResPosts, err
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
			if err != nil {
				logs.Warning(err)
			} else {
				allResPosts = append(allResPosts, eachPost)
			}

		}
		return allResPosts, nil
	}

	return allResPosts, err
}

func _assignMapToPost(fromMap orm.Params) (eachPost Posts, err error) {
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
	mypost.Story_id = NewPost.Story_id
	mypost.Header = NewPost.Header
	mypost.Rel = NewPost.Rel
	mypost.Content = NewPost.Content
	mypost.Update_time = NewPost.Update_time
	mypost.Create_date = NewPost.Create_date
	res, err := o.Insert(mypost)

	if err != nil {
		jsonData, _ := json.Marshal(mypost)
		logs.Warning("post fail %s error: %s ", jsonData, err.Error())
	} else {
		postId = res
	}

	return postId, err
}


/**

 */
func (postDb *PostDb)ClearPostOut(postList []Posts) (allPubPost []map[string]interface{}) {

	for _, eachPost := range postList {
		mapVal := library.Struct2Map(eachPost)
		delete(mapVal, "Uid")
		allPubPost = append(allPubPost, mapVal)
	}
	return allPubPost
}

