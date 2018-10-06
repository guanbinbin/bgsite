package routers

import (
	"bgsite/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.PagesController{}, "get,post:Index")
	beego.Router("/map", &controllers.PagesController{}, "get:Map")
	beego.Router("/thanks", &controllers.PagesController{}, "get,post:Thanks")
	beego.Router("/test", &controllers.PagesController{}, "get:Test")
	beego.Router("/auth", &controllers.LoginController{}, "get,post:Auth")
	beego.Router("/signup", &controllers.LoginController{}, "post:SignUp")
	beego.Router("/signin", &controllers.LoginController{}, "post:SignIn")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
	beego.Router("/latest", &controllers.ProductsController{}, "get:Latest")
	beego.Router("/catalog/:id:int", &controllers.ProductsController{}, "get:Catalog")
	beego.Router("/product/:id:int", &controllers.ProductsController{}, "get:Product")
	beego.Router("/cart", &controllers.CartController{}, "get:CartAndOrder")
	beego.Router("/addtocart", &controllers.CartController{}, "post:AddToCart")
	beego.Router("/rm/:id:int", &controllers.CartController{}, "get:RmFromCart")
	beego.Router("/qnt/:mark([0-1])/:pr:int", &controllers.CartController{}, "get:ChQuantity")
	beego.Router("/order", &controllers.CartController{}, "get:CartAndOrder")
	beego.Router("/makeorder", &controllers.CartController{}, "post:MakeOrder")
	beego.Router("/home", &controllers.UserController{}, "get,post:Home")
	beego.Router("/orderinfo/:id:int", &controllers.UserController{}, "get,post:OrderInfo")
	beego.Router("/admin/home", &controllers.AdminController{}, "get:AdmIndex")
	beego.Router("/admin/login", &controllers.AdminController{}, "get:Login")
	beego.Router("/admin/check", &controllers.AdminController{}, "post:CheckLogin")
}