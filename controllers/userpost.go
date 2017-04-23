package controllers

import (
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/logs"
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
	logs.Warn(timeFormatInt)
	postData := models.Posts{
		Uid: 16,
		Title: title,
		Content: content,
		Update_time: time.Now().Unix(),
		Create_date: timeFormatInt,
	}
	logs.Warn(postData)
	c.EnableXSRF = false
	var newPostDb = models.NewPost()
	//var getUser = newUser.GetUserProfile()
	//logs.Warning(getUser)
	newPostDb.AddNewUserPost(postData)
	//redirect
	c.Abort("500")
	//c.Ctx.Redirect(302, "/hello/123/5")
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


