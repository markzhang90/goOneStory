package routers

import (
	"onestory/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/test", &controllers.TestController{})
	beego.Router("/", &controllers.ShowController{})
	beego.Router("/wea", &controllers.WeatherController{})
	beego.Router("/main", &controllers.MainController{})
	beego.Router("/edit", &controllers.EditController{})
	beego.Router("/show", &controllers.ShowController{})
	beego.Router("/showdetail/?:id", &controllers.ShowDetailController{})
	beego.Router("/user/activeuser", &controllers.ActiveUserProfileController{})
	beego.Router("/user/register", &controllers.RegisterController{})
	beego.Router("/user/profile", &controllers.ProfileController{})
	beego.Router("/user/adduserprofile", &controllers.AddUserProfileController{})
	beego.Router("/user/updateuserprofile", &controllers.UpdateUserProfileController{})
	beego.Router("/user/getuserprofile", &controllers.GetUserProfileController{})
	beego.Router("/user/logintosys", &controllers.LoginUserController{})
	beego.Router("/uploader", &controllers.UploadController{})
	beego.Router("/post/adduserpost", &controllers.AddUserPostController{})
	beego.Router("/post/getuserpostbyid", &controllers.GetUserPostController{})
	beego.Router("/post/getuserpostdaterange", &controllers.GetUserPostDateRangeController{})
	beego.Router("/post/getuserpostdate", &controllers.GetUserPostDateController{})
	beego.Router("/post/getuserclosestpost", &controllers.GetUserPostClosestController{})

	beego.Router("api/wechat/logintosys", &controllers.LoginWehchatController{})
	beego.Router("api/wechat/initinfo", &controllers.InitWehchatController{})

	beego.Router("api/activeuser", &controllers.EmailConfirmController{})
	beego.Router("api/authcode", &controllers.SendEmailAuthController{})

}
