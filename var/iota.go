package main

import "fmt"

const (
	a = iota + 1	// 1
	_
	b				// 3
	c				// 4
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
