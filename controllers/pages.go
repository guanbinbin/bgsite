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

func (c *PagesController) Cart() {
	SetIsLogin(&c.BaseController)
	GetCategoriesForSideNav(&c.BaseController)

	//Get product id, call models.GetProductById and render in tpl
	//cart := c.GetSession("cart").(map[int]int)

	c.Data["is_cart"] = true //Lighting
	c.Layout = "layout.html"
	c.TplName = "cart.html"
}

func (c *PagesController) AddToCart() {
	val, _ := c.GetInt("sel") // Get the data "sel" from request
	cart := c.GetSession("cart").(map[int]int) //type assertion interface to map
	//Check if product id exists. Plus count or add to cart
	_, ok := cart[val]
	if ok == true {
		cart[val]++
	} else {
		cart[val] = 1
	}
	c.SetSession("cart",cart)
	//Count items in cart
	var count int
	for _, v := range cart {
		count = count + v
	}
	c.SetSession("cartCount", count)
	c.Data["json"] = count
	c.ServeJSON()
}

func (c *PagesController) Filters() {
	val, _ := c.GetInt("sel") // Get the data "sel" from request
	c.Data["json"] = val
	c.ServeJSON()
}

func (c *PagesController) Test() {

	c.Layout = "layout.html"
	c.TplName = "test.html"
}

func (c *PagesController) Latest(){
	SetIsLogin(&c.BaseController)
	GetCategoriesForSideNav(&c.BaseController)
	pageNum := GetPageNum(&c.BaseController) 	//Get current page number
	totalProducts, _ := models.CountAllProducts() //Get count of all products

	//TODO: Get from view
	productsToShow := 3
	products, _ := models.GetLatestProducts(productsToShow, pageNum) //Gets products for current page

	pagination.SetPaginator(c.Ctx, productsToShow, totalProducts) 	//Pagination (set c.Data["paginator"])

	c.Data["products"] = products
	c.Data["is_latest"] = true //lighting
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