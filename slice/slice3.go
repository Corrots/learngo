package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s :=  a[2:4]

	fmt.Println("Before -> a: ", a)
	fmt.Println("Before -> s: ", s)
	fmt.Printf("Before -> len(s): %d, cap(s): %d\n", len(s), cap(s))
	fmt.Println("Before -> &a[2] == &s[0] is", &a[2] == &s[0])

	fmt.Println()

	s = append(s, 50, 60, 70, 80, 90, 100, 110)
	fmt.Println("After -> a: ", a)
	fmt.Println("After -> s: ", s)
	fmt.Printf("After -> len(s): %d, cap(s): %d\n", len(s), cap(s))
	fmt.Println("After -> &a[2] == &s[0] is", &a[2] == &s[0])
}
