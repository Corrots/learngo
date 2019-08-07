package main

import "fmt"

func main() {
	s1 := make([]int, 0, 10)
	fmt.Println("Before s1 -> ", s1)
	fmt.Printf("len(s1): %d cap(s1): %d\n", len(s1), cap(s1))

	s2 := []int{1, 2, 3}
	s1 = append(s1, s2...)
	fmt.Println("After s1 -> ", s1)
	fmt.Printf("len(s1): %d cap(s1): %d\n", len(s1), cap(s1))
}
