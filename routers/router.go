package routers

import (
	"bgsite/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.PagesController{}, "get,post:Index")
	beego.Router("/map", &controllers.PagesController{}, "get:Map")
	beego.Router("/latest", &controllers.PagesController{}, "get:Latest")
	beego.Router("/thanks", &controllers.PagesController{}, "get,post:Thanks")
	beego.Router("/catalog/:id([0-9]+)", &controllers.PagesController{}, "get:Catalog")
	beego.Router("/product/:id([0-9]+)", &controllers.PagesController{}, "get:Product")
	beego.Router("/test", &controllers.PagesController{}, "get:Test")
	beego.Router("/auth", &controllers.LoginController{}, "get,post:Auth")
	beego.Router("/signup", &controllers.LoginController{}, "post:SignUp")
	beego.Router("/signin", &controllers.LoginController{}, "post:SignIn")
	beego.Router("/home", &controllers.LoginController{}, "get,post:Home")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
	beego.Router("/cart", &controllers.CartController{}, "get:Cart")
	beego.Router("/addtocart", &controllers.CartController{}, "post:AddToCart")
	beego.Router("/rmfromcart", &controllers.CartController{}, "post:RmFromCart")
	beego.Router("/chquantity", &controllers.CartController{}, "post:ChQuantity")
	beego.Router("/order", &controllers.CartController{}, "get:Order")
	beego.Router("/makeorder", &controllers.CartController{}, "post:MakeOrder")
	beego.Router("/orderinfo/:id([0-9]+)", &controllers.CartController{}, "get,post:OrderInfo")
}