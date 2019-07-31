package main

import "fmt"

func swap(a , b int) (int, int) {
	//*b, *a = *a, *b
	return b, a
}

func main()  {
	a, b := 3, 4
	//swap(&a, &b)
	a, b = swap(a, b)
	fmt.Println(a, b)
}