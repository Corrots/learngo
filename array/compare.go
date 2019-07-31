package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := [3]int{1, 3, 2}
	c := [3]int{1, 1, 1}
	d := [3]int{1, 2, 3}

	fmt.Println("a == b ", a == b)
	fmt.Println("a == c ", a == c)
	fmt.Println("a == d ", a == d)
	fmt.Println("b == d ", b == d)
}
