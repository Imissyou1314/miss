package utils

import (
    "fmt"
    "encoding/json"
    "github.com/astaxie/beego/logs"
)

var logger = logs.GetLogger()
/**
 *  构造函数，初始化Logger
 */
func init() {
    fmt.Println("<==== json 构造 函数 ======>")
    logger.Println("配置json Utils 初始化....")
}

/**
 *  对象转Json字符串
 */
func ToJson(json_data interface{}) (b []byte, err error) {
    logger.Println("to json data is:%s", json_data)
    if json_data == nil {
            logger.Println("json data is nil")
    }else{
        b, err:= json.Marshal(json_data)
        return b, err
    }
    return nil, nil
}

/**
 *   json数据转成对象
 */
func JsonToStruct(json_data []byte, instance interface{}) error {
    logger.Println("json to Struct:%s", json_data)
    if json_data == nil {
        logger.Println("json_data is nil")
    }else{
        err := json.Unmarshal(json_data, &instance)
        return err
    }
    return nil
}
