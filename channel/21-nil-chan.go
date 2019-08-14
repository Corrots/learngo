package main

import "fmt"

func service(c chan string)  {
	c <- "response"
}

func main() {
	fmt.Println("main started")

	var chan1 chan string

	go service(chan1)

	//应该时刻去检查chan是否为nil
	select {
	case res := <-chan1:
		fmt.Println("response from chan1, ", res)
	default:
		fmt.Println("No goroutines available")
	}

	fmt.Println("main stopped")
}
