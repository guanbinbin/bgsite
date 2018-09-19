package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

func SetIsLogin(c *BaseController) {
	if c.GetSession("auth") == nil {
		c.Data["is_login"] = false
	} else {
		c.Data["is_login"] = true
	}
}

func (c *BaseController) Logout(){
	c.DestroySession()
	c.Redirect("/",302)
}