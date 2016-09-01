package main

import (
	_ "miss/routers"
	_ "miss/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

)

func init()  {
	//配置数据库
	//配置mysql的数据库驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//配置数据库的默认链接
	orm.RegisterDataBase("default", "mysql",
		 "root:root@tcp(127.0.0.1:3366)/miss?charset=utf8")
}

func main() {
	// NOTE: 添加数据库操作
	// set Bebug
	orm.Debug = true
	//create db table
	orm.RunSyncdb("default", false, true)

	beego.Run()
}
