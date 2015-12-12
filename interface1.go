package main

import "fmt"

/**
1.interface 就是方法的组合，通过interface 来定义对象的行为
2.interface 能存储什么样的值，可以存储实现了这个接口的任意类型的变量
3. interface{}
*/
type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  //匿名字段Human
	school string
	loan   float32
}
type Employee struct {
	Human
	company string
	money   float32
}

func (h Human) SayHi() {
	fmt.Printf("Hi,I am %s you can call me on %s \n", h.name, h.phone)
}
func (h Human) Sing(lyrics string) {
	fmt.Println("la la la la .......", lyrics)
}

func (h Human) Guzzle(beeStenin string) {
	fmt.Println("Guzzle GUzzle .......", beeStenin)
}

//employ 重载了huamn的SayHi方法
func (employee Employee) SayHi() {
	fmt.Printf("Hi ,I am %s,I work at %s ,Call me on %s \n", employee.name, employee.company, employee.phone)
}
func (student Student) BorrowMoney(amount float32) {
	student.loan += amount
}

func (e Employee) SpendSalary(amount float32) {
	e.money -= amount
}

//定义接口
type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStenin string)
}
type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}
type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}

func main() {
	mike := Student{Human{"mike", 25, "222-222-xxx"}, "山中", 0.00}
	paul := Student{Human{"paul", 26, "111-222-xxx"}, "上海", 100}
	sam := Employee{Human{"sam", 28, "444-222-333"}, "peixian", 35000}
	Tom := Employee{Human{"tom", 32, "444-333-444"}, "Things Ltd.", 5000}
	//定义Men 类型的变量 i
	var i Men
	//i 能存储student
	i = mike
	fmt.Println("this is Mike, a student:")
	i.SayHi()
	i.Sing("爱情")
	//i 也能存储 employee
	i = Tom
	fmt.Println("this is Tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	//定义了slice Men
	fmt.Println("let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	x[0], x[1], x[2] = paul, sam, mike
	for _, value := range x {
		value.SayHi()
	}
	var a interface{}
	var i1 int = 5
	s := "hello world"
	a = i1
	a = s

}
