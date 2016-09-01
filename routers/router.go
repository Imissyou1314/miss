package routers

import (
	"miss/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

	beego.Router("/poll", &controllers.MainController{})


	// beego.Get("/miss", func (ctx *Context) {
	// 	ctx.Output.Body([]byte("Hello word"))
	// })
}
