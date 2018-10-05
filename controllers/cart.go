package controllers

import (
	"bgsite/models"
	"encoding/json"
	"fmt"
	"strconv"
)

type CartController struct {
	BaseController
}

func (c *CartController) OrderInfo() {
	SetIsLogin(&c.BaseController)
	GetCategoriesForSideNav(&c.BaseController)
	//TODO: Пересмотри код на повторения и на ненужные переменные (использующиеся один раз/для вывода и т.д.). Просто вывод ф-ии

	//Get order id from link, convert to int
	orderId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	//Get products in JSON by userId
	products, _ := models.GetProductsFromOrder(c.GetSession("auth").(int))

	//Decode
	var cart map[int]int
	json.Unmarshal([]byte(products[0]), &cart) //Always only one element, convert to []byte to decode JSON
	fmt.Println(cart)

	prod, _ := models.GetProductsForCart(cart) //Get products by ids
	sum := CountSum(prod, cart)

	c.Data["cart"] = prod
	c.Data["sum"] = sum

	orders, _ := models.GetOneOrder(c.GetSession("auth").(int), orderId)
	c.Data["ordinfo"] = orders
	c.Data["orderNo"] = orderId
	c.Data["is_inlogin"] = true
	c.Layout = "layout.html"
	c.TplName = "orderinfo.html"
}

func (c *CartController) MakeOrder() {
	SetIsLogin(&c.BaseController)
	GetCategoriesForSideNav(&c.BaseController)

	//Get products, encode to JSON
	cart := c.GetSession("cart").(map[int]int)
	products, _ := json.Marshal(cart)

	prod, _ := models.GetProductsForCart(cart) //Get products by ids
	sum := CountSum(prod, cart)

	//Check if authorized
	var userId int
	if c.GetSession("auth") == nil {
		userId = 0
	} else {
		userId = c.GetSession("auth").(int)
	}

	//Register order in DB
	var address string
	if c.GetString("street") == ""{
		address = ""
	} else {
		address = c.GetString("street") + " д. " + c.GetString("home") + " кв. " + c.GetString("flat") + ", подьезд " + c.GetString("entrance")
	}

	models.RegisterOrder(c.GetString("name"), c.GetString("tel"), c.GetString("email"), string(products), address, c.GetString("comments"), userId, sum)

	c.DelSession("cart")
	c.DelSession("cartCount")
	c.Redirect("/thanks",302)
}

func (c *CartController) Order() {
	SetIsLogin(&c.BaseController)
	GetCategoriesForSideNav(&c.BaseController)

	cart := c.GetSession("cart").(map[int]int) //Get products from session
	prod, _ := models.GetProductsForCart(cart)		//Get products by ids

	sum := CountSum(prod, cart)
	c.Data["cart"] = prod
	c.Data["sum"] = sum

	c.Layout = "layout.html"
	c.TplName = "order.html"
}

func (c *CartController) Cart() {
	SetIsLogin(&c.BaseController)
	GetCategoriesForSideNav(&c.BaseController)

	cart := c.GetSession("cart").(map[int]int) //Get products from session
	prod, _ := models.GetProductsForCart(cart)			//Get products by ids

	sum := CountSum(prod, cart)

	c.Data["cart"] = prod
	c.Data["sum"] = sum


	c.Data["is_cart"] = true //Lighting
	c.Layout = "layout.html"
	c.TplName = "cart.html"
}

func (c *CartController) AddToCart() {
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

func (c *CartController) RmFromCart() {
	prID, _ := c.GetInt("del")
	cart := c.GetSession("cart").(map[int]int)
	delete(cart,prID)
	c.SetSession("cart", cart)
	CountCartSetJson(&c.BaseController)
	c.ServeJSON()
}

func (c* CartController) ChQuantity() {
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
