package models

import (
	"fmt"
	"github.com/astaxie/beego/validation"
)

type MyUser struct {
	Id        uint64 `form:"-"`
	Email     string `orm:"unique" form:"email" valid:"Required"`
	Password  string `form:"password" valid:"Required"`
	password2 string `orm:"-" valid:"Required"`
}

func (self *MyUser) String() string {
	return fmt.Sprintf("MyUser: %v, with password: %v, ID: %v", self.Email, self.Password, self.Id)

}

func (u *MyUser) TableName() string {
	return "auth_user"
}

func (u *MyUser) Valid(v *validation.Validation) {
	//if u.Password != u.password2 {
	//	v.SetError("password", "Passwords do not match")
	//}
}
