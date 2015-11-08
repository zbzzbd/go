package main 
import "fmt"

type person struct{

	name string
	age int 
}

func main() {
	fmt.Println(person{"bob",20})//直接赋值
	fmt.Println(person{name:"alice",age:30})//加上变量赋予值,通过field:value的方式初始化，这样可以任意顺序
	fmt.Println(person{name:"zhangsan"})// 只赋予name值，age 初始化为0
	fmt.Println(&person{name:"ann",age :40})//加上取地址符号&，输出 &{ann 40}

	s:=person{name:"sean",age :50}
	fmt.Println(s.name)
	sp := &s//使用指针使得sp 与s使用同一个对象
	fmt.Println(sp.age)//使用点进行调用s中的age
     sp.age=51//直接赋值
     fmt.Println(sp.age)
     fmt.Println(s.age)

     ms:= new (person)//使用new 对象，ms是一个指向person 的指针，此时结构题字段的值时他们所属类型的灵芝，此时ms 就是一个person的对象，new返回指针
     ms.name="zhangbingzhen"
     ms.age=10
     fmt.Println(ms.name)
     fmt.Println("ms:",ms)

     ms2:=&person{"zhangbingzhen",25}
     fmt.Println(ms2)
 
    
    
 }

