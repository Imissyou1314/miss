package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Post()  {
	id := c.Ctx.Input.Param(":id")
	fmt.Println(id)
	c.Data["Website"] = "Miss.go"
	c.Data["Email"] = "1255418233@qq.com"
	c.TplName = "index.tpl"
}
