package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	_ "go-sso/routers"
)

func main() {
	orm.RegisterDataBase("default", "mysql", "root:Root!!2018@tcp(47.105.36.188:3306)/CentaSSO?charset=utf8")
	beego.Run()
}
