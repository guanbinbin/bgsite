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
	c.Layout = "layout.html"
	c.TplName = "session.html"
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
		c.Data["login"] = true
	}
	c.Layout = "layout.html"
	c.TplName = "register.html"
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
		c.Data["is_login"] = true
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
	//Use db
	o := orm.NewOrm()
	o.Using("default")
	//Check if user is authorized
	id := c.GetSession("auth")
	if id == nil {
		c.Data["is_login"] = false
		c.Data["msg"] = "Вы не авторизованы"
	} else {
		c.Data["is_login"] = true
		//Read name from db by user id
		userId := &models.User{Id: id.(int)}
		o.Read(userId, "id")
		c.Data["msg"] = "Вы авторизованы как " + string(userId.Name)
	}
	//Logout
	if c.Ctx.Request.PostForm.Get("quit") != "" {
		c.DestroySession()
		c.Redirect("/",302)
	}
	c.Layout = "layout.html"
	c.TplName = "user.html"
}
