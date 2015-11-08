package main 
import "fmt"


var ch chan int  =make(chan int )

func foo(msg string) {
	for i := 0; i < 4; i++ {
		fmt.Println(msg,":",i)
	}
   ch <-0	//向 ch 中加数据，如果没有其他goroutine 来取出这个数据的话,那么就会挂起foo，直到main函数把0这个数据取走,否则该信道不会被打开
}

func main() {
	messages:=make(chan string)//创建一个信道
  go func (message string ) {
		messages <- message //给信道进行存消息
	}("ping le")
	msg:= <-messages//从信道中取出消息
	fmt.Println(msg)
   go foo("go")
   <- ch 
   
}
