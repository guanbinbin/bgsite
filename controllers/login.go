package controllers

import (
	"bgsite/models"
	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	BaseController
}

func (c* LoginController) Session () {
	//Button "Login"
	SetIsLogin(&c.BaseController)
	if c.Ctx.Request.FormValue("submit") == "login" {
		c.Redirect("/login",307) //307 - using POST
	}
	//Button "Register"
	if c.Ctx.Request.FormValue("submit") == "register"{
		c.Redirect("/register",307)
	}
	c.Layout = "layout.html"
	c.TplName = "session.html"
}

func (c* LoginController) Register () {
	//TODO: remove orm to models
	SetIsLogin(&c.BaseController)
	//Use DB
	o := orm.NewOrm()
	o.Using("default")
	//Registration
	name, pass := c.GetString("name"), c.GetString("pass")
	checkUserName := &models.User{Name: name}
	if o.Read(checkUserName, "name") == nil { 	//Check if Username exists (if err=nil)
		c.Data["msg"] = "Choose another login"
	} else if  name == "" && pass == "" {
		c.Data["msg"] = "Enter correct login and pass"
	} else if  pass == "" {
		c.Data["msg"] = "Enter correct pass"
	} else if  name == "" {
		c.Data["msg"] = "Enter correct login"
	} else {
		addUser := models.User{Name: name, Pass: pass}
		o.Insert(&addUser) 							//add to DB
		c.Data["msg"] = "Вы зарегистрировались как  " + string(name) + ". Пожалуйста, авторизуйтесь."
		c.Data["login"] = true
	}
	c.Layout = "layout.html"
	c.TplName = "register.html"
}

func (c* LoginController) Login () {
	//TODO: remove orm to models
	SetIsLogin(&c.BaseController)
	//Use db
	o := orm.NewOrm()
	o.Using("default")
	//Login
	name, pass := c.GetString("name"), c.GetString("pass")
	check := &models.User{Name:name,Pass:pass}
	if o.Read(check,"name","pass") == nil { //Read from POST name, pass, if matching with db, err==nil
		c.SetSession("auth", check.Id)     //Id to session
		c.Data["msg"] = "Вы авторизовались как  " + string(name)
		c.Data["is_login"] = true
		c.Redirect("/user",302)
	} else if o.Read(&models.User{Name:name}, "name") == nil && o.Read(&models.User{Pass:pass}, "pass") != nil {
		c.Data["msg"] = "Неправильный пароль"
		c.Data["is_login"] = false
	} else {
		c.Data["msg"] = "Зарегистрируйтесь"
		c.Data["is_login"] = false
	}
	c.Layout = "layout.html"
	c.TplName = "login.html"
}

func (c *LoginController) User() {
	SetIsLogin(&c.BaseController)
	c.Data["msg"] = models.GetUserName(c.GetSession("auth")) //Read name from db by user id
	c.TplName = "user.html"
	c.Layout = "layout.html"
}
