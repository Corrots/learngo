package main

import "fmt"

func testSquares(c chan int)  {
	for i := 0; i <= 3; i++ {
		num := <-c
		fmt.Println(num * num)
	}
}

func main() {
	fmt.Println("main started")
	c := make(chan int, 3)

	go testSquares(c)

	c <- 1
	c <- 2
	c <- 3
	c <- 4

	fmt.Println("main stopped")
}
