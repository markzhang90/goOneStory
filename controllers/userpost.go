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

	GetUserPostDateController struct {
		beego.Controller
	}

	GetUserPostDateRangeController struct {
		beego.Controller
	}
)

func (c *AddUserPostController) Post() {

	cookiekey := beego.AppConfig.String("passid")

	//get from cache
	passId, _ := c.GetSecureCookie(cookiekey, "passid")

	if len(passId) <= 0 {
		passId = c.GetString("passid", "")
		if len(passId) < 1{
			output, _ := library.ReturnJsonWithError(library.GetUserFail, "ref", nil)
			c.Ctx.WriteString(output)
			return
		}
	}
	cahchedUser, err := models.GetUserFromCache(passId, true)
	if err != nil {
		output, _ := library.ReturnJsonWithError(library.GetUserFail, "ref", nil)
		c.Ctx.WriteString(output)
		return
	}
	uid := cahchedUser.UserProfile.Id

	header := c.GetString("header", "无题")
	content := c.GetString("content", "今日无事。。。")
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
	logs.Warning(postData);

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




func (c *GetUserPostController) Get() {
	c.EnableXSRF = false

	cookiekey := beego.AppConfig.String("passid")

	//get from cache
	passId, _ := c.GetSecureCookie(cookiekey, "passid")

	if len(passId) <= 0 {
		passId = c.GetString("passid", "")
		if len(passId) < 1{
			output, _ := library.ReturnJsonWithError(library.GetUserFail, "ref", nil)
			c.Ctx.WriteString(output)
			return
		}
	}
	cahchedUser, err := models.GetUserFromCache(passId, true)
	if err != nil {
		output, _ := library.ReturnJsonWithError(library.GetUserFail, "ref", err.Error())
		c.Ctx.WriteString(output)
		return
	}
	uid := cahchedUser.UserProfile.Id

	pageId, err := c.GetInt("id")
	if err != nil{
		output, _ := library.ReturnJsonWithError(library.ParamFail, "ref", nil)
		c.Ctx.WriteString(output)
		return
	}

	var newPostDb = models.NewPost()
	//var getUser = newUser.GetUserProfile()
	//logs.Warning(getUser)
	postList, err := newPostDb.GetPostByPassidAndId(uid, pageId)
	var output string

	if err != nil{
		output, _ = library.ReturnJsonWithError(library.CodeErrCommen, err.Error(), nil)

	}else {
		output, _ = library.ReturnJsonWithError(library.CodeSucc, "ref", postList)
	}

	c.Ctx.WriteString(output)
}

/**
get most recent post

@vars limit
@vars option : (next , previous)
@vars date : 20170101,2017/01/01
 */
func (c *GetUserPostClosestController) Post() {
	c.EnableXSRF = false

	cookiekey := beego.AppConfig.String("passid")

	//get from cache
	passId, _ := c.GetSecureCookie(cookiekey, "passid")
	if len(passId) <= 0 {
		passId = c.GetString("passid", "")
		if len(passId) < 1{
			output, _ := library.ReturnJsonWithError(library.GetUserFail, "ref", nil)
			c.Ctx.WriteString(output)
			return
		}
	}
	cahchedUser, err := models.GetUserFromCache(passId, true)
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

/**
get user post by date range

@vars limit
@vars order : (desc , asc)
@vars start : start date (20170101,2017/01/01)
@vars end : end date (20170101,2017/01/01)
 */
func (c *GetUserPostDateRangeController) Get() {
	c.EnableXSRF = false

	cookiekey := beego.AppConfig.String("passid")
	//get from cache
	passId, _ := c.GetSecureCookie(cookiekey, "passid")

	if len(passId) <= 0 {
		passId = c.GetString("passid", "")
		if len(passId) < 1{
			output, _ := library.ReturnJsonWithError(library.GetUserFail, "ref", nil)
			c.Ctx.WriteString(output)
			return
		}
	}
	cahchedUser, err := models.GetUserFromCache(passId, true)
	if err != nil {
		output, _ := library.ReturnJsonWithError(library.GetUserFail, "ref", err.Error())
		c.Ctx.WriteString(output)
		return
	}
	uid := cahchedUser.UserProfile.Id

	total, errTotal := c.GetInt("total")
	if errTotal != nil{
		total = 0
	}

	limit, errLimit := c.GetInt("limit")
	if errLimit != nil{
		limit = 10
	}

	isDesc := true
	order := c.GetString("order", "desc")
	if order != "desc"{
		isDesc = false
	}

	startDate := c.GetString("start", "0")
	dateFormatStart := strings.Replace(startDate, "/", "", -1)
	startDateInt ,_ := strconv.Atoi(dateFormatStart)

	endDate := c.GetString("end", "20201231")
	dateFormatEnd := strings.Replace(endDate, "/", "", -1)
	endDateInt ,_ := strconv.Atoi(dateFormatEnd)
	
	var newPostDb = models.NewPost()

	var allResult = make(map[string]interface{})
	//allResult["total"] = -1

	var output string
	postList, errList := newPostDb.QueryUserPostByDateRange(uid, startDateInt, endDateInt, isDesc, limit)
	if total == 1{
		allNum, errNum := newPostDb.QueryCountUserPostByDateRange(uid, startDateInt, endDateInt)
		if errNum != nil {
			allNum = -1
			logs.Warning("get count all query fail" , errNum)
			//output, _ = library.ReturnJsonWithError(library.CodeErrCommen, errNum.Error(), nil)
		}
		allResult["total"] = allNum
	}
	allResult["limit"] = limit
	allResult["order"] = order
	if errList != nil{
		allResult["list"] = make(map[string]interface{})
		output, _ = library.ReturnJsonWithError(library.CodeErrCommen, errList.Error(), nil)
	} else {
		psotMap := newPostDb.ClearPostOut(postList)
		allResult["list"] = psotMap
		output, _ = library.ReturnJsonWithError(library.CodeSucc, "ref", allResult)
	}

	c.Ctx.WriteString(output)
}

/**
get user post by dates

@vars limit auto
@vars order : (desc , asc)
@vars date : split by comma 20170101,2017/01/01
 */
func (c *GetUserPostDateController) Get() {
	c.EnableXSRF = false

	cookiekey := beego.AppConfig.String("passid")

	//get from cache
	passId, _ := c.GetSecureCookie(cookiekey, "passid")

	if len(passId) <= 0 {
		passId = c.GetString("passid", "")
		if len(passId) < 1{
			output, _ := library.ReturnJsonWithError(library.GetUserFail, "ref", nil)
			c.Ctx.WriteString(output)
			return
		}
	}
	cahchedUser, err := models.GetUserFromCache(passId, true)

	if err != nil {
		output, _ := library.ReturnJsonWithError(library.GetUserFail, "ref", err.Error())
		c.Ctx.WriteString(output)
		return
	}
	uid := cahchedUser.UserProfile.Id
	isDesc := true
	order := c.GetString("order", "desc")
	if order != "desc"{
		isDesc = false
	}

	dateStr := c.GetString("date", "0")
	dateMap := strings.Split(dateStr, ",")

	var queryDateList []int
	for _, eachDateStr := range dateMap{
		dateFormatStart := strings.Replace(eachDateStr, "/", "", -1)
		dateInt ,_ := strconv.Atoi(dateFormatStart)
		queryDateList = append(queryDateList, dateInt)
	}

	limit, err := c.GetInt("limit")
	if err != nil{
		limit = len(queryDateList)
		if limit < 1 {
			limit = 1
		}
	}

	var newPostDb = models.NewPost()

	postList, err := newPostDb.QueryUserPostByDate(uid, queryDateList, isDesc, limit)
	//for _, eachPost := range postList{
	//	eachPost
	//}
	var output string

	if err != nil{

		output, _ = library.ReturnJsonWithError(library.CodeErrCommen, err.Error(), nil)
	}else {
		var resultMap = make(map[string]interface{})
		psotMap := newPostDb.ClearPostOut(postList)
		for _, eachPost := range psotMap {
			date := strconv.FormatInt(eachPost["Create_date"].(int64),10)
			resultMap[date] = eachPost
		}
		output, _ = library.ReturnJsonWithError(library.CodeSucc, "ref", resultMap)
	}

	c.Ctx.WriteString(output)
}

