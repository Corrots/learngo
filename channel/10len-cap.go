package main

import "fmt"

func sender(c chan int)  {
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	c <- 5
	close(c)
}

func main() {
	c := make(chan int, 3)

	go sender(c)

	fmt.Printf("1--len(c) = %v, cap(c) = %v\n", len(c), cap(c))

	for val := range c{
		fmt.Printf("val = %v, len(c) = %v\n", val, len(c))
	}
}