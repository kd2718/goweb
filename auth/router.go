package auth

import "github.com/astaxie/beego"

func init(){
	beego.Router("/", &HomeController{})
	beego.SetViewsPath("views/")
}