package main

import "fmt"

func makeSquares(slice []int)  {
	for index, elem := range slice {
		slice[index] = elem * elem
	}
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(s)

	makeSquares(s)

	fmt.Println(s)
}
