package main

import (
	"fmt"
)

/***
1.继承：包含这个匿名字段的 struct 也能调用该method
2.重写：不同接受器，方法名相同
*/
type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  //这个匿名字段， 这就实现了类似java 语言的继承的现象
	school string
}

type Employee struct {
	Human
	company string
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s  you can call me on %s \n", h.name, h.phone)
}

//employee  重写了human 的method
func (e *Employee) SayHi() {
	fmt.Printf("HI ,I am %s i am work at %s ,you can call me on %s \n", e.name, e.company, e.phone)
}

func main() {
	mark := Student{Human{"Mark", 25, "123234234"}, "山中"}
	sam := Employee{Human{"zbz", 25, "13611873856"}, "go langage"}
	mark.SayHi()
	sam.SayHi()

}
