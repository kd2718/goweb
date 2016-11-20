package models

import "fmt"

type MyUser struct {
	Id       int    `form:"-"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

func (self *MyUser) String() string {
	return fmt.Sprintf("MyUser: %v, with password: %v, ID: %v", self.Email, self.Password, self.Id)

}
