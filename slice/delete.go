package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// delete element at index=2
	s = append(s[:2], s[3:]...)

	fmt.Println(s)
}
