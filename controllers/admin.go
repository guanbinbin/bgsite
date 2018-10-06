package controllers

type AdminController struct {
	BaseController
}

func (c *AdminController) AdmIndex() {


	c.Layout = "admin/layout.html"
	c.TplName = "admin/home.html"
}

func (c *AdminController) Login() {


	c.Layout = "admin/layout.html"
	c.TplName = "admin/login.html"
}

func (c *AdminController) CheckLogin() {


}
