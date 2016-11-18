package auth

import (
	"github.com/astaxie/beego"
	"fmt"
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
	var u MyUser
	ctrl.ParseForm(&u)
	fmt.Println(u)
	ctrl.Ctx.WriteString(u.String())
	return
}
