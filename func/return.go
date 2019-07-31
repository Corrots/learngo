package main

import "fmt"

func addMulti(a, b int) (add, multi int) {
	add = a + b
	multi = a * b

	return	//necessary
}

func main() {
	add, multi := addMulti(3, 4)

	fmt.Println(add, multi)
}
