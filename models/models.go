package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id int
	Name string
	Pass string
}

func init() {
	orm.RegisterModel(new(User))
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