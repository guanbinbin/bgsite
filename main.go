package main

import (
	_ "bgsite/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
	beego.BConfig.WebConfig.Session.SessionOn = true
}

