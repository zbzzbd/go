package main 
import "fmt"



/**
1.传指针比较轻量级 (8bytes),只是传内存地址，我们可以用指针传递体积大的结构体。
如果用参数值传递的话, 在每次copy上面就会花费相对较多的系统开销（内存和时间）。
所以当你要传递大的结构体的时候，用指针是一个明智的选择。

2.Go语言中channel，slice，map这三种类型的实现机制类似指针，所以可以直接传递，
而不用取地址后传递指针。（注：若函数需改变slice的长度，则仍需要取地址传递指针）
*/
func zeroval(ival int ) {
	ival=0
}

func zeroptr(iptr *int) {
   *iptr =0	
}

func add1(a *int) int {
	*a = *a +1
	return  *a
}

func main() {
	i:=1
	fmt.Println("initial:",i)
	
	zeroval(i)
	fmt.Println("zerval",i)
    
    zeroptr(&i)
    fmt.Println("zeroptr:",i)//指针既可以使用该指针的地址，也可以使用指针的指向的值
    fmt.Println("pointer",&i)
    

    x:=1
    fmt.Println("x",x)
    x1:=add1(&x)//传指针使得多个函数能操作同一个对象
    fmt.Println("x+1",x1)
    fmt.Println("x",x)
}