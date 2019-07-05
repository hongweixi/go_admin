package main

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "zadmin/routers"
	"github.com/astaxie/beego"

)

func init() {
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)

	_ = orm.RegisterDataBase("default", "mysql", "root:root@/zadmin?charset=utf8")
}


func main() {
	beego.SetStaticPath("/images","images")
	beego.SetStaticPath("/css","css")
	beego.SetStaticPath("/js","js")
	beego.SetStaticPath("/proto","static/proto")

	beego.BConfig.WebConfig.Session.SessionOn = true

	//导入公共方法


	beego.Run()
}

