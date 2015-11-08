package main 
import  "fmt"

func main() {
  var a  [5]int //定义一个一维数组
  fmt.Println("emp:",a)

  a[4]=1000//赋值
  fmt.Println("set",a)
  fmt.Println("get:",a[4])
  fmt.Println("len:",len(a))

  b:= [5]int{1,2,3,4,5}//定义数组并初始化数组
  fmt.Println("dcl",b)
  var twoD  [2][3]int //定义二维数组
  for i :=0;i<2;i++{
  	for j:=0;j<3;j++{
  		twoD[i][j]=i+j
  	}
  }
fmt.Println("2d:",twoD)


a1:=[3]int{1,2,3}//定义了长度为3的int数组
b1:=[10]int{1,2,3} //定义了长度为10的数组，其中3个元素初始化，其它默认为0
c:=[...]int{4,5,6}//使用...的方式，go 根据元素个数来计算长度

fmt.Println(a1)
fmt.Println(b1)
fmt.Println(c)

/*
在很多应用场景中，数组并不能满足我们的需求。在初始定义数组时，我们并不知道需要多大的数组，因此我们就需要“动态数组”。在Go里面这种数据结构叫slice
slice并不是真正意义上的动态数组，而是一个引用类型。slice总是指向一个底层array，slice的声明也可以像array一样，只是不需要长度。
*/
s:=make([]string,3)//定义一个切片，
fmt.Println("emp",s)
s[0]="a"
s[1]="b"
s[2]="c"
fmt.Println("set1:",s)
fmt.Println("get1:",s[2])
fmt.Println("len1:",len(s))
s= append(s,"d")
fmt.Println(s)

slice :=[]byte{'a','b','c','d'}
fmt.Println("2-4 的数字",slice[2:4])//输出第3个开始，第4个结束



}