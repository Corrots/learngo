package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := a[2:4]

	newS := append(s, 55, 66)

	fmt.Println(s)
	fmt.Println(newS)
	fmt.Println(a)

}
