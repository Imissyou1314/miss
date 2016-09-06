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

// /** 登陆注册控制器*/
// type LoginController struct {
//     beego.Controller
// }


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
func (this *UserController) Post() {
    inputs := this.Input()
    fmt.Println(inputs)
    var user = new(models.User)
    user.UserName = inputs.Get("userName")
    user.Password = inputs.Get("password")
    user.CreateTime = time.Now().Format("2006-01-02 15:04:05")
    user.Accout = inputs.Get("accout")
    user.PhoneNumber = inputs.Get("phoneNumber")
    user.PhoneMmid = inputs.Get("phoneMmid")
    fmt.Println(user)
    models.Insert(user)
}

func (this *UserController) Login() {
    inputs := this.Input()
    logger.Println("<=== %s", inputs)
    var user = new(models.User)
    user.UserName = inputs.Get("userName")
    user.Password = inputs.Get("password")
    user.Accout = inputs.Get("accout")
    user.PhoneNumber = inputs.Get("phoneNumber")
    var userInfo = models.Query(user)
    if userInfo == nil {
        fmt.Println("user is Nil")
    }else{
        fmt.Println(userInfo)
    }
}

// func setUser(inputs beego.Conttroller.Input()) (user models.User) {
//     var user = new(models.User)
//     user.UserName = inputs.Get("userName")
//     user.Password = inputs.Get("password")
//     user.CreateTime = time.Now().Format("2006-01-02 15:04:05")
//     user.Accout = inputs.Get("accout")
//     user.PhoneNumber = inputs.Get("phoneNumber")
//     user.PhoneMmid = inputs.Get("phoneMmid")
//     return user
// }
