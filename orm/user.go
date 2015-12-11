package  main 

import (
"fmt"
"github.com/astaxie/beego/orm"
_ "github.com/go-sql-driver/mysql"
)


/*
Beego中数据库的处理很简答，用beego/orm这个工具就可以了。
对于数据库操作主要就是增删查改CURD操作。ORM可以自己建表，
我们只需要定义好相关的结构体就好了，然后用相关的函数去注册好就OK
*/
type User struct {
	Id   int 
	Name  string `orm:"size(100)"`  //表示限制大小长度
}
 
 //注册模型，init 函数自动执行
func init() {
	//设置默认数据库
	//  参数default        数据库的别名，用来在ORM中切换数据库使用
    // mysql        driverName
    // "root"        对应的链接字符串
    //30 设置最大连接数
	orm.RegisterDataBase("default","mysql","root:root@/zhifu?charset=utf8",30)

	//注册数据库
	orm.RegisterModel(new(User))

	//创建表
    //orm.RunSyncdb("default",false,true)

}

func main() {
   o:=orm.NewOrm()  //注册新的orm
   orm.Debug =true
   user:= User{Name:"小明"} // 创建一个user对象

   //insert
   id,err := o.Insert(&user)
   fmt.Printf("Id: %d,ERR: %v \n", id,err)

   var maps []orm.Params
   nums ,err := o.Raw("SELECT * FROM user where name =?","slene").Values(&maps)
    if nums >0 {
    	fmt.Println(maps[0]["id"])
    }
    var users []*User
    num,err := o.QueryTable("user").Filter("name","小明").All(&users,"Id","Name")
    fmt.Printf("Returned Rows Num: %d,%v", num,err)
   
 }