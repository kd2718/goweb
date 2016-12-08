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
	ctrl.TplName = "home_login.html"
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
		ctrl.TplName = "home_login.html"
	}
	ctrl.Ctx.WriteString(extracted.String())
	return
}

type RegisterController struct {
	beego.Controller
}

type User struct {
    Id    int         `form:"-"`
    Name  interface{} `form:"username"`
    Age   int         `form:"age,text,age:"`
    Sex   string
    Intro string `form:",textarea"`
}

func (ctrl *RegisterController) Get() {
	ctrl.Data["Form"] = & models.MyUser{}
	ctrl.TplName = "home_register.tpl"
}

func (ctrl *RegisterController) Post() {
	ctrl.TplName = "new_register.tpl"
	o := orm.NewOrm()
	errors := map[string]string{}
	var count int64

	valid := validation.Validation{}
	var u models.MyUser
	ctrl.ParseForm(&u)
	log.Println("test in register...", valid)
	log.Println("email:", u.Email, "pass1", u.Password, "pass2:", u.Password2)
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
		errors["blah"] = err.Message
		//ctrl.Data["errors"] = &errors
		//ctrl.TplName = "new_register.html"
		//return

	}
	if b {
		count, err = o.Insert(&u)
		if err != nil {
			fmt.Println("error inserting into db", err)
			errors["Insert"] = fmt.Sprint(err)
		}
	}

	if len(errors) > 0 {
		log.Println("errors found")
		ctrl.TplName = "home_register.tpl"
		ctrl.Data["errors"] = errors
		ctrl.Data["Form"] = &models.MyUser{}
	}

	log.Println("inserted u", u, "total inserted:", count)
	ctrl.Data["u"] = &u
}
