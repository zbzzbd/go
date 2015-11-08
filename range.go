package main 
import "fmt"

func main() {
	nums:=[]int {1,2,3,4,5}
	sum:=0
	//range on silder ,忽略 index, 直接使用_代替
	for _,num := range nums {
		sum +=num 
	}
	fmt.Println("sum:",sum)

//按照条件输出下标
for i,num:= range nums {
	if num ==3 {
		fmt.Println("index:",i)
	}
}
//range on map
kvs := map[string]string{"a":"apple","b":"banana"}
for k,v:=range kvs {
	fmt.Println("%s -> %s \n",k,v)
}

//按照”go“ unicode 编码输出
for i,c:=range "go"{
	fmt.Println(i,c)
}
for _,num:=range nums {
	fmt.Println(num)
}
}