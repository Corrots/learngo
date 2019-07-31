package main

import "fmt"

func main()  {
	c, quit := make(chan int), make(chan int)

	go func() {
		c <- 1
		quit <- 0
	}()

	fmt.Println(<- c)
	fmt.Println(<- quit)
}
