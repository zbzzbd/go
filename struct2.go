package main 
import "fmt"

//声明一个新的类型person 
type person struct{
	name string 
	age int 
}
type Skills [] string //定义一个动态的切片

type student struct {
	person    //匿名字段
	student_no int 
    number  int 
    Skills
}


//比较2个人的年龄，返回年龄大的那个人，并切返回年龄差
func Older(p1,p2 person) (person ,int )  {
	
	if p1.age>p2.age{
		return p1,p1.age-p2.age
	} 
	return p2,p2.age-p1.age
}

func main() {
	
	var tom person
	tom.name,tom.age="tom",18
	bobo:=person{age:25,name:"bobo"}
	paul:=person{"paul",26}

	tb_older,tb_diff := Older(tom, bobo)
	tp_older,tp_diff :=Older(bobo, paul)

	fmt.Println("of %s  and  %d",tb_older,tb_diff)
	fmt.Println("of %s and %d ",tp_older,tp_diff)

	bingzhen:=student{person{"bingzhen",25},01,{"anatomy"},number:100}//匿名字段就是这样，能够实现字段的继承
	fmt.Println(bingzhen.student_no,bingzhen.number)
}
