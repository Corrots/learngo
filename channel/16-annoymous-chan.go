package main

import (
	"fmt"
)

func main() {
	fmt.Println("<<-- main started -->>")

	c := make(chan string)

	go func(ch chan string) {
		fmt.Printf("Hello %v!\n", <-ch)
	}(c)

	c <- "Mike"

	fmt.Println("<<-- main stopped -->>")
}
