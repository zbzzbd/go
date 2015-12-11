package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}
type Circle struct {
	radius float64
}

func (r *Rectangle) area() float64 {
	return r.width * r.height
} //加入指针对象是对实例进行操作，而普通的对象则是对对象的副本进行操作
func (c *Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

//reciver
func main() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{5}

	fmt.Println("Area of r1 is:", r1.area())
	fmt.Println("Area of r2 is:", r2.area())
	fmt.Println("Area of c1 is:", c1.area())
	fmt.Println("Area of c2 is:", c2.area())
}
