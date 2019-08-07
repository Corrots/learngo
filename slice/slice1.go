package main

import "fmt"

func main() {
	var s []int

	fmt.Println(s)
	fmt.Println(s == nil)

	// create an array of int
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//create new slice
	s1 := a[2:4]

	a[2] = 33
	a[3] = 44

	fmt.Println(s1)

	fmt.Println(len(s1))
	fmt.Println("Capacity of s1: ", cap(s1))

	fmt.Println("address of a[2]: ", &a[2])
	fmt.Println("address of s1[0]: ", &s1[0])

	fmt.Println("&a[2] == &s1[0] is", &a[2] == &s1[0])

	s1[1] = 55
	fmt.Println(s1)
}
