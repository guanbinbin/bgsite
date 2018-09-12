package controllers

import (
	"bgsite/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	beego.Controller
}

func (c* LoginController) Session () {
	//Button "Login"
	if c.Ctx.Request.FormValue("submit") == "login" {
		c.Redirect("/login",307) //307 - using POST
	}
	//Button "Register"
	if c.Ctx.Request.FormValue("submit") == "register"{
		c.Redirect("/register",307)
	}

	c.TplName = "session.tpl"
}

func (c* LoginController) Register () {
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
	}

	c.TplName = "register.tpl"
}

func (c* LoginController) Login () {
	//Use db
	o := orm.NewOrm()
	o.Using("default")

	//Login
	name, pass := c.GetString("name"), c.GetString("pass")
	check := &models.User{Name:name,Pass:pass}
	if o.Read(check,"name","pass") == nil { //Read from POST name, pass, if matching with db, err==nil
		c.SetSession("auth", check.Id)     //Id to session
		c.Data["msg"] = "Вы авторизовались как  " + string(name)
	} else if o.Read(&models.User{Name:name}, "name") == nil && o.Read(&models.User{Pass:pass}, "pass") != nil {
		c.Data["msg"] = "Неправильный пароль"
	} else {
		c.Data["err"] = "Зарегистрируйтесь"
	}

	c.TplName = "login.tpl"
}
