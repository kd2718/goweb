package models

import (
	"fmt"
	"github.com/astaxie/beego/validation"
)

type MyUser struct {
	Id        uint64 `form:"-"`
	Email     string `orm:"unique;size(100)" form:"email" valid:"Required"`
	Password  string `orm:"size(102)" form:"password,password" valid:"Required"`
	Password2 string `orm:"-" form:"password2,password,Password Again" valid:"Required"`
}

func (self *MyUser) String() string {
	return fmt.Sprintf("MyUser: %v, with password: %v, ID: %v", self.Email, self.Password, self.Id)

}

func (u *MyUser) TableName() string {
	return "auth_user"
}

func (u *MyUser) Valid(v *validation.Validation) {
	if u.Password != u.Password2 {
		fmt.Println("passwords don't match", u.Password, u.Password2)
		v.SetError("password", "Passwords do not match")
	}
}
