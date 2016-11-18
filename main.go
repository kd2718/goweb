package main

import (
	//_ "github.com/kd2718/goweb/routers"
	"github.com/astaxie/beego"
	_ "github.com/kd2718/goweb/auth"
)

//type MyController struct{
//	beego.Controller
//}
//
//func (self *MyController) Get() {
//	//self.Ctx.WriteString("I am kory!!!")
//	self.TplName = "home.tpl"
//}

func main() {
	//beego.Router("/", &MyController{})
	beego.Run()
}