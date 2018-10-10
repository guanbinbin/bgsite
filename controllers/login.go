package controllers

import (
	"bgsite/models"
)

type LoginController struct {
	BaseController
}

func (c* LoginController) Auth () {
	//Button "Login"
	SetHeader(&c.BaseController)
	if c.Ctx.Request.FormValue("submit") == "signin" {
		c.Redirect("/signin",307) //307 - using POST
	}

	//Button "Register"
	if c.Ctx.Request.FormValue("submit") == "signup"{
		c.Redirect("/signup",307)
	}

	c.Layout = "layout.html"
	c.TplName = "login/auth.html"
}

func (c* LoginController) SignUp () {
	SetHeader(&c.BaseController)

	//Check if login exists, return error OR register & return name
	name, err  := models.CheckRegistration(c.GetString("name"), c.GetString("pass"))
	if err == "" {
		c.Data["msg"] = "Вы зарегистрировались как  " + name + ". Пожалуйста, авторизуйтесь."
	} else {
		c.Data["msg"] = err
	}

	c.Layout = "layout.html"
	c.TplName = "login/signup.html"
}

func (c* LoginController) SignIn () {
	SetHeader(&c.BaseController)

	//Check name, pass, return id if ok OR return error
	id, err  := models.CheckLogin(c.GetString("name"), c.GetString("pass"))
	if err == "" {
		c.SetSession("auth", id)     //Id to session
		if models.CheckIfAdmin(id) == true {
			c.Redirect("/admin",302)
		} else {
			c.Redirect("/home",302)
		}

	} else {
		c.Data["msg"] = err
	}

	c.Layout = "layout.html"
	c.TplName = "login/signin.html"
}

func (c *BaseController) Logout(){
	c.DelSession("auth")
	c.Redirect("/",302)
}
