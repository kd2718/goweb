package routers

import (
	"github.com/astaxie/beego"
	"github.com/kd2718/goweb/controllers"
)

func init() {
	beego.Router("/register/", &controllers.RegisterController{})
	beego.Router("/", &controllers.HomeController{})
	beego.SetViewsPath("views/")
}
