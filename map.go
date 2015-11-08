package main 
import "fmt"


/**
   1.map 无序，每次打印出来都不一样，只能通过key值进行获取
   2.map 长度是不固定的，也就是类似slice ,
   3.map和其他基本型别不同，它不是thread-safe，在多个go-routine存取时，必须使用mutex lock机制
   4.map也是一种引用类型，如果两个map同时指向一个底层，那么一个改变，另一个也相应的改变

   5.make，ake用于内建类型（map、slice 和channel）的内存分配,make返回初始化后的（非零）值 nil。
*/
func main() {
	m:=make(map[string]int)//创建一个map[key-type]value-type
    m["k1"] =7  //设置值
    m["k2"] = 13

   fmt.Println("map:",m)
    
    v1 := m["k1"]// 取出key的值
    fmt.Println("v1:",v1)
    
    fmt.Println("len",len(m)) //内置的len函数同样适用于map，返回key的数量

    delete(m,"k2")  //删除
    fmt.Println(m)
    _,prs :=m["k2"]
    fmt.Println("prs:",prs)

    n:=map[string]int{"foo":1,"bar":2}//创建并初始化数据
    fmt.Println("map:",n)

    rating :=map[string]float32{"c":5,"go":4.5,"Python":4.5,"c++":2}//创建一个字典并初始化

    csharpRating,ok:=rating["c#"] // map有两个返回值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true
    if ok {
        fmt.Println("c# is in map and its rating is :",csharpRating)
    }else{
        fmt.Println("c# is not in  map")
    }

}