package main

import (
	//_ "github.com/kd2718/goweb/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kd2718/goweb/models"
	_ "github.com/kd2718/goweb/routers"
)

type Test struct {
	Id   uint64
	Name string
}

func init() {
	orm.RegisterDataBase("default", "mysql", "goweb:goweb@/goweb?charset=utf8")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterModel(new(models.MyUser), new(Test))

}

func main() {
	fmt.Println("here we go")

	name := "default"

	//if true,  Drop table and re-create.
	force := false

	// Print log.
	verbose := true

	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println("syncdb err", err)
	}

	beego.Run()
}
