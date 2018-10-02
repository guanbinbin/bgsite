package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
)

type User struct {
	Id int
	Name string
	Pass string
}

type Category struct {
	Id int
	Name string
}

type Product struct {
	Id int
	Name string
	Category_id int //TODO: Correct name to CategoryId
	Code int
	Price float64
	Availability int
	Brand string
	Image string
	Description string
}

type Cart struct {
	Id int
	Code int
	Name string
	Price float64
	Quantity int
}

func init() {
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Category))
	orm.RegisterModel(new(Product))
}

func GetProductsForCart(id []string)([]Cart,error){
	o := orm.NewOrm()
	var products []Cart
	var query string
	if len(id) > 0  {
		query = "select id, code, name, price from product where id in (?" + strings.Repeat(",?", len(id) - 1) + ")"
	} else {
		return products, nil
	}

	_, err := o.Raw(query, id).QueryRows(&products)
	if err == nil {
		return products, nil
	} else {
		return nil, err
	}
}

//Count all products
func CountAllProducts()(int64, error){
	o := orm.NewOrm()
	res, err := o.QueryTable("product").Count()
	if err == nil {
		return res, nil
	} else {
		return 0, err
	}
}

//Count all products in category
func CountProductsById(id int)(int64, error){
	o := orm.NewOrm()
	res, err := o.QueryTable("product").Filter("category_id",id).Count()
	if err == nil {
		return res, nil
	} else {
		return 0, err
	}
}

//Get list of n (num) latest products for page
func GetLatestProducts(num, page int)([]Product,error){
	o := orm.NewOrm()
	var latestProducts []Product
	//Offset for pagination
	if page == 0 { page = 1}
	offset := (page - 1) * num
	_, err := o.Raw("select id, name, price, image from product order by id desc limit ? offset ?",num, offset).QueryRows(&latestProducts)
	if err == nil {
		return latestProducts, nil
	} else {
		return nil, err
	}
}

//Get list of num products by id for page
func GetProductsById(id string, num, page int)([]Product,error){
	o := orm.NewOrm()
	var latestProducts []Product
	//Offset for pagination
	if page == 0 { page = 1}
	offset := (page - 1) * num
	_, err := o.Raw("select id, name, price, image from product where category_id = ? order by id desc limit ? offset ?",id, num, offset).QueryRows(&latestProducts)
	if err == nil {
		return latestProducts, nil
	} else {
		return nil, err
	}
}

func GetCategories() ([]Category, error) {
	o := orm.NewOrm()
	var categories []Category
	_, err := o.Raw("SELECT id, name FROM category").QueryRows(&categories)
	if err == nil {
		return categories, nil
	} else {
		return nil, err
	}
}

//Get single product by id
func GetProductById(id int)([]Product, error){
	o := orm.NewOrm()
	var product []Product
	_, err := o.Raw("SELECT * FROM product where id = ?", id).QueryRows(&product)
	if err == nil {
		return product, nil
	} else {
		return nil, err
	}
}

func GetUserName(userId interface{}) string {
	if userId == nil {
		return ""
	} else {
		o := orm.NewOrm()
		userId := &User{Id: userId.(int)}
		o.Read(userId, "id")
		return userId.Name
	}
}

func CheckRegistration(name, pass string) (registeredName, err string) {
	o := orm.NewOrm()
	checkUserName := &User{Name: name}
	if o.Read(checkUserName, "name") == nil { 	//Check if Username exists (if err=nil)
		return "", "Выбранный логин уже занят. Выберите другой"
	} else {
		addUser := User{Name: name, Pass: pass}
		o.Insert(&addUser) 	//add to DB
		return string(name), ""
	}
}

func CheckLogin (name, pass string) (id int, err string){
	o := orm.NewOrm()
	check := &User{Name:name,Pass:pass}
	if o.Read(check,"name","pass") == nil { //Read from POST name, pass, if matching with db, err==nil
		return check.Id, ""
	} else if o.Read(&User{Name:name}, "name") == nil && o.Read(&User{Pass:pass}, "pass") != nil {
		err = "Неправильный пароль"
		return 0, err
	} else {
		err = "Сначала пройдите регистрацию"
		return 0, err
	}
}