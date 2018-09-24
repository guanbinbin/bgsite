package controllers

import (
	"bgsite/models"
	"strconv"
)

type PagesController struct {
	BaseController
}

func (c *PagesController) Index() {
	c.Layout = "layout.html"
	c.TplName = "index.html"
}

func (c *PagesController) Catalog(){
	SetIsLogin(&c.BaseController)
	GetCategoriesForSideNav(&c.BaseController)

	//For pagination
	page := c.Ctx.Input.Param(":page")
	pageInt, _ := strconv.Atoi(page)
	c.Data["page"] = page

	//num - number of products to show
	products, _ := models.GetLatestProducts(6, pageInt)
	c.Data["products"] = products
	c.Layout = "layout.html"
	c.TplName = "catalog.html"
}

func (c *PagesController) Category(){
	SetIsLogin(&c.BaseController)
	GetCategoriesForSideNav(&c.BaseController)

	//Get products, add to tpl
	categoryId := c.Ctx.Input.Param(":id") //id from router (from GET)
	products, _ := models.GetProductsById(categoryId,10)
	c.Data["products"] = products

	c.Data["page"] = c.Ctx.Input.Param(":page")
	c.Data["categ"] = categoryId

	//Check what category the user is in (for lighting it in tpl)
	catIdInt, _ := strconv.Atoi(categoryId)
	categoryIdSlice := []models.Category{{catIdInt, ""}} //For correct comparison in tpl
	c.Data["categoryId"] = categoryIdSlice
	c.Layout = "layout.html"
	c.TplName = "category.html"
}

func (c *PagesController) Product(){
	SetIsLogin(&c.BaseController)
	GetCategoriesForSideNav(&c.BaseController)

	//Id from GET
	productId := c.Ctx.Input.Param(":id")
	prodIdInt, _ := strconv.Atoi(productId)

	//Get single product by id
	product, _ := models.GetProductById(prodIdInt)
	c.Data["product"] = product

	//Check what category the user is in (for lighting it in tpl)
	categoryIdSlice := []models.Category{{product[0].Category_id, ""}}
	c.Data["categoryId"] = categoryIdSlice
	c.Layout = "layout.html"
	c.TplName = "product.html"
}

func (c *PagesController) Map() {
	SetIsLogin(&c.BaseController)
	GetCategoriesForSideNav(&c.BaseController)
	c.Data["isActive"] = true //For lighting in side nav
	c.Layout = "layout.html"
	c.TplName = "map.html"
}