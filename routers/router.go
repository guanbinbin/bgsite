package routers

import (
	"bgsite/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.PagesController{}, "get,post:Index")
	beego.Router("/nameold", &controllers.PagesController{}, "get:NameOld")
	beego.Router("/hello", &controllers.PagesController{}, "post:Hello")
	beego.Router("/answer", &controllers.PagesController{}, "get,post:Answer")
	beego.Router("/calc", &controllers.PagesController{}, "get,post:Calc")
	beego.Router("/session", &controllers.LoginController{}, "get,post:Session")
	beego.Router("/register", &controllers.LoginController{}, "post:Register")
	beego.Router("/login", &controllers.LoginController{}, "post:Login")
	beego.Router("/user", &controllers.LoginController{}, "get,post:User")
	beego.Router("/map", &controllers.PagesController{}, "get:Map")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
}