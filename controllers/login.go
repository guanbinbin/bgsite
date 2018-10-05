package controllers

import (
	"bgsite/models"
)

type LoginController struct {
	BaseController
}

func (c* LoginController) Auth () {
	//Button "Login"
	SetIsLogin(&c.BaseController)
	if c.Ctx.Request.FormValue("submit") == "signin" {
		c.Redirect("/signin",307) //307 - using POST
	}
	//Button "Register"
	if c.Ctx.Request.FormValue("submit") == "signup"{
		c.Redirect("/signup",307)
	}
	c.Layout = "layout.html"
	c.TplName = "auth.html"
}

func (c* LoginController) SignUp () {
	SetIsLogin(&c.BaseController)
	//Check if login exists, return error OR register & return name
	name, err  := models.CheckRegistration(c.GetString("name"), c.GetString("pass"))
	if err == "" {
		c.Data["msg"] = "Вы зарегистрировались как  " + name + ". Пожалуйста, авторизуйтесь."
	} else {
		c.Data["msg"] = err
	}
	c.Layout = "layout.html"
	c.TplName = "signup.html"
}

func (c* LoginController) SignIn () {
	SetIsLogin(&c.BaseController)
	//Check name, pass, return id if ok OR return error
	id, err  := models.CheckLogin(c.GetString("name"), c.GetString("pass"))
	if err == "" {
		c.SetSession("auth", id)     //Id to session
		c.Data["msg"] = "Вы авторизовались как  " + models.GetUserName(c.GetSession("auth"))
		c.Redirect("/home",302)
	} else {
		c.Data["msg"] = err
	}
	c.Layout = "layout.html"
	c.TplName = "signin.html"
}

func (c *LoginController) Home() {
	SetIsLogin(&c.BaseController)

	orders, _ := models.GetOrdersForUser(c.GetSession("auth").(int))

	c.Data["orders"] = orders
	c.Data["msg"] = models.GetUserName(c.GetSession("auth")) //Read name from db by user id
	c.Data["is_inlogin"] = true
	c.TplName = "home.html"
	c.Layout = "layout.html"
}
