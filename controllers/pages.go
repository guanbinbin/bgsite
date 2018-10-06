package controllers

type PagesController struct {
	BaseController
}

//TODO: Make one controller for easy pages
func (c *PagesController) Index() {
	SetHeader(&c.BaseController)
	SetSidenav(&c.BaseController)

	c.Layout = "layout.html"
	c.TplName = "pages/index.html"
}

func (c *PagesController) Thanks() {
	SetHeader(&c.BaseController)
	SetSidenav(&c.BaseController)

	c.Layout = "layout.html"
	c.TplName = "pages/thanks.html"
}

func (c *PagesController) Test() {
	SetHeader(&c.BaseController)
	SetSidenav(&c.BaseController)

	c.Layout = "layout.html"
	c.TplName = "pages/test.html"
}

func (c *PagesController) Map() {
	SetHeader(&c.BaseController)
	SetSidenav(&c.BaseController)

	c.Data["isActive"] = true //For lighting in side nav
	c.Layout = "layout.html"
	c.TplName = "pages/map.html"
}