package controllers

import (
	"bgsite/models"
	"github.com/astaxie/beego/utils/pagination"
	"strconv"
)

type PagesController struct {
	BaseController
}

func (c *PagesController) Index() {
	c.Layout = "layout.html"
	c.TplName = "index.html"
}

func (c *PagesController) Latest(){
	SetIsLogin(&c.BaseController)
	GetCategoriesForSideNav(&c.BaseController)
	pageNum := GetPageNum(&c.BaseController) 	//Get current page number
	totalProducts, _ := models.CountAllProducts() //Get count of all products

	//TODO: Get from view
	productsToShow := 3

	products, _ := models.GetLatestProducts(productsToShow, pageNum) //Gets products for current page

	pagination.SetPaginator(c.Ctx, productsToShow, totalProducts) 		//Pagination (set c.Data["paginator"])
	c.Data["products"] = products
	c.Layout = "layout.html"
	c.TplName = "latest.html"
}

func (c *PagesController) Catalog(){
	SetIsLogin(&c.BaseController)
	GetCategoriesForSideNav(&c.BaseController)

	categoryId := c.Ctx.Input.Param(":id") //Get num of category,
	catIdInt, _ := strconv.Atoi(categoryId)		//convert it to int
	pageNum := GetPageNum(&c.BaseController) 	//Get current page number
	totalProducts, _ := models.CountProductsById(catIdInt) //Get count of products by id

	//TODO: Get from view
	productsToShow := 2
	products, _ := models.GetProductsById(categoryId, productsToShow, pageNum) //Get list of products

	//Check what category the user is in (for lighting it in tpl)
	categoryIdSlice := []models.Category{{catIdInt, ""}} //For correct comparison in tpl

	pagination.SetPaginator(c.Ctx, productsToShow, totalProducts) //Set c.Data["paginator"]
	c.Data["categoryId"] = categoryIdSlice
	c.Data["products"] = products
	c.Layout = "layout.html"
	c.TplName = "catalog.html"
}

func (c *PagesController) Product(){
	SetIsLogin(&c.BaseController)
	GetCategoriesForSideNav(&c.BaseController)
	productId := c.Ctx.Input.Param(":id") // GET id, convert to int
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