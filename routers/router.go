package routers

import (
	"onestory/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/test", &controllers.TestController{})
	beego.Router("/main", &controllers.MainController{})
	beego.Router("/edit", &controllers.EditController{})
	beego.Router("/user/adduserprofile", &controllers.AddUserProfileController{})
	beego.Router("/user/updateuserprofile", &controllers.UpdateUserProfileController{})
	beego.Router("/user/getuserprofile", &controllers.GetUserProfileController{})
	beego.Router("/user/logintosys", &controllers.LoginUserController{})
	beego.Router("/post/adduserpost", &controllers.AddUserPostController{})
	beego.Router("/post/getuserpost", &controllers.GetUserPostController{})
	beego.Router("/post/getuserclosestpost", &controllers.GetUserPostClosestController{})
}
