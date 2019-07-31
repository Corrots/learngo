package main

import "fmt"

func calcMultiples(factor int, args ...int) []int  {
	//multiples := make([]int, len(args))

	for index, val := range args {
		args[index] = val * factor
	}

	return args
}

func main() {

	s := []int{10, 20, 30}

	mult1 := calcMultiples(2, s...)
	//mult2 := calcMultiples(10, 2, 3, 4, 5)

	//fmt.Println(mult1)
	fmt.Println(s, mult1)
}
