package main

import (
	"fmt"
	"strconv"
)
import "math"

type (
	byte   int8
	rune   int32
	wenben string //设置别名,最好用英文
)

const x int = 1
const y int = 'A'

const (
	x1 = 1
	x2
	x3
)

func main() {
	fmt.Println(x1, x2)
	var a [0]byte

	fmt.Println(a)
	fmt.Println(math.MaxFloat32)
	var b int
	b = 1
	var c = 1
	d := "hehe"
	fmt.Println(c)
	fmt.Println(b)
	fmt.Println(d)

	f, g, h := 1, 2, 3
	fmt.Println(f, g, h)

	var k int = 65
	m := strconv.Itoa(k)

	fmt.Println(m)

}
