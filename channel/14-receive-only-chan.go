package main

import "fmt"

func main() {
	roc := make(chan<- int)
	soc := make(<-chan int)

	fmt.Printf("type of roc is %T, val = %v\n", roc, roc)
	fmt.Printf("type of soc is %T, val = %v\n", soc, soc)
}
