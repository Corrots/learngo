package main

import "fmt"

func main()  {
	slice := []int{1, 2, 3, 4}
	slice2 := slice[1:4]

	slice4 := make([]int, len(slice2))
	copy(slice4,slice2)

	slice[1] = 111;

	fmt.Println(slice2)	//[2 3 4]
	fmt.Println(slice4)

}
