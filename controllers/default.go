package controllers

import (
	"bgsite/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Index() {
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
		c.Data["msg"] = "Добро пожаловать, " + string(userId.Name) + "!"
		}
	//Logout
	if c.Ctx.Request.PostForm.Get("quit") != "" {
		c.DestroySession()
		c.Redirect("/",302)
	}
	c.Layout = "layout.html"
	c.TplName = "index.html"
}

func (c *MainController) NameOld() {
	c.Layout = "layout.html"
	c.TplName = "nameold.html"
}

func (c *MainController) Answer() {
	answer, _ := c.GetUint32("answer")
		if answer == 4 {
			c.Data["reply"] = "Верно!"
		} else if answer == 0 {
			c.Data["reply"] = " "
		} else {
			c.Data["reply"] = "Неверно"
		}
	c.Layout = "layout.html"
	c.TplName = "answer.html"
}

func (c *MainController) Calc() {
	firsta, _ := c.GetFloat("firsta")
	seconda, _ := c.GetFloat("seconda")
	action := c.GetString("action")
	if action == "+" {
		c.Data["reply"] = firsta + seconda
	} else if action == "-" {
		c.Data["reply"] = firsta - seconda
	} else if action == "/" {
			if seconda == 0 {
				c.Data["reply"] = "Себя на ноль помножь, пес"
			} else {
				c.Data["reply"] = firsta / seconda
			}
	} else if action == "*" {
		c.Data["reply"] = firsta * seconda
	} else  if action == "" {
		c.Data["reply"] = " "
	} else {
		c.Data["reply"] = "Введите +, -, / или * !"
	}
	c.Layout = "layout.html"
	c.TplName = "calc.html"
}

func (c *MainController) Hello(){
	c.Data["name"] = c.GetString("name")
	c.Data["old"] = c.GetString("old")
	c.Layout = "layout.html"
	c.TplName = "hello.html"
}