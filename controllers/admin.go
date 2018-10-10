package controllers

import "bgsite/models"

type AdminController struct {
	BaseController
}

func (c *AdminController) AdmIndex() {
	CheckIfAdmin(&c.BaseController)
	SetAdmHeader(&c.BaseController)

	c.Layout = "admin/layout.html"
	c.TplName = "admin/home.html"
}

func (c *AdminController) Orders() {
	CheckIfAdmin(&c.BaseController)
	SetAdmHeader(&c.BaseController)

	c.Data["orders"], _ = models.GetAllOrders()
	c.Data["is_orders"] = true
	c.Layout = "admin/layout.html"
	c.TplName = "admin/orders.html"
}

func (c *AdminController) Products() {
	CheckIfAdmin(&c.BaseController)
	SetAdmHeader(&c.BaseController)

	c.Data["is_products"] = true
	c.Layout = "admin/layout.html"
	c.TplName = "admin/products.html"
}

func (c *AdminController) Categories() {
	CheckIfAdmin(&c.BaseController)
	SetAdmHeader(&c.BaseController)

	c.Data["is_categories"] = true
	c.Layout = "admin/layout.html"
	c.TplName = "admin/categories.html"
}

func (c *AdminController) Login() {

	c.Layout = "admin/layout.html"
	c.TplName = "admin/login.html"
}

