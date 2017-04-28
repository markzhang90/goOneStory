package controllers

import (
	"github.com/astaxie/beego"
	"onestory/models"
	"onestory/library"
	"github.com/astaxie/beego/logs"
	"time"
	"strconv"
)

type (
	AddUserPostController struct {
		beego.Controller
	}

	GetUserPostController struct {
		beego.Controller
	}
)

func (c *AddUserPostController) Get() {

	title := "this is title"
	content := "this is content"
	timeNow := time.Now().Unix()
	timeFormat := time.Unix(timeNow, 0).Format("20060102")
	timeFormatInt, err := strconv.ParseInt(timeFormat, 10, 64)
	if err != nil{
		logs.Warn("failed convert", err)
	}
	postData := models.Posts{
		Uid: 16,
		Title: title,
		Content: content,
		Update_time: time.Now().Unix(),
		Create_date: timeFormatInt,
	}
	c.EnableXSRF = false
	var newPostDb = models.NewPost()
	//var getUser = newUser.GetUserProfile()
	//logs.Warning(getUser)
	res, err := newPostDb.AddNewUserPost(postData)
	var output string

	if err != nil{
		output, _ = library.ReturnJsonWithError(library.AddPostFail, err.Error(), nil)

	}else {
		output, _ = library.ReturnJsonWithError(library.AddPostFail, "ref", res)
	}
	c.Ctx.WriteString(output)
}


func (c *GetUserPostController) Get() {
	uid := 15
	limit, err := c.GetInt("num")
	if err != nil{
		limit = 1
	}

	c.EnableXSRF = false
	var newPostDb = models.NewPost()
	//var getUser = newUser.GetUserProfile()
	//logs.Warning(getUser)
	postList, err := newPostDb.GetUserAllRecentPosts(uid, limit)

	var output string

	if err != nil{
		output, _ = library.ReturnJsonWithError(library.CodeErrCommen, err.Error(), nil)

	}else {
		output, _ = library.ReturnJsonWithError(library.CodeErrCommen, "ref", postList)
	}

	c.Ctx.WriteString(output)
}


