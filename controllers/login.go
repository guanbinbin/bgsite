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
	//Use DB
	o := orm.NewOrm()
	o.Using("default")

	sess := c.StartSession()
	sess.Set("test","123")

	c.Data["key"] = sess.Get("test")
	c.Data["sessid"] = sess.SessionID()

	//Registration
	name, pass := c.GetString("name"), c.GetString("pass")
	checkUserName := &models.User{Name: name}
	if o.Read(checkUserName, "UserId") == nil { 	//Check if Username exists (if err=nil)
	c.Data["err"] = "Choose another login"
	} else {
		addUser := models.User{Name: name, Pass: pass}
		o.Insert(&addUser) 								//add to DB
	}

	c.TplName = "session.tpl"
}

func (c* LoginController) Register () {

	c.TplName = "register.tpl"

}
