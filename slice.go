package  main 
import "fmt"

func main() {
	
	var array  =[10]int {1,2,3,4,5,6,7,8,9,0}
	var aslice ,bslice  []int 


	aslice= array[:3]//等价于：［0:3］
	bslice=array[5:] //等价于 ［5:10］
	aslice=array[:]//等价于全部数据

	slice:=array[:4:6]
	fmt.Println(aslice)
	fmt.Println(bslice)
	fmt.Println(slice)

//or配合range可以用于读取slice和map的数据
	for i, v:=range bslice{
		fmt.Println(i,v)
	}
//由于 Go 支持 “多值返回”, 而对于“声明而未被调用”的变量, 编译器会报错, 在这种情况下, 可以使用_来丢弃不需要的返回
	for _,v:=range bslice {
		fmt.Println("v value is :",v)
	}
}