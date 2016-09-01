package modles

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
    logger := logs.GetLogger()
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

func insert(user User) {
    o := orm.NewOrm()
    err := o.Begin()

    if err != nil {
        err = o.Rollback()
    } else {
        err = o.Commit()
    }
    checkError(err)
    fmt.Print("miss")
}

// NOTE: 检查是否有错误爆出
func checkError(e error)  {
    if e == nil {
        fmt.Println(" Success do it ")
    } else {
        fmt.Println(" get the Error %v", e)
    }
}
