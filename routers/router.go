package routers

import (
	"bgsite/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.PagesController{}, "get,post:Index")
	beego.Router("/auth", &controllers.LoginController{}, "get,post:Auth")
	beego.Router("/signup", &controllers.LoginController{}, "post:SignUp")
	beego.Router("/signin", &controllers.LoginController{}, "post:SignIn")
	beego.Router("/home", &controllers.LoginController{}, "get,post:Home")
	beego.Router("/map", &controllers.PagesController{}, "get:Map")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
	beego.Router("/catalog", &controllers.PagesController{}, "get:Catalog")
	beego.Router("/catalog/p=:page([0-9]+)", &controllers.PagesController{}, "get:Catalog")
	beego.Router("/category/:id([0-9]+)", &controllers.PagesController{}, "get:Category")
	beego.Router("/product/:id([0-9]+)", &controllers.PagesController{}, "get:Product")
	beego.Router("/category/:id([0-9]+)/p=:page([0-9]+)", &controllers.PagesController{}, "get:Category")
}