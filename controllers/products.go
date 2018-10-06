package controllers

import (
	"bgsite/models"
	"github.com/astaxie/beego/utils/pagination"
	"strconv"
)

type ProductsController struct {
	BaseController
}

func (c *ProductsController) Latest(){
	SetHeader(&c.BaseController)
	SetSidenav(&c.BaseController)
	pageNum := GetPageNum(&c.BaseController) 	//Get current page number
	totalProducts, _ := models.CountAllProducts() //Get count of all products

	//TODO: Add filters, get from view
	productsToShow := 3
	products, _ := models.GetLatestProducts(productsToShow, pageNum) //Gets products for current page

	pagination.SetPaginator(c.Ctx, productsToShow, totalProducts) 	//Pagination (set c.Data["paginator"])

	c.Data["products"] = products
	c.Data["is_latest"] = true //lighting
	c.Layout = "layout.html"
	c.TplName = "products/latest.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["AddCart"] = "scripts/addcart.html"
	c.LayoutSections["Pagination"] = "pagination.html"
}

func (c *ProductsController) Catalog(){
	SetHeader(&c.BaseController)
	SetSidenav(&c.BaseController)

	categoryId := c.Ctx.Input.Param(":id") //Get num of category,
	catIdInt, _ := strconv.Atoi(categoryId)		//convert it to int
	pageNum := GetPageNum(&c.BaseController) 	//Get current page number
	totalProducts, _ := models.CountProductsById(catIdInt) //Get count of products by id

	productsToShow := 2
	products, _ := models.GetProductsById(categoryId, productsToShow, pageNum) //Get list of products

	//Check what category the user is in (for lighting it in tpl)
	categoryIdSlice := []models.Category{{catIdInt, ""}} //For correct comparison in tpl

	pagination.SetPaginator(c.Ctx, productsToShow, totalProducts) //Set c.Data["paginator"]
	c.Data["categoryId"] = categoryIdSlice
	c.Data["products"] = products
	c.Layout = "layout.html"
	c.TplName = "products/catalog.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["AddCart"] = "scripts/addcart.html"
	c.LayoutSections["Pagination"] = "pagination.html"
}

func (c *ProductsController) Product(){
	SetHeader(&c.BaseController)
	SetSidenav(&c.BaseController)
	productId := c.Ctx.Input.Param(":id") // GET id, convert to int
	prodIdInt, _ := strconv.Atoi(productId)

	//Get single product by id
	product, _ := models.GetProductById(prodIdInt)
	c.Data["product"] = product

	//Check what category the user is in (for lighting it in tpl)
	categoryIdSlice := []models.Category{{product[0].Category_id, ""}}

	c.Data["categoryId"] = categoryIdSlice
	c.Layout = "layout.html"
	c.TplName = "products/product.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["AddCart"] = "scripts/addcart.html"
}