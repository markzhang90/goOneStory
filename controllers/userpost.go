package controllers

import (
	"github.com/astaxie/beego"
	"onestory/models"
	"onestory/library"
	"github.com/astaxie/beego/logs"
	"time"
	"strconv"
	"strings"
)

type (
	AddUserPostController struct {
		beego.Controller
	}

	GetUserPostController struct {
		beego.Controller
	}

	GetUserPostClosestController struct {
		beego.Controller
	}
)

func (c *AddUserPostController) Post() {

	cookiekey := beego.AppConfig.String("passid")

	//get from cache
	passId, _ := c.GetSecureCookie(cookiekey, "passid")
	logs.Warning(passId)
	if len(passId) <= 0 {
		output, _ := library.ReturnJsonWithError(library.GetUserFail, "ref", nil)
		c.Ctx.WriteString(output)
		return
	}
	cahchedUser, err := models.GetUserFromCache(passId)
	if err != nil {
		output, _ := library.ReturnJsonWithError(library.GetUserFail, "ref", err.Error())
		c.Ctx.WriteString(output)
		return
	}
	uid := cahchedUser.UserProfile.Id

	header := c.GetString("header", "无题")
	content := c.GetString("content", "")
	ref := c.GetString("ref", "")

	timeNow := time.Now().Unix()
	timeFormat := time.Unix(timeNow, 0).Format("20060102")
	timeFormatInt, err := strconv.ParseInt(timeFormat, 10, 64)
	if err != nil{
		logs.Warn("failed convert", err)
	}
	postData := models.Posts{
		Uid: uid,
		Header: header,
		Content: string(content),
		Rel: string(ref),
		Update_time: time.Now().Unix(),
		Create_date: timeFormatInt,
	}

	var newPostDb = models.NewPost()
	//var getUser = newUser.GetUserProfile()
	//logs.Warning(getUser)
	res, err := newPostDb.AddNewUserPost(postData)
	var output string

	if err != nil{
		output, _ = library.ReturnJsonWithError(library.AddPostFail, err.Error(), nil)

	}else {
		output, _ = library.ReturnJsonWithError(library.CodeSucc, "ref", res)
	}
	c.Ctx.WriteString(output)
}




func (c *GetUserPostController) Post() {
	cookiekey := beego.AppConfig.String("passid")


	//get from cache
	passId, _ := c.GetSecureCookie(cookiekey, "passid")
	logs.Warning(passId)
	if len(passId) <= 0 {
		output, _ := library.ReturnJsonWithError(library.GetUserFail, "ref", nil)
		c.Ctx.WriteString(output)
		return
	}
	cahchedUser, err := models.GetUserFromCache(passId)
	if err != nil {
		output, _ := library.ReturnJsonWithError(library.GetUserFail, "ref", err.Error())
		c.Ctx.WriteString(output)
		return
	}
	uid := cahchedUser.UserProfile.Id

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
		output, _ = library.ReturnJsonWithError(library.CodeSucc, "ref", postList)
	}

	c.Ctx.WriteString(output)
}


func (c *GetUserPostClosestController) Post() {
	cookiekey := beego.AppConfig.String("passid")


	//get from cache
	passId, _ := c.GetSecureCookie(cookiekey, "passid")
	logs.Warning(passId)
	if len(passId) <= 0 {
		output, _ := library.ReturnJsonWithError(library.GetUserFail, "ref", nil)
		c.Ctx.WriteString(output)
		return
	}
	cahchedUser, err := models.GetUserFromCache(passId)
	if err != nil {
		output, _ := library.ReturnJsonWithError(library.GetUserFail, "ref", err.Error())
		c.Ctx.WriteString(output)
		return
	}
	uid := cahchedUser.UserProfile.Id

	option := c.GetString("option")
	var isNext bool = false
	if option == "next"{
		isNext = true
	}
	date := c.GetString("date")
	dateCorrect := strings.Replace(date, "/", "", -1)
	intDate ,_ := strconv.Atoi(dateCorrect)
	logs.Warning(uid)
	c.EnableXSRF = false
	var newPostDb = models.NewPost()
	//var getUser = newUser.GetUserProfile()
	//logs.Warning(getUser)
	postList, err := newPostDb.GetUserClosestPost(uid, intDate, isNext)

	var output string

	if err != nil{
		output, _ = library.ReturnJsonWithError(library.CodeErrCommen, err.Error(), nil)

	}else {
		output, _ = library.ReturnJsonWithError(library.CodeSucc, "ref", postList)
	}

	c.Ctx.WriteString(output)
}

