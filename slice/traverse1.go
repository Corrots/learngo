package main

import "fmt"

func makeSquares1(array [10]int)  {
	for index, elem := range array {
		array[index] = elem * elem
	}
}

func main() {
	s := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(s)
	//数组不会改变
	//makeSquares1 函数只是拿到了数组参数的副本
	makeSquares1(s)

	fmt.Println(s)
}
