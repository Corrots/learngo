package main

import "fmt"

func greet(c chan string)  {
	//程序从chan中读取数据，但是chan中无数据，故当前goroutine阻塞，直到goroutine中有数据可被读取
	fmt.Println("Hello " + <-c + "!")
}

func main()  {
	//main goroutine开始执行
	fmt.Println("<<-- main started -->>")
	c := make(chan string)

	go greet(c)

	c <- "Ross"

	fmt.Println("<<-- main stopped -->>")
}