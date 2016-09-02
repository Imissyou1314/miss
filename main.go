package main

import (
	_ "miss/routers"
	"miss/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

)

func init()  {
	// NOTE: 配置数据库
	//配置mysql的数据库驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//配置MYSQL数据库的默认链接
	orm.RegisterDataBase("default", "mysql",
		 "root:root@tcp(127.0.0.1:3366)/miss?charset=utf8")
	 // NOTE: 测试多表结构
	 orm.RegisterDataBase("miss","mysql",
	 	"root:root@tcp(127.0.0.1:3366)/golang?charset=utf8")
}

func main() {
	// NOTE: 添加数据库操作
	// set Bebug
	orm.Debug = true
	//create db table
	orm.RunSyncdb("default", false, true)
	//create golang db
	orm.RunSyncdb("miss", false, true)

	o := orm.NewOrm()
	o.Using("default")
	user := models.User{UserName:"Imissyou",
                 Password:"123456",
                 Accout:"123456",
                 CreateTime:"2019-16-.8",
                 Status: true,
                 PhoneNumber:"13763012723",
                 PhoneMmid:"123456"}

	o.Insert(&user)

	beego.Run()
}
