package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func main() {
	fmt.Printf("Type of function Add is %T\n", Add)
	fmt.Printf("Type of function Subtract is %T\n", Subtract)
}
