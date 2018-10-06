package controllers

import (
	"bgsite/models"
	"encoding/json"
	"strconv"
)

type CartController struct {
	BaseController
}

func (c *CartController) CartAndOrder() {
	SetHeader(&c.BaseController)
	SetSidenav(&c.BaseController)

	cart := c.GetSession("cart").(map[int]int)
	prod, sum, _ := models.GetProductsAndSum(cart)

	c.Data["cart"] = prod
	c.Data["sum"] = sum
	c.Layout = "layout.html"
	c.Data["is_cart"] = true //Lighting

	if c.Ctx.Request.URL.Path == "/order"{
		c.TplName = "cart/order.html"
	} else {
		c.TplName = "cart/cart.html"
	}
}

func (c *CartController) AddToCart() {
	val, _ := c.GetInt("add") // From script
	cart := c.GetSession("cart").(map[int]int) //type assertion interface to map

	//Check if product id exists. Plus count or add to cart
	_, ok := cart[val]
	if ok == true {
		cart[val]++
	} else {
		cart[val] = 1
	}

	count := CountCartSetSessions(&c.BaseController)
	c.Data["json"] = count
	c.ServeJSON()
}

func (c *CartController) RmFromCart() {
	prID, _ := strconv.Atoi(c.Ctx.Input.Param(":id")) //id from router
	cart := c.GetSession("cart").(map[int]int)
	delete(cart,prID)

	CountCartSetSessions(&c.BaseController)
	c.Redirect("/cart",302)
}

func (c* CartController) ChQuantity() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":pr"))
	cart := c.GetSession("cart").(map[int]int)

	if c.Ctx.Input.Param(":mark") == "0"{
		if cart[id] == 1 {
			c.Redirect("/rm/" + c.Ctx.Input.Param(":pr"),302)
		} else {
			cart[id] = cart[id] - 1
		}
	} else {
		cart[id] = cart[id] + 1
	}

	CountCartSetSessions(&c.BaseController)
	c.Redirect("/cart",302)
}

func (c *CartController) MakeOrder() {
	//Get products, encode to JSON
	cart := c.GetSession("cart").(map[int]int)
	products, _ := json.Marshal(cart)
	
	//Get sum of products
	_, sum, _ := models.GetProductsAndSum(cart)

	//Register order in DB:
	//1. Check if authorized
	var userId int
	if c.GetSession("auth") == nil {
		userId = 0
	} else {
		userId = c.GetSession("auth").(int)
	}
	//2. Make address
	var address string
	if c.GetString("street") == ""{
		address = ""
	} else {
		address = c.GetString("street") + " д. " + c.GetString("home") + " кв. " + c.GetString("flat") + ", подьезд " + c.GetString("entrance")
	}
	//3. Perform registration
	models.RegisterOrder(c.GetString("name"), c.GetString("tel"), c.GetString("email"), string(products), address, c.GetString("comments"), userId, sum)

	c.DelSession("cart")
	c.DelSession("cartCount")
	c.Redirect("/thanks",302)
}