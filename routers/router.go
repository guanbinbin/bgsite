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
	beego.Router("/latest", &controllers.PagesController{}, "get:Latest")
	beego.Router("/catalog/:id([0-9]+)", &controllers.PagesController{}, "get:Catalog")
	beego.Router("/product/:id([0-9]+)", &controllers.PagesController{}, "get:Product")
	beego.Router("/test", &controllers.PagesController{}, "get:Test")
	beego.Router("/filters", &controllers.PagesController{}, "post:Filters")
	beego.Router("/addtocart", &controllers.PagesController{}, "post:AddToCart")
}