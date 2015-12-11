package main

import (
	"fmt"
)

const (
	A         = 15
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
)

func main() {
	fmt.Println(B)
	a := 1
	for a <= 3 {
		a++
		fmt.Println(a)

	}

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	switch {
	case a >= 0:
		fmt.Println("a=0")
	case a >= 1:
		fmt.Println("a=1")
	default:
		fmt.Println("None")
	}

LABEL1:
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				break LABEL1
			}
		}
	}
	fmt.Println("OK")

	b := [...]int{99: 1}
	var p *[100]int = &b
	fmt.Println(p)

}
