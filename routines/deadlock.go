package main

import "fmt"

var ch1, ch2 = make(chan int), make(chan int)

func say(s string)  {
	fmt.Println(s)
	ch1 <- <- ch2
}

func main()  {
	go say("Hello")
	<- ch1
}