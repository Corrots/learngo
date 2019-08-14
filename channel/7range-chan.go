package main

import "fmt"

func getSquares(c chan int)  {
	for i := 0; i< 10; i++ {
		c <- i * i
	}
	close(c) //不关闭channel会造成deadlock
	//getSquares协程执行完毕后退出，main goroutine继续执行，for循环继续从c通道中接收数据，但是已无数据可接收，主线程阻塞，程序死锁
}

func main()  {
	fmt.Println("<<-- main started -->>")
	c := make(chan int)

	go getSquares(c)

	for val := range c {
		fmt.Printf("Val: %v\n", val)
	}

	fmt.Println("<<-- main stopped -->>")
}