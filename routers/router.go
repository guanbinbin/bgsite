package routers

import (
	"bgsite/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get,post:Index")
	beego.Router("/nameold", &controllers.MainController{}, "get:NameOld")
	beego.Router("/hello", &controllers.MainController{}, "post:Hello")
	beego.Router("/answer", &controllers.MainController{}, "get,post:Answer")
	beego.Router("/calc", &controllers.MainController{}, "get,post:Calc")
	beego.Router("/session", &controllers.LoginController{}, "get,post:Session")
	beego.Router("/register", &controllers.LoginController{}, "post:Register")
	beego.Router("/login", &controllers.LoginController{}, "post:Login")
}