package main 
import "fmt"
import "math"

type geometry interface {
	area() float64
	perim() float64

}

type rect struct {
	width,height  float64
}
type cirlce struct {
	radius  float64
}
// 创建area方法此方法属于rect 对象下的一个方法
func (r rect) area() float64{
	return  r.width*r.height
}
func (r rect) perim() float64 {
	return 2*r.width+2*r.height
}

func (c cirlce) area() float64 {
	return  math.Pi*c.radius*c.radius
}
func (c cirlce) perim() float64  {
	return 2*math.Pi*c.radius
}

 func  measure(g geometry) {
   fmt.Println(g)
   fmt.Println(g.area())
   fmt.Println(g.perim())

 }
func main() {
	r:=rect{width:3,height:4}
	c:=cirlce{radius:5}

	measure(r)
	measure(c)
}
