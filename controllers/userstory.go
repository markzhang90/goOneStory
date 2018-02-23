package controllers

import (
	"onestory/models"
	"onestory/library"
	"time"
)

type (
	AddUserStoryController struct {
		LogedInUserController
	}

	GetStoryPostController struct {
		LogedInUserController
	}

	GetStoryListController struct {
		LogedInUserController
	}

	UpdateUserStoryController struct {
		LogedInUserController
	}
)

func (c *GetStoryListController) Get() {
	defer func() {
		errRecover := recover()
		output, alertErr := library.PanicFunc(errRecover)
		if alertErr != nil{
			c.Ctx.WriteString(output)
		}
	}()
	page, _ := c.GetInt("page", 1)
	limit, _ := c.GetInt("limit", 5)

	c.EnableXSRF = false
	c.requireUserLogIn()
	var getRes = make(map[string]interface{})
	var newStoryDb = models.NewStoryInfo()
	chCounter := make(chan int64)
	stroyList, errList := newStoryDb.GetUserStoryInfoByUid(c.user.Id, page-1, limit, chCounter)
	if errList != nil {
		panic("获取story信息失败")
	}
	count := <-chCounter

	getRes["page"] = page
	getRes["limit"] = limit
	getRes["story_list"] = stroyList
	getRes["count"] = count
	outputSucc, _ := library.ReturnJsonWithError(library.CodeSucc, "succ", getRes)
	c.Ctx.WriteString(outputSucc)
}

func (c *GetStoryPostController) Get() {
	defer func() {
		errRecover := recover()
		output, alertErr := library.PanicFunc(errRecover)
		if alertErr != nil{
			c.Ctx.WriteString(output)
		}
	}()

	c.EnableXSRF = false
	c.requireUserLogIn()
	var getRes = make(map[string]interface{})

	user := c.user

	storyId, _ := c.GetInt64("story_id", 0)
	page, _ := c.GetInt("page", 1)
	needStoryInfo, _ := c.GetInt("get_story", 1)
	limit, _ := c.GetInt("limit", 5)
	getRes["limit"] = limit
	getRes["page"] = page

	if storyId == 0 {
		panic("story参数无效")
	}

	var newStoryDb = models.NewStoryInfo()
	var newPostDb = models.NewPost()

	var chPosts = make(chan *[]models.Posts)
	var counter int64
	go func(myStoryId int64, uid int) {
		defer close(chPosts)
		var chCounter = make(chan int64)
		getPostList, err := newPostDb.QueryUserPostByStoryId(myStoryId, uid, limit, page-1, true, chCounter)
		if err != nil {
			chPosts <- nil
		}
		counter = <- chCounter
		chPosts <- &getPostList
	}(storyId, user.Id)

	if needStoryInfo == 1 {
		var chStoryInfo = make(chan *models.OpenData)
		go func(userStoryId int64, uid int) {
			defer close(chStoryInfo)
			getRes, err := newStoryDb.GetUserStory(userStoryId, uid)
			if err != nil {
				chStoryInfo <- nil
			}
			chStoryInfo <- &getRes.OpenData
		}(storyId, user.Id)
		storyInfo := <-chStoryInfo
		if storyInfo == nil {
			panic("获取story信息失败")
		}
		getRes["story_info"] = &storyInfo
	}

	postInfo := <-chPosts
	if postInfo == nil {
		panic("获取记录失败")
	}

	getRes["post_info"] = postInfo
	getRes["count"] = counter
	outputSucc, _ := library.ReturnJsonWithError(library.CodeSucc, "succ", getRes)
	c.Ctx.WriteString(outputSucc)
}

//新增story
func (c *AddUserStoryController) Get() {
	defer func() {
		errRecover := recover()
		output, alertErr := library.PanicFunc(errRecover)
		if alertErr != nil{
			c.Ctx.WriteString(output)
		}
	}()

	c.EnableXSRF = false
	c.requireUserLogIn()
	user := c.user

	cover := c.GetString("cover", "cover")
	desc := c.GetString("desc", "desc")
	title := c.GetString("title", "无题")
	pen_name := c.GetString("pen_name", user.Nick_name)
	library.CreateRandId(8)
	nowTime := time.Now().Unix()

	var storyData models.StoryInfo

	storyData.Id = 0
	storyData.Uid = user.Id
	storyData.Story_id = library.CreateRandId(1)
	storyData.Cover = cover
	storyData.Desc = desc
	storyData.Title = title
	storyData.Pen_name = pen_name
	storyData.Create_time = nowTime
	storyData.Update_time = nowTime
	storyData.Extend = ""
	storyData.Is_open = 0

	var newUserDb = models.NewStoryInfo()
	insertId, err := newUserDb.AddNewUserStory(storyData)

	if err != nil {
		panic("创建失败")
	}
	storyData.Id = insertId
	outputSucc, _ := library.ReturnJsonWithError(library.CodeSucc, "succ", storyData.OpenData)
	c.Ctx.WriteString(outputSucc)
}

//新增story
func (c *UpdateUserStoryController) Get() {
	defer func() {
		errRecover := recover()
		output, alertErr := library.PanicFunc(errRecover)
		if alertErr != nil{
			c.Ctx.WriteString(output)
		}
	}()

	c.EnableXSRF = false
	c.requireUserLogIn()
	user := c.user

	cover := c.GetString("cover", "")
	storyId, _ := c.GetInt64("story_id", 0)
	isOpen, _ := c.GetInt("is_open", 0)
	title := c.GetString("title", "")
	desc := c.GetString("desc", "")
	penName := c.GetString("pen_name", "")
	extend := c.GetString("extend", "")
	library.CreateRandId(8)

	if storyId == 0 {
		panic("获取StoryId失败")
	}


	var storyData models.StoryInfo

	storyData.Uid = user.Id
	storyData.Story_id = storyId
	storyData.Cover = cover
	storyData.Title = title
	storyData.Desc = desc
	storyData.Pen_name = penName
	storyData.Extend = extend
	storyData.Is_open = isOpen

	var newUserDb = models.NewStoryInfo()
	updateCount, err := newUserDb.UpdateUserStory(storyData)

	if err != nil || updateCount == nil{
		panic("更新失败 " + err.Error())
	}
	outputSucc, _ := library.ReturnJsonWithError(library.CodeSucc, "succ", updateCount.OpenData)
	c.Ctx.WriteString(outputSucc)
}

