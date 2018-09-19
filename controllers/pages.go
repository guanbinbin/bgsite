package controllers

type PagesController struct {
	BaseController
}

func (c *PagesController) Index() {
	SetIsLogin(&c.BaseController)
	c.Layout = "layout.html"
	c.TplName = "index.html"
}

func (c *PagesController) Map() {
	SetIsLogin(&c.BaseController)
	c.Layout = "layout.html"
	c.TplName = "map.html"
}

func (c *PagesController) NameOld() {
	SetIsLogin(&c.BaseController)
	c.Layout = "layout.html"
	c.TplName = "nameold.html"
}

func (c *PagesController) Answer() {
	SetIsLogin(&c.BaseController)
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

func (c *PagesController) Calc() {
	SetIsLogin(&c.BaseController)
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

func (c *PagesController) Hello(){
	SetIsLogin(&c.BaseController)
	c.Data["name"] = c.GetString("name")
	c.Data["old"] = c.GetString("old")
	c.Layout = "layout.html"
	c.TplName = "hello.html"
}