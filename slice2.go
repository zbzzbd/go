package main

import "fmt"

func main() {
	var s1 []int
	fmt.Println(s1)
	s2 := make([]int, 3)
	fmt.Println(s2)

	a := []int{1, 2, 3, 4, 5, 6, 7}
	sa := a[2:5]
	sb := a[4:5]
	fmt.Println(sa)
	fmt.Println(sb)

}
