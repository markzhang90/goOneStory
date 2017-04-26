package routers

import (
	"onestory/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hello/?:id/?:test", &controllers.HelloController{})
	beego.Router("/user/adduserprofile", &controllers.AddUserProfileController{})
	beego.Router("/user/updateuserprofile", &controllers.UpdateUserProfileController{})
	beego.Router("/user/getuserprofile", &controllers.GetUserProfileController{})
	beego.Router("/user/logintosys", &controllers.LoginUserController{})
	beego.Router("/post/adduserpost", &controllers.AddUserPostController{})
	beego.Router("/post/getuserpost", &controllers.GetUserPostController{})
}
