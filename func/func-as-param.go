package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

type CalcFunc func(int, int) int

func calc(a , b int,f CalcFunc) int {
	r := f(a, b)
	return r
}


func main() {
	addRes := calc(5, 3, add)
	subRes := calc(5, 3, subtract)

	fmt.Println("5 + 3 = ", addRes)
	fmt.Println("5 - 3 = ", subRes)
}
