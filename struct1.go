package main

import (
	"fmt"
)

type test struct{}
type human struct {
	Sex int
}

type person struct {
	Name string
	Age  int
	human
}

type Student struct {
	Name    string
	Age     int
	Contact struct {
		Phone, City string
	}
	human
}

func main() {
	a := test{}
	p := person{}
	p.Age = 13 //赋值

	p1 := person{Name: "zbz", Age: 19, human: human{Sex: 0}}
	p1.human.Sex = 1
	fmt.Println(a, p, p1)
	A(&p1)
	fmt.Println(p1)

	p3 := struct {
		Name string
		Age  int
	}{
		Name: "job",
		Age:  19,
	}
	//p3就是匿名结构
	fmt.Println(p3)
	student1 := Student{Name: "zbz", Age: 19}
	student1.Contact.Phone = "13611873856" //匿名结构的初始化职能通过这个方式进行初始化
	student1.Contact.City = "shanghai"
	fmt.Println(student1)
}

func A(per *person) {
	per.Age = 13
	fmt.Println("A", per)
}
