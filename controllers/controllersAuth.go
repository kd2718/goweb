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
	var input models.MyUser
	ctrl.ParseForm(&input)
	fmt.Println(input)
	extracted := models.MyUser{Email: input.Email}
	err := o.Read(&extracted, "Email")

	fmt.Println("err", err)

	fmt.Println("extracted pass", extracted.Password, "input pass", input.Password)

	if extracted.Password != input.Password {
		ctrl.Ctx.WriteString("BAD!!! passwords don't match")
	}
	ctrl.Ctx.WriteString(extracted.String())
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
	b, err := valid.Valid(&u)
	fmt.Println("registered user:", u, u.Password)
	if err != nil {
		fmt.Println("an error occured", err)
	}
	if !b {
		fmt.Println("bad stuff occured", b)
	}
	for _, err := range valid.Errors {
	    log.Println(err.Key, err.Message)
	}
	count, err := o.Insert(&u)
	if err != nil {
		fmt.Println("error inserting into db", err)
	}

	log.Println("inserted u", u, "total inserted:", count)
	ctrl.Data["u"] = &u
}
