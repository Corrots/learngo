package main

import "fmt"

//main goroutine deadlock example
func main() {
	fmt.Println("<<-- main started -->>")
	c := make(chan int)
	c <- 12		//fatal error: all goroutines are asleep - deadlock!
	//chan中只有写入，没有读取，导致当前main goroutine阻塞，及all goroutines阻塞

	fmt.Println("<<-- main stopped -->>")
}
