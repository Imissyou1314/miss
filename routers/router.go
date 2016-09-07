package routers

import (
	"miss/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

	beego.Router("/addNewUser",&controllers.LoginController{})

	beego.Router("/poll", &controllers.MainController{})
	// 登陆接口
	beego.Router("/login", &controllers.LoginController{}, "post:Login")
	// 注销接口
	beego.Router("/login/out", &controllers.LoginController{}, "*:LoginOut")
	// 注册接口
	beego.Router("/register",&controllers.LoginController{}, "*:Register")
	// beego.Get("/miss", func (ctx *Context) {
	// 	ctx.Output.Body([]byte("Hello word"))
	// })
}
