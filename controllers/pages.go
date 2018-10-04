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

func (c *PagesController) Thanks() {
	SetIsLogin(&c.BaseController)
	GetCategoriesForSideNav(&c.BaseController)
	//Создать в БД таблицу с заказами (со ссылками на товары из др. таблицы),
	//присвоить id заказа, записать заказ, послать письмо админу и клиенту о заказе

	//Required
	c.GetString("name")
	c.GetString("tel")
	c.GetString("email")
	//If courier
	c.GetString("street")
	c.GetString("house")
	c.GetString("padik")
	c.GetString("flat")
	c.GetString("comments")

	c.Layout = "layout.html"
	c.TplName = "thanks.html"
}

func (c *PagesController) Order() {
 //TODO:Repeating code!!!!
	SetIsLogin(&c.BaseController)
	GetCategoriesForSideNav(&c.BaseController)

	cart := c.GetSession("cart").(map[int]int) //Get products from session
	v := GetIdsForCart(cart) 						//Convert product ids for SQL query
	prod, _ := models.GetProductsForCart(v)			//Get products by ids
	//Check if cart is empty
	if len(prod)>0 {
		//Add quantity to each element, count sum price
		var sum float64
		for key := range prod {
			prod[key].Quantity = cart[prod[key].Id]
			sum = sum + prod[key].Price * float64(prod[key].Quantity)
		}

		c.Data["cart"] = prod
		c.Data["sum"] = sum
	}

	c.Layout = "layout.html"
	c.TplName = "order.html"
}

func (c *PagesController) Cart() {
	SetIsLogin(&c.BaseController)
	GetCategoriesForSideNav(&c.BaseController)

	cart := c.GetSession("cart").(map[int]int) //Get products from session
	v := GetIdsForCart(cart) 						//Convert product ids for SQL query
	prod, _ := models.GetProductsForCart(v)			//Get products by ids
		//Check if cart is empty
	if len(prod)>0 {
		//Add quantity to each element, count sum price
		var sum float64
		for key := range prod {
			prod[key].Quantity = cart[prod[key].Id]
			sum = sum + prod[key].Price * float64(prod[key].Quantity)
		}

		c.Data["cart"] = prod
		c.Data["sum"] = sum
	}

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
	CountCartSetJson(&c.BaseController)
	c.ServeJSON()
}

func (c *PagesController) RmFromCart() {
	prID, _ := c.GetInt("del")
	cart := c.GetSession("cart").(map[int]int)
	delete(cart,prID)
	c.SetSession("cart", cart)
	CountCartSetJson(&c.BaseController)
	c.ServeJSON()
}

func (c* PagesController) ChQuantity() {
	mark := c.GetString("mark")
	id, _ := c.GetInt("id")
	cart := c.GetSession("cart").(map[int]int)
	if mark == "-"{
		if cart[id] == 1 {
			//TODO: rm from cart (redirect to link)
		} else {
			cart[id] = cart[id] - 1
		}
	} else {
		cart[id] = cart[id] + 1
	}
	c.SetSession("cart",cart)
	CountCartSetJson(&c.BaseController)
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