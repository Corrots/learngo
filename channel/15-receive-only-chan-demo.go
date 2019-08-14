package main

import "fmt"

func receiveOnlyChan(roc <-chan string)  {
	//roc <- "aaa"	//invalid operation: roc <- "aaa" (send to receive-only type <-chan string)
	fmt.Println("hello ", <-roc + "!")
}

func main() {
	fmt.Println("main started")
	c := make(chan string)

	go receiveOnlyChan(c)

	c <- "John"

	fmt.Println("main stopped")
}
