package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/kd2718/goweb/models"
	"github.com/astaxie/beego/validation"
	"log"
)

type HomeController struct {
	beego.Controller
}

func (ctrl *HomeController) Get() {
	ctrl.Data["Website"] = "beego.me"
	ctrl.Data["Email"] = "astaxie@gmail.com"
	ctrl.TplName = "home_page.tpl"
}

func (ctrl *HomeController) Post() {
	var u models.MyUser
	ctrl.ParseForm(&u)
	fmt.Println(u)
	ctrl.Ctx.WriteString(u.String())
	return
}

type RegisterController struct {
	beego.Controller
}

func (ctrl *RegisterController) Post() {
	ctrl.TplName = "new_register.html"

	valid := validation.Validation{}
	var u models.MyUser
	ctrl.ParseForm(&u)
	b, err := valid.Valid(&u)
	fmt.Println("registered user:", u, u.Password)
	if err != nil {
		fmt.Println("an error occured", err)
	}
	if !b {
		fmt.Println("bad stuff occured", b)
	for _, err := range valid.Errors {
	    log.Println(err.Key, err.Message)
	}
	}
}
