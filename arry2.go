package main

import "fmt"

func main() {
	a := [...]int{1, 3, 4, 5, 8, 7, 6, 9}

	num := len(a)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if a[i] < a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}

	fmt.Println(a)

	for i := 0; i < 3; i++ {
		v := 1
		fmt.Println(&v)
	}
}
