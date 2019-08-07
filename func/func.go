package main

import "fmt"

func addMult(a, b int) (int, int) {
	return a + b, a * b
}

func main() {
	addRes, multRes := addMult(2, 5)

	_, m1 := addMult(3,6)

	fmt.Println(addRes, multRes)

	fmt.Println(m1)
}
