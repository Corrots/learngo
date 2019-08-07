package main

import "fmt"

var plus = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("5 + 3 = ", plus(5, 3))
}
