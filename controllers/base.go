package controllers

import (
	"bgsite/models"
	"github.com/astaxie/beego"
	"net/url"
	"strconv"
)

type BaseController struct {
	beego.Controller
}

//Add quantity to each element, count sum price
func CountSum (prod []models.Cart, cart map[int]int) float64{
	var sum float64
	for key := range prod {
		prod[key].Quantity = cart[prod[key].Id]
		sum = sum + prod[key].Price * float64(prod[key].Quantity)
	}
	return sum
}

//Count items in cart, add them to session and set json
func CountCartSetJson (c *BaseController) {
	cart := c.GetSession("cart").(map[int]int)
	var count int
	for _, v := range cart {
		count = count + v
	}
	c.SetSession("cartCount", count)
	c.Data["json"] = count
}

//Parse URL to get number after ? (paginator makes urls: /catalog?p=2..999)
func GetPageNum(c *BaseController) int {
	endOfUrl := c.Ctx.Request.URL.RawQuery //Get p = 1..999
	if endOfUrl == "" {
		endOfUrl = "p=1"
	}                                      //if p is empty (first page)
	m, _ := url.ParseQuery(endOfUrl)
	page, _ := strconv.Atoi(m["p"][0]) //get number of page (2..999)
	return page
}

//Get categories for left nav, add them to template
func GetCategoriesForSideNav (c *BaseController){
	categories, _ := models.GetCategories() //return Id, name from db (slice)
	c.Data["categories"] = categories
}

func SetIsLogin(c *BaseController) {
	if c.GetSession("auth") == nil {
		c.Data["is_login"] = false
	} else {
		c.Data["is_login"] = true
	}
	//Cart
	if c.GetSession("cart") == nil {
		cart := make(map[int]int)
		c.SetSession("cart", cart)
	}
	//Number of products in cart
	if c.GetSession("cartCount") == nil {
		c.SetSession("cartCount",0)
	}
	c.Data["cartCount"] = c.GetSession("cartCount")
}

func (c *BaseController) Logout(){
	c.DelSession("auth")
	c.Redirect("/",302)
}