package main

import "fmt"

func getDataFromChan(c chan int)  {
	fmt.Println("data in chan is ", <-c)
}

func main() {
	c := make(chan int)

	fmt.Printf("type of `c` is %T\n", c)
	fmt.Printf("value of `c` is %v\n", c)
	go getDataFromChan(c)

	c <- 12
	//fmt.Println("<<-- main stopped -->>")
}
