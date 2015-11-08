package main 
import "fmt"
import "time"
func main() {
	
	c1:=make(chan string ,1)
	go func () {
		//休息2秒
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()
    
    //当取出c1的值，如果超过1秒打印出超时
	select {
		case res:= <-c1: fmt.Println(res)
		case  <-time.After(time.Second*1): fmt.Println("timeout 1")
	}


	c2:=make(chan string ,1)
	go func () {
		time.Sleep(time.Second * 2)
	     c2<-"result 2"
	}()

	select {
		case res:=<-c2:fmt.Println(res)
		case <-time.After(time.Second *3): fmt.Println("timeout 2")
	}


}