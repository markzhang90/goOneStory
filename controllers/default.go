package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	//"onestory/services/rediscli"
	"onestory/services/request"
	"html/template"
	"onestory/library"
	"fmt"
	"time"
	"onestory/models"
)

type (
	MainController struct {
		beego.Controller
	}
	EditController struct {
		beego.Controller
	}
	TestController struct {
		beego.Controller
	}
	UploadController struct {
		beego.Controller
	}
)


func (c *MainController) Get() {

	//get userinfo
	cookiekey := beego.AppConfig.String("passid")

	//get from cache
	var cahchedUser models.UserCache
	passId, resBool := c.GetSecureCookie(cookiekey, "passid")
	if resBool {
		cahchedUser, _ = models.GetUserFromCache(passId)
	}

	if len(cahchedUser.UserProfile.Nick_name)>1 {
		c.Data["nickname"] = cahchedUser.UserProfile.Nick_name
	}else {
		c.Data["nickname"] = ""
	}

	if len(cahchedUser.UserProfile.Avatar)>1 {
		c.Data["avatar"] = cahchedUser.UserProfile.Avatar
	}else {
		c.Data["avatar"] = ""
	}

	c.Data["xsrfdata"]= template.HTML(c.XSRFFormHTML())
	c.Layout = "onestory/base.html"
	c.TplName = "onestory/feed.html"

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Fixheader"] = "onestory/fixheader.html"
	c.LayoutSections["Footer"] = "onestory/footer.html"
}

func (c *UploadController) Post() {
	//c.EnableXSRF = false
	f, h, _ := c.GetFile("myfile")
	nowTimging := time.Now().Format("2006-01-02-03-04-05")
	path := "./temp/" + nowTimging + h.Filename;
	defer f.Close()
	c.SaveToFile("myfile", path)
	qiuniuApi := library.NewQiNiu(false)
	imgKey, errUp := qiuniuApi.Upoloader(path)
	var output string
	if errUp == nil{
		var outRes = make(map[string]string)
		outRes["key"] = imgKey
		url := qiuniuApi.DownloadUrl(imgKey)
		outRes["url"] = url
		output, _ = library.ReturnJsonWithError(library.CodeSucc, library.CodeString(library.CodeSucc), outRes)
	}else {
		output, _ = library.ReturnJsonWithError(library.CodeErrCommen, "upload pic fail", "")
	}
	c.Ctx.WriteString(output)
}

func (c *EditController) Get() {
	c.Data["xsrfdata"]= template.HTML(c.XSRFFormHTML())
	c.Layout = "onestory/base.html"
	c.TplName = "onestory/edit.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Fixheader"] = "onestory/fixheader.html"
	c.LayoutSections["Footer"] = "onestory/footer.html"
}


func (c *TestController) Get() {

	defer func(){ // 必须要先声明defer，否则不能捕获到panic异常
		logs.Warning("begin defer")

		if err:=recover();err!=nil{
			logs.Warning(err)
		}
		logs.Warning("ends defer")
	}()

	//varId := c.Ctx.Input.Param(":id")
	//varTest := c.Ctx.Input.Param(":test")
	//
	//v := c.GetSession("asta")
	//if v == nil {
	//	c.SetSession("asta", int(1))
	//	c.Data["num"] = 0
	//} else {
	//	c.SetSession("asta", v.(int)+1)
	//	c.Data["num"] = v.(int)
	//}
	//logs.Warning("coooool")
	city := c.GetString("city")
	mapRes, err := request.GetWeatherInfo(city)

	if err != nil{
		c.Ctx.WriteString(err.Error())
	}

	//conn := rediscli.RedisClient.Get()
	//_, err2 := conn.Do("SET", "hello", "world")
	//
	//if err2!=nil{
	//	panic(err2)
	//}
	//
	//defer conn.Close()
	mapVal, ok := mapRes.(map[string]interface{})
	if !ok {
		fmt.Print("get wether fail")
	}

	realtimeVal , ok := mapVal["realtime"].(map[string]interface{})
	if !ok {
		fmt.Print("get wether fail")
	}
	stringRes, _ := library.ReturnJsonWithError(0,"", realtimeVal["weather"])
	c.Ctx.WriteString(stringRes)
}

