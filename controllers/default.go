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
	"onestory/services/request/third"
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
	c.Data["xsrfdata"]= template.HTML(c.XSRFFormHTML())
	c.Layout = "onestory/base.html"
	c.TplName = "onestory/feed.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Fixheader"] = "onestory/fixheader.html"
	c.LayoutSections["Footer"] = "onestory/footer.html"
}

func (c *UploadController) Post() {
	c.EnableXSRF = false
	f, h, _ := c.GetFile("myfile")
	nowTimging := time.Now().Format("2006-01-02-03-04-05")
	path := "./temp/" + nowTimging + h.Filename;
	defer f.Close()
	c.SaveToFile("myfile", path)
	qiuniuApi := third.NewQiNiu(true)
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
	fmt.Print(mapVal["realtime"])
	fmt.Print("11111111111111")
	realtimeVal , ok := mapVal["realtime"].(map[string]interface{})
	if !ok {
		fmt.Print("get wether fail")
	}


	stringRes, _ := library.ReturnJsonWithError(1,"", realtimeVal)
	c.Ctx.WriteString(stringRes)
}

