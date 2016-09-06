package controllers
/**
 *  @desc: 登陆模块
 *  @time:16-09-04
 * 提供用户登陆，注册，注销，更改密码
 */
import (
    "fmt"
    "time"
    "github.com/astaxie/beego"
    "miss/models"
)

type LoginController struct {
    beego.Controller
}

//  添加日志管理
// var logger = logs.GetLogger()


func init()  {
    fmt.Println("<============= User Controller =====================>")
    logger.Println("<========================================================>")
    logger.Println("<=============== User Controller ========================>")
    logger.Println("<========================================================>")
}
/**
 *  用户注册接口
 */
func (this *LoginController) Post() {
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

/**
 *  用户注销接口
 */
func (this *LoginController) LoginOut() {
    inputs := this.Input()
    logger.Println("<=== User LoginOut：%s", inputs)
    var user = new(models.User)
    user.PhoneNumber = inputs.Get("phoneNumber")
    user.Status = false
    models.Update(user)
}
