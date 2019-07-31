package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(s)
	fmt.Printf("len(s): %d, cap(s): %d\n\n", len(s), cap(s))

	s = append(s, 7, 8)
	fmt.Println(s)
	fmt.Printf("len(s): %d, cap(s): %d\n", len(s), cap(s))
}
