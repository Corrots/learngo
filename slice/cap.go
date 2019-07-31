package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	s = append(s, 7, 8)

	fmt.Println(s)
	fmt.Printf("len=%d, cap=%d", len(s), cap(s))

	fmt.Println()
	s =  append(s, 9, 10, 11, 12, 13)

	fmt.Printf("len=%d, cap=%d", len(s), cap(s))

}
