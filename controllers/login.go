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
	u := o.QueryTable("user")
	c.Data["bd"] = u
	// c.Redirect("/login", 302)
	c.TplName = "register.tpl"
}

func (c* LoginController) Login () {
	sess := c.StartSession()
	sess.Set("test","123")

	c.Data["key"] = sess.Get("test")
	c.Data["sessid"] = sess.SessionID()

	c.TplName = "login.tpl"

}
