package models

import(
    "fmt"
    "github.com/astaxie/beego/orm"
    "github.com/astaxie/beego/logs"
    _ "github.com/go-sql-driver/mysql"
)

/**
 *  用户个体信息数据库表
 */
type User struct {
    /** 主键*/
    Id              int64
    /** 用户名*/
    UserName        string
    /** 用户密码*/
    Password        string
    /** 用户账号*/
    Accout          string
    /** 用户注册时间*/
    CreateTime      string
    /** 用户手机号*/
    PhoneNumber     string
    /** 用户当前状态*/
    Status          bool
    /** 用户手机MMID号*/
    PhoneMmid       string
}

/**
 *   地区信息数据库表
 */
type AddressMsg struct {
    /** 主键*/
    Id              int64
    /** 地址名*/
    AddressName     string
    /** 经度*/
    AddressLog      string
    /** 维度*/
    AddressLan      string
    /** 进行销售的次数*/
    Count_number    int
    /** 最后的销售时间*/
    LastTime        string
    /** 总销售数*/
    SaleCountNumber   int
    /** 最后一次销售的数量*/
    LastSaleNumber    int
    /** 计划销售的数量*/
    PlanSaleNumber     int
    /** 最后一名销售员进行的备注*/
    AddressDesc        string
}

/**
 *  销售员的销售记录数据库表
 */
type SaleLog struct {
    /** 主键*/
    Id              int64
    /** 销售员的ID*/
    SaleerId        int64
    /** 销售数量*/
    SaleCountNumber int64
}

/**
 *  销售地区记录数据库表
 */
 type AddressLog struct {
     /** 主键*/
     Id             int64
     /** 地区的主键*/
     AddressId      int64
     /**  销售员的Id号*/
     SaleerId       int64
     /** 进行销售的时间*/
     SaleTime       string
     /** 销售的数量*/
     SaleNumber     int
     /** 销售员进行地区的备注*/
     AddressDesc    string
 }

type Instace interface {

}


var logger = logs.GetLogger()

func init()  {

    //添加日志管理
    logs.SetLogger(logs.AdapterFile,
        `{"filename":"project.log",
          "level":7,
          "maxlines":0,
          "maxsize":0,
          "daily":true,
          "maxdays":10}`)
    //get logger

    logger.Println("miss")
    logs.Info("this %s cat is %v years old", "yellow", 3)
    // set miss database
    // orm.RegisterDataBase("default", "mysql",
        //  "root:root@tcp(127.0.0.1:3366)/miss?charset=utf8", 35)

     // set Debug
    //  orm.Debug  = true

     //需要再init中注册定义的model
    orm.RegisterModel(new(User), new(AddressMsg), new(SaleLog),new(AddressLog))

    //create table
    // orm.RunSyncdb("default", false, true)
}

func main() {
    mysqlOrm := orm.NewOrm()
    user := User{UserName:"Imissyou",
                 Password:"123456",
                 Accout:"123456",
                 CreateTime:"2019-16-.8",
                 Status: true,
                 PhoneNumber:"13763012723",
                 PhoneMmid:"123456"}
    id, err := mysqlOrm.Insert(&user)
    fmt.Print("Id : %d, Error: %v\n", id, err)
}

// NOTE: 获取Orm对象
func  getOrm() orm.Ormer {
    return orm.NewOrm()
}

// NOTE: 读取对象
func Query(instance interface{}) interface{} {
    logger.Println("<=== read %s", instance)
    o := getOrm()
    err := o.Read(&instance)
    if checkOrmErr(err) {
        logger.Println("%s ===>", instance)
        return instance
    }else {
        logger.Println("获取数据为空 ===>")
        return nil
    }
}

/**
 *   插入操作
 */
func Insert(instace interface{}) {
    o := orm.NewOrm()
    err := o.Begin()

    o.Insert(instace)

    if err != nil {
        err = o.Rollback()
    } else {
        err = o.Commit()
    }
    checkError(err)
    fmt.Print("miss")
}

/**
 *  删除数据
 * instace: 操作对象
 */
func Delete(instace interface{}) {

}

func checkOrmErr(err interface{}) bool {
    if err == orm.ErrNoRows {
        fmt.Println("====>查询不到")
        logger.Println("=== 查询不到 ===")
        return false
    } else if err == orm.ErrMissPK {
        fmt.Println("=====> 找不到主键")
        logger.Println("=== 找不到主键 ===")
        return false
    } else if err != nil {
        fmt.Println("=====> 操作数据库出错")
        logger.Println("=== 操作数据库异常 ===")
        return false
    }else {
        return true
    }
}

// NOTE: 检查是否有错误爆出
func checkError(e error)  {
    if e == nil {
        fmt.Println(" Success do it ")
    } else {
        fmt.Println(" get the Error %v", e)
    }
}
