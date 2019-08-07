package main

import (
	"fmt"
	"time"
)

func printHello()  {
	//增加一个sleep调用
	time.Sleep(time.Millisecond * 15)
	fmt.Println("Hello Golang")
}

func main() {
	fmt.Println("<<--main function started-->")
	// call function
	go printHello()

	time.Sleep(time.Millisecond * 10)
	fmt.Println("<<--main function stopped-->")
}
