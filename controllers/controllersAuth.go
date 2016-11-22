package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/kd2718/goweb/models"
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
	o := orm.NewOrm()
	var u models.MyUser
	ctrl.ParseForm(&u)
	fmt.Println(u)
	err := o.Read(&u, "Password")
	fmt.Println("err", err)
	ctrl.Ctx.WriteString(u.String())
	return
}

type RegisterController struct {
	beego.Controller
}

func (ctrl *RegisterController) Post() {
	ctrl.TplName = "new_register.html"
	o := orm.NewOrm()

	valid := validation.Validation{}
	var u models.MyUser
	ctrl.ParseForm(&u)
	log.Println("test in register...", valid)
	//b, err := valid.Valid(&u)
	//fmt.Println("registered user:", u, u.Password)
	//if err != nil {
	//	fmt.Println("an error occured", err)
	//}
	//if !b {
	//	fmt.Println("bad stuff occured", b)
	//for _, err := range valid.Errors {
	//    log.Println(err.Key, err.Message)
	//}
	o.Insert(&u)
	log.Println("inserted u", u)
	ctrl.Data["u"] = &u
}
