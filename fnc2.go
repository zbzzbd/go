package main 
import "fmt"

 	/**
       1.函数也是一种变量，我们可以通过type来定义它，它的类型就是所有拥有相同的参数，相同的返回值的一种类型
   	*/
type testint func(int) bool //声明了一个函数类型

func isOdd(integer int) bool {
	if integer %2==0{
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer %2==0{
		return true
	}
	return false 
}

func filter(slice []int,f testint) []int  {
	var result [] int 
	for _,value:=range slice {
		if f (value){
			result =append(result,value)
		} 
	}
	return result

}

 func main() {
   	
   	slice:=[]int {1,2,3,4,5,6}
   	fmt.Println("slice=",slice)
  
   	odd:=filter(slice,isOdd)//函数当作值来进行传递
   	fmt.Println("odd element of slice are :",odd)
   	even:=filter(slice,isEven)
   	fmt.Println("even element of slice are :",even)
}