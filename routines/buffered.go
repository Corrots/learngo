package main

import "fmt"

func main()  {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 3
	ch <- 5

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}