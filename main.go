package main

import (
	_ "bgsite/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:ormtest@/dev?charset=utf8")
}

func main() {
	//Autocreate tables
	name := "default"
	force := false	 // Drop table and re-create.
	verbose := false // Print log.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}

	//Session
	beego.BConfig.WebConfig.Session.SessionOn = true

	beego.Run()
}

