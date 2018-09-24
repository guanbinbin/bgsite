package controllers

import (
	"bgsite/models"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

//Get categories for left nav, add them to template
func GetCategoriesForSideNav(c *BaseController){
	categories, _ := models.GetCategories() //return Id, name from db (slice)
	c.Data["categories"] = categories
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