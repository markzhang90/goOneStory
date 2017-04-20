package routers

import (
	"onestory/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hello/?:id/?:test", &controllers.HelloController{})
	beego.Router("/adduserprofile", &controllers.UserProfileController{})
	beego.Router("/getuserprofile", &controllers.GetUserProfileController{})
	beego.Router("/post/adduserpost", &controllers.UserPostController{})
	beego.Router("/post/getuserpost", &controllers.GetUserPostController{})
}
