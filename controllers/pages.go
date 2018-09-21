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