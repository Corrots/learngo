package main

import "fmt"

//自定义可变参数函数
func getMultiples(factor int, args ...int) []int {
	multiples := make([]int, len(args))

	for index, val := range args {
		multiples[index] = val * factor
	}

	return multiples
}



func main() {
	s := []int{10, 20, 30}
	mult1 := getMultiples(2, s...)
	mult2 := getMultiples(3, 1, 2, 3, 4)


	fmt.Println(mult1)
	fmt.Println(mult2)
}
