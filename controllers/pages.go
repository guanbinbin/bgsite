package controllers

import (
	"bgsite/models"
	"fmt"
	"strconv"
)

type PagesController struct {
	BaseController
}

func (c *PagesController) Index() {
	cat, _ := models.GetCategories() //return Id, name from db
	c.Data["category"] = cat
	c.Layout = "layout.html"
	c.TplName = "index.html"
}

func (c *PagesController) Catalog(){
	SetIsLogin(&c.BaseController)
	//num - number of products to show
	products, _ := models.GetLatestProducts(12) //return Id, name, price, image path from db
	c.Data["products"] = products

	cat, _ := models.GetCategories() //return Id, name from db
	c.Data["category"] = cat
	c.Layout = "layout.html"
	c.TplName = "catalog.html"
}

//Ended here
func (c *PagesController) Category(){
	SetIsLogin(&c.BaseController)
	cat, _ := models.GetCategories() //return Id, name from db
	c.Data["category"] = cat

	categoryId := c.Ctx.Input.Param(":id") //id from router (from GET)
	products, _ := models.GetProductsById(categoryId,10)

	//Right here, activate buttons
	i,err := strconv.ParseFloat(categoryId,32)
	fmt.Println(err)
	c.Data["data"] = i
	c.Data["products"] = products
	//

	c.Layout = "layout.html"
	c.TplName = "category.html"
}

func (c *PagesController) Map() {
	SetIsLogin(&c.BaseController)
	cat, _ := models.GetCategories() //return Id, name from db
	c.Data["category"] = cat
	c.Layout = "layout.html"
	c.TplName = "map.html"
}