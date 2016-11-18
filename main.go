package main

import (
	//_ "github.com/kd2718/goweb/routers"
	"github.com/astaxie/beego"
	_ "github.com/kd2718/goweb/auth"
)

func main() {
	beego.Run()
}