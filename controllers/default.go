package controllers

import (
	"github.com/astaxie/beego"
	"onestory/library"
	"time"
	"encoding/base64"
	"onestory/services/request/third"
	"onestory/services/request/lbs"
	"io/ioutil"
	"onestory/models"
	"github.com/astaxie/beego/logs"
)

type LogedInUserController struct {
	beego.Controller
	user models.UserCache
}

func (c *LogedInUserController)requireUserLogIn() {
	defer func() {
		if err := recover(); err != nil {
			logs.Warning(err)
			output, _ := library.ReturnJsonWithError(library.CodeErrApi, "11111", "")
			c.Ctx.WriteString(output);
		}
	}()

	cookiekey := beego.AppConfig.String("passid")

	//get from cache
	passId, resBool := c.GetSecureCookie(cookiekey, "passid")
	if !resBool {
		output, _ := library.ReturnJsonWithError(1, "用户未登录", "")
		c.Ctx.WriteString(output)
		c.StopRun()
	}

	cahchedUser, getUserErr := models.GetUserFromCache(passId, false)
	if getUserErr != nil {
		output, _ := library.ReturnJsonWithError(1, "获取用户信息失败", "")
		c.Ctx.WriteString(output)
		c.StopRun()
	}
	c.user = cahchedUser
}


type (

	TestController struct {
		beego.Controller
	}
	UploadController struct {
		beego.Controller
	}
	WeatherController struct {
		beego.Controller
	}
)

func (c *UploadController) Post() {
	c.EnableXSRF = false
	fileType := c.GetString("type", "file")
	var path string
	nowTimging := time.Now().Format("2006-01-02-03-04-05")
	var output string

	if fileType == "base64" {
		dataSource := c.GetString("myfile")
		fileData, _ := base64.StdEncoding.DecodeString(dataSource)
		path = "./temp/" + nowTimging + library.RandSeq(3) + "_.png";
		errWriteFile := ioutil.WriteFile(path, fileData, 0777)
		if errWriteFile != nil {
			output, _ = library.ReturnJsonWithError(library.CodeErrCommen, "upload pic fail", "")
			c.Ctx.WriteString(output)
			return
		}
	}else{
		f, h, _ := c.GetFile("myfile")
		path = "./temp/" + nowTimging + h.Filename;
		defer f.Close()
		c.SaveToFile("myfile", path)
	}

	qiuniuApi := third.NewQiNiu(true)
	imgKey, errUp := qiuniuApi.Upoloader(path)
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

func (c *TestController) Get() {
	c.EnableXSRF = false

	email := "e930300047@163.com"
	subject := "激活账户通知"
	openUrl :=  "12345"
	message := "<html><body><a href='"+openUrl+"'>注册成功，请点击链接激活账户<a> <br> 或复制以下链接至浏览器 " + openUrl +" </body></html>"
	errEmail := third.SendToMail(email, subject, message, "html")
	//defer func(){ // 必须要先声明defer，否则不能捕获到panic异常
	//	logs.Warning("begin defer")
	//
	//	if err:=recover();err!=nil{
	//		stringRes, _ := library.ReturnJsonWithError(1,"获取信息失败", "")
	//		c.Ctx.WriteString(stringRes)
	//		return
	//		logs.Warning(err)
	//	}
	//	logs.Warning("ends defer")
	//}()
	//res := third.SingleSendMail()
	//errMail := library.SendToMail("onestory90@163.com", "test", "hahahaha", "")
	//logs.Warning(errMail)
	stringRes, _ := library.ReturnJsonWithError(1, "", errEmail)
	c.Ctx.WriteString(stringRes)
	return

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
	//var city string
	//city = c.GetString("city")
	//if len(city) < 1 {
	//	res := library.GetClientIp(c.Ctx.Request)
	//	city, _ = lbs.GetLocationByIp(res)
	//}
	//if len(city) < 1 {
	//	stringRes, _ := library.ReturnJsonWithError(1,"获取信息失败", "")
	//	c.Ctx.WriteString(stringRes)
	//	return
	//}
	//
	//mapRes, err := lbs.GetWeatherByLocation(city)
	//if err != nil{
	//	stringRes, _ := library.ReturnJsonWithError(1,err.Error(), "")
	//	c.Ctx.WriteString(stringRes)
	//}else{
	//	stringRes, _ := library.ReturnJsonWithError(0,"", mapRes)
	//	c.Ctx.WriteString(stringRes)
	//}

	return
	//conn := rediscli.RedisClient.Get()
	//_, err2 := conn.Do("SET", "hello", "world")
	//
	//if err2!=nil{
	//	panic(err2)
	//}
	//
	//defer conn.Close()
	//mapVal, ok := mapRes.(map[string]interface{})
	//if !ok {
	//	fmt.Print("get wether fail")
	//}
	//fmt.Print(mapVal["realtime"])
	//fmt.Print("11111111111111")
	//realtimeVal , ok := mapVal["realtime"].(map[string]interface{})
	//if !ok {
	//	fmt.Print("get wether fail")
	//}
	//
	//

	//stringRes, _ := library.ReturnJsonWithError(1,"", realtimeVal)
	//c.Ctx.WriteString(stringRes)
}

func (c *WeatherController) Get() {

	var city string
	city = c.GetString("city")
	if len(city) < 1 {
		res := library.GetClientIp(c.Ctx.Request)
		city, _ = lbs.GetLocationByIp(res)
	}
	if len(city) < 1 {
		stringRes, _ := library.ReturnJsonWithError(1,"获取信息失败", "")
		c.Ctx.WriteString(stringRes)
		c.StopRun()
		return
	}

	mapRes, err := lbs.GetWeatherByLocation(city)
	if err != nil{
		stringRes, _ := library.ReturnJsonWithError(1,err.Error(), "")
		c.Ctx.WriteString(stringRes)
		c.StopRun();
	}else{
		stringRes, _ := library.ReturnJsonWithError(0,"", mapRes)
		c.Ctx.WriteString(stringRes)
		c.StopRun();
	}
	return
}