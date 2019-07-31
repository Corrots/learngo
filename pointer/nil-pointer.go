package main

import "fmt"

func main() {
	var pa *int
	var a1 *string

	fmt.Printf("type: %T, value: %v\n", pa, pa)
	fmt.Printf("type: %T, value: %v\n", a1, a1)
}
