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

func init() {
	orm.RegisterModel(new(models.MyUser))
	orm.RegisterDataBase("default", "mysql", "goweb:goweb@/goweb?charset=utf8")
}

func main() {

	// register model
	name := "default"

	// Drop table and re-create.
	force := true

	// Print log.
	verbose := true

	// Error.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}

	// set default database
	//beego.AppConfig.String("pguser")
	//beego.AppConfig.String("pgpass")
	//beego.AppConfig.String("pgurls")
	//beego.AppConfig.String("pgdb")

	beego.Run()
}
