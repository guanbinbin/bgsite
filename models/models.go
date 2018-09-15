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
		o.Using("default")

		userId := &User{Id: userId.(int)}
		o.Read(userId, "id")
		return userId.Name
	}
}