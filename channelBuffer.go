package main 
import "fmt"

func main() {
	messages :=make(chan string , 2)
	messages  <- "bufferde"
	messages  <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	 
}