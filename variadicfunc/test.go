package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	s3 := []int{7, 8, 9}

	c := append(s1, append(s2, s3...)...)
	fmt.Println(c)
}
