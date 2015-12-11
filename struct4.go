package main

import (
	"fmt"
)

const (
	WHITE = iota
	BLACK
	BULE
	RED
	YELLOW
)

type Color byte //Color 作为byte 的别名
type Box struct {
	width, height, depth float64
	color                Color
}
type BoxList []Box //

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

func (b *Box) SetColor(c Color) {
	b.color = c
} //这里的b.color 也可以写成*b.color， 之所以b.color可以，是因为go 语言中自动转化啦

func (b1 BoxList) BiggestsColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range b1 {
		if b.Volume() > v {
			v = b.Volume()
			k = b.color
		}
	}
	return k
}

func (b1 BoxList) PaintItBlack() {
	for i, _ := range b1 {
		b1[i].SetColor(BLACK)
	}
}

func (c Color) String() {
	stirngs := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return stirngs[c]
}
func main() {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{2, 4, 5, BULE},
		Box{3, 4, 5, WHITE},
		Box{4, 3, 2, YELLOW},
	}
	fmt.Printf("We have %d boxes in our set \n", len(boxes))
	fmt.Println("the volume of first one is :", boxes[0].Volume(), "cm3")
	fmt.Println("The color of the last one is ", boxes[len(boxes)-1].color.String())
	fmt.Println("The biggest one is", boxes.BiggestsColor().String())
	fmt.Println("let's paint them all black")
	boxes.PaintItBlack()
	fmt.Println("The color of the second one is ", boxes[1].color.String())
	fmt.Println("Obviously,now, the biggest one is", boxes.BiggestsColor().String())
}
