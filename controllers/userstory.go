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

	UpdateUserStoryController struct {
		LogedInUserController
	}
)

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

	cover := c.GetString("cover", "")
	title := c.GetString("title", "无题")
	pen_name := c.GetString("pen_name", user.Nick_name)
	library.CreateRandId(8)
	nowTime := time.Now().Unix()

	var storyData models.StoryInfo

	storyData.Id = 0
	storyData.Uid = user.Id
	storyData.Story_id = library.CreateRandId(1)
	storyData.Cover = cover
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
