package main 
import "fmt"
import "time"
//定义一个函数，此函数输出 0，1，2
func f (from  string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from,":",i)
	}
}

func main() {
	f("direct")
	go f("goroutine")//启动了一个goountine，在goroutine还没来得及跑f的时候，主函数已经退出了
    time.Sleep(time.Second)
    go func(msg string) {
		fmt.Println(msg)
	}("going")
	
	 var input string 
     fmt.Scanln(&input)
	 fmt.Println("done")


	 //使用make建立一个信道：
	 var channel chan  int = make (chan int )
	 channel1:=make(chan int)

	 

}