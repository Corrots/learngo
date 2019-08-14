package main

import (
	"fmt"
	"time"
)

var start time.Time

func init()  {
	start = time.Now()
}

func service1(ch chan string)  {
	time.Sleep(3 * time.Second)
	ch <- "Hello from service 1"
}

func service2(ch chan string)  {
	time.Sleep(5 * time.Second)
	ch <- "Hello from service 2"
}

func main() {
	fmt.Println("main started ", time.Since(start))

	chan1 := make(chan string)
	chan2 := make(chan string)

	go service1(chan1)
	go service2(chan2)

	select {
	case res := <-chan1:
		fmt.Println("service1: ", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("service2: ", res, time.Since(start))
	case <-time.After(2 * time.Second):
		fmt.Println("No response received", time.Since(start))
		
	}

	fmt.Println("main stopped ", time.Since(start))
}