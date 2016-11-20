package main

import (
	//_ "github.com/kd2718/goweb/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/kd2718/goweb/models"
	_ "github.com/kd2718/goweb/routers"
	_ "github.com/lib/pq"
)

func init() {
	// register model
	orm.RegisterModel(new(models.MyUser))

	// set default database
	beego.AppConfig.String("pguser")
	beego.AppConfig.String("pgpass")
	beego.AppConfig.String("pgurls")
	beego.AppConfig.String("pgdb")
	beego.Run()
}
