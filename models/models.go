package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id int
	UserId string
	Name string
	Pass string
	Hash string
}

func init() {
	orm.RegisterModel(new(User))
}
