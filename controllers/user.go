package controllers

import (
	"bgsite/models"
	"encoding/json"
	"strconv"
)

type UserController struct {
	BaseController
}

func (c *UserController) Home() {
	SetHeader(&c.BaseController)

	orders, _ := models.GetOrdersForUser(c.GetSession("auth").(int))

	c.Data["orders"] = orders
	c.Data["msg"] = models.GetUserName(c.GetSession("auth")) //Read name from db by user id
	c.Data["is_inlogin"] = true
	c.TplName = "user/home.html"
	c.Layout = "layout.html"
}

func (c *UserController) OrderInfo() {
	SetHeader(&c.BaseController)
	SetSidenav(&c.BaseController)

	//Get order id from link, convert to int
	orderId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	//Get products in JSON by userId
	products, _ := models.GetProductsFromOrder(c.GetSession("auth").(int))

	//Decode
	var cart map[int]int
	json.Unmarshal([]byte(products[0]), &cart) //Always only one element, convert to []byte to decode JSON

	prod, sum, _ := models.GetProductsAndSum(cart)
	order, _ := models.GetOneOrder(c.GetSession("auth").(int), orderId)

	c.Data["cart"] = prod
	c.Data["sum"] = sum
	c.Data["ordinfo"] = order
	c.Data["orderNo"] = orderId
	c.Data["is_inlogin"] = true //Lighting
	c.Layout = "layout.html"
	c.TplName = "user/orderinfo.html"
}