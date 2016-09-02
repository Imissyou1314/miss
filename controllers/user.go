package controllers

import (
    "fmt"
    "time"
    "github.com/astaxie/beego/logs"
    "github.com/astaxie/beego"
    "miss/models"
)

/** 用户控制器*/
type UserController struct {
    beego.Controller
}

/** 登陆注册控制器*/
type LoginController struct {
    beego.Controller
}


var logger = logs.GetLogger()


func init()  {
    fmt.Println("<============= User Controller =====================>")
    logger.Println("<========================================================>")
    logger.Println("<=============== User Controller ========================>")
    logger.Println("<========================================================>")
}

/**
 *  用户注册登陆接口
 */
func (this *LoginController) Post() {
    inputs := this.Input()
    var user models.User
    user.UserName = inputs.Get("userName")
    user.Password = inputs.Get("password")
    user.CreateTime = time.Now().Format("2006-01-02 15:04:05")
    user.Accout = inputs.Get("accout")
    user.PhoneNumber = inputs.Get("phoneNumber")
    user.PhoneMmid = inputs.Get("PhoneMmid")
    models.Insert(user)
}
