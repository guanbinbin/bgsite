package controllers

import (
	"bgsite/models"
	"fmt"
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
	c.Data["err"] = "Choose another login"
	} else if  name == "" && pass == "" {
		c.Data["err"] = "Enter correct login and pass"
	} else if  pass == "" {
		c.Data["err"] = "Enter correct pass"
	} else if  name == "" {
		c.Data["err"] = "Enter correct login"
	} else {
		addUser := models.User{Name: name, Pass: pass}
		o.Insert(&addUser) 							//add to DB
		c.Data["err"] = "Добро пожаловать, "
		c.Data["name"] = name
	}

	c.TplName = "register.tpl"
}

func (c* LoginController) Login () {
	c.SetSession("auth", "id")
	c.GetSession("auth")

	//TODO: repeating code!
	o := orm.NewOrm()
	o.Using("default")

	//Read from db
	user:= &models.User{Name:"12"}
	err := o.Read(user, "name")
	fmt.Println(err, user.Id, user.Pass)

	//c.SetSecureCookie("namesd", "asd","aaaaaaaaaa")
	c.TplName = "login.tpl"
}
