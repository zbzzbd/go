package main 
import "fmt"


/**
1.函数可以返回多个值,比如返回2个参数，直接就两个类型即可， 
2.有返回值， 那么必须在函数的外层添加return语句

*/
//创建一个函数puls
func puls(a int ,b int ) int {
   return a+b
}
//参数类型多个，且类型相同的可以在最后写
func pulsPluss(a ,b ,c int) int {
	return  a+b+c
}
//返回多个数值
func valus()(int,int) {
	return 3,7
} 
//定义匿名函数,closure
func intSeq() func() int  {
	i:=0
	return func () int  {
		i +=1
		return i

	}
}
//递归（Recursion）
func fact(n int ) int {
	if n ==0{
		return 1
	}
	return n*fact(n-1)
}


func main() {
	res:=puls(1,2)//调用函数puls
	fmt.Println("1+2:",res)

	 res=pulsPluss(1, 2, 3)//直接覆盖原来的res
	fmt.Println("1+2+3=",res)

   a,b:= valus()
   fmt.Println(a)
   fmt.Println(b)
   
    _,c:=valus()//如果不想获取第一个值的时候可以使用_代替，进行忽略
  fmt.Println(c)

   //
  nextInt := intSeq()
  fmt.Println(nextInt())
  fmt.Println(nextInt())
  fmt.Println(nextInt())
  fmt.Println(nextInt())
 // fmt.Println(newInts())
  fmt.Println(fact(7))//调用递归函数

}