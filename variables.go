package main
import "fmt"

var i =1//在函数外部使用则会无法编译通过，所以一般用var方式来定义全局变量
func main() {
	var a string ="initial"
	fmt.Println(a)
	var b,c int = 1,2
	fmt.Println(b,c)
	var d = true
	fmt.Println(d)
	var e int 
	fmt.Println(e)

	f:="short"
	fmt.Println(f)

	var t1,t2,t3 = 1,"3",true

	fmt.Println(t2)
	fmt.Println(t1)
	fmt.Println(t3)

	t4,t5,t6:=2,"5",false    //这种方式最简洁，方便,go 会自动根基响应值初始化t4,t5,t6,；在函数外部使用则会无法编译通过，所以一般用var方式来定义全局变量
	fmt.Println(t4)
	fmt.Println(t5)
	fmt.Println(t6)
}