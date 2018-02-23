package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
	"onestory/services"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"time"
)

type (
	StoryDb struct {
		tableName string
		DbConnect *services.DbService
	}

	OpenData struct{
		Story_id    int64
		Cover       string
		Title       string
		Desc        string
		Pen_name    string
		Create_time int64
		Update_time int64
		Extend      string
	}

	StoryInfo struct {
		Id          int64 `orm:"auto"`
		Uid         int
		Is_open     int
		OpenData
	}

	StoryAllFormat struct {
		Id          int64
		Uid         int
		Is_open     int
		Story_id    int64
		Cover       string
		Desc        string
		Title       string
		Pen_name    string
		Create_time int64
		Update_time int64
		Extend      string
	}
)

func NewStoryInfo() (*StoryDb) {

	dbService, err := services.NewService("onestory")
	if err != nil {
		logs.Warn(err)
	}
	return &StoryDb{"story_info", dbService}
}

/**
get all user posts
 */
func (storyDb *StoryDb) GetUserStoryInfoByUid(uid int, page int, limit int, counter chan int64) (allStorys []StoryInfo, err error) {

	o := storyDb.DbConnect.Orm
	o.Using(storyDb.DbConnect.DbName)
	var maps []orm.Params

	if limit <= 0 {
		limit = 20 //max 10
	}
	offset := page * limit
	qs := o.QueryTable(storyDb.tableName).Filter("uid", uid).OrderBy("-update_time").Limit(limit, offset)
	_, err = qs.Values(&maps)

	if counter != nil{
		go func() {
			defer close(counter)
			countRow, errCounter := o.QueryTable(storyDb.tableName).Filter("uid", uid).Count()
			if errCounter != nil {
				countRow = 0
			}
			counter <- countRow
		}()
	}


	var allResStory []StoryInfo
	if err == nil {
		for _, story := range maps {
			eachStory, err := _assignMapToStoryInfo(story)
			if err != nil {
				logs.Warning(err)
			} else {
				allResStory = append(allResStory, eachStory)
			}
		}
		return allResStory, nil
	}
	return allResStory, err
}

func _assignMapToStoryInfo(fromMap orm.Params) (eachStory StoryInfo, err error) {

	var storyAllFormat *StoryAllFormat
	err = mapstructure.Decode(fromMap, &storyAllFormat)
	if err != nil {
		return eachStory, err
		panic(err)
	}
	eachStory.Id = storyAllFormat.Id
	eachStory.Story_id = storyAllFormat.Story_id
	eachStory.Create_time = storyAllFormat.Create_time
	eachStory.Update_time = storyAllFormat.Update_time
	eachStory.Uid = storyAllFormat.Uid
	eachStory.Is_open = storyAllFormat.Is_open
	eachStory.Cover = storyAllFormat.Cover
	eachStory.Desc = storyAllFormat.Desc
	eachStory.Title = storyAllFormat.Title
	eachStory.Pen_name = storyAllFormat.Pen_name
	eachStory.Extend = storyAllFormat.Extend
	return eachStory, nil
}

/**
add a new story
 */
func (storyDb *StoryDb) AddNewUserStory(storyInfo StoryInfo) (storyId int64, err error) {
	o := storyDb.DbConnect.Orm
	o.Using(storyDb.DbConnect.DbName)

	storyId = -1
	mystory := new(StoryInfo)
	mystory.Uid = storyInfo.Uid
	mystory.Story_id = storyInfo.Story_id
	mystory.Cover = storyInfo.Cover
	mystory.Desc = storyInfo.Desc
	mystory.Title = storyInfo.Title
	mystory.Pen_name = storyInfo.Pen_name
	mystory.Is_open = storyInfo.Is_open
	mystory.Update_time = storyInfo.Update_time
	mystory.Create_time = storyInfo.Create_time
	res, err := o.Insert(mystory)

	if err != nil {
		jsonData, _ := json.Marshal(mystory)
		logs.Warning("post fail %s error: %s ", jsonData, err.Error())
	} else {
		storyId = res
	}

	return storyId, err
}

func (storyDb *StoryDb) UpdateUserStory(storyInfo StoryInfo) (*StoryInfo, error) {
	o := storyDb.DbConnect.Orm
	o.Using(storyDb.DbConnect.DbName)

	requireUpdate := false
	var story StoryInfo
	story.Story_id = storyInfo.Story_id
	story.Uid = storyInfo.Uid

	errRead := o.QueryTable(storyDb.tableName).Filter("Story_id", storyInfo.Story_id). Filter("Uid", storyInfo.Uid).One(&story)

	if errRead == nil {
		//update fields
		story.Update_time = time.Now().Unix()
		if len(storyInfo.Title) > 0 {
			requireUpdate = true
			story.Title = storyInfo.Title
		}
		if len(storyInfo.Desc) > 0 {
			requireUpdate = true
			story.Desc = storyInfo.Desc
		}
		if len(storyInfo.Cover) > 0 {
			requireUpdate = true
			story.Cover = storyInfo.Cover
		}
		if len(storyInfo.Pen_name) > 0 {
			requireUpdate = true
			story.Pen_name = storyInfo.Pen_name
		}
		if len(storyInfo.Extend) > 0 {
			requireUpdate = true
			story.Extend = storyInfo.Extend
		}
		if storyInfo.Is_open != 0 {
			requireUpdate = true
			story.Is_open = storyInfo.Is_open
		}
		if !requireUpdate {
			return nil, nil
		}
		if num, err := o.Update(&story); err == nil && num == 1{
			logs.Trace(string(story.Id) + " update user story succ " + string(num))
			return &story, nil
		}else{
			logs.Warning(string(story.Id) + " update user story fail " + err.Error())
			return nil, err
		}

	}else{
		if errRead == orm.ErrMultiRows {
			// 多条的时候报错
			return nil, fmt.Errorf("记录错误")
		}
		if errRead == orm.ErrNoRows {
			// 没有找到记录
			return nil, fmt.Errorf("没有查找到记录")
		}

		return nil, fmt.Errorf(" 更新错误")
	}
	return nil, fmt.Errorf("更新错误")
}


func (storyDb *StoryDb) GetUserStory(storyId int64, uid int) (storyInfo StoryInfo, err error) {
	o := storyDb.DbConnect.Orm
	o.Using(storyDb.DbConnect.DbName)

	var maps []orm.Params
	qs := o.QueryTable(storyDb.tableName).Filter("uid", uid).Filter("story_id", storyId).Limit(1)
	_, err = qs.Values(&maps)
	if len(maps) != 1 {
		return storyInfo, fmt.Errorf("no row found")
	}
	if err == nil {
		for _, getData := range maps{
			getStory, _ := _assignMapToStoryInfo(getData)
			return getStory, nil
		}
	}
	return storyInfo, err
}
