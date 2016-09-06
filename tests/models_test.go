package test

import (
    "fmt"
    "testing"
    "github.com/astaxie/beego/logs"
)

/**
 *  初始化函数
 */
func init()  {
    fmt.Println("<===============================================================>")
    fmt.Println("<==================== Start Test Models ========================>")
    fmt.Println("<===============================================================>")
    // NOTE: set the test Log Config
    //
    logs.SetLogger(logs.AdapterFile,
        `{"filename":"test_project.log",
          "level":7,
          "maxlines":0,
          "maxsize":0,
          "daily":true,
          "maxdays":10}`)
    //get logger
    logs.Info("<======================= Start Test ===========================>")
}

func  TestInsert(t *testing.T) {
    user := User{UserName:"missyou",
                 Password:"123456",
                 Accout:"123456",
                 CreateTime:"2019-16-.8",
                 Status: true,
                 PhoneNumber:"13763012723",
                 PhoneMmid:"123456"}
    Insert(&user)
}
