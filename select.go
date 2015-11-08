package main 
import "fmt"
import "time"
func main() {
	
	c1:=make(chan string)
	c2:=make(chan string)

	go func() {
		time.Sleep(time.Second*1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second*2)
		c2 <- "two"
	}()

//使用select，当收到c1的值，打印one，当收到c2 打印two
for i := 0; i < 2; i++ {
	select {
		case  msg1:=<-c1: fmt.Println("recived",msg1)
		case  msg2:=<-c2: fmt.Println("received",msg2)
	}
}



}