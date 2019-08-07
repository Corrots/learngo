package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		//return a / b
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

func div(a , b int) (q, r int) {
	//q = a / b
	//r = a % b
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()

	fmt.Printf("Calling function with %s with arguments " + "(%d, %d)\n", opName, a, b)

	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main()  {
/*
	result, err := eval(12, 3, "^")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	//fmt.Println(apply(pow, 12, 3))
	//匿名函数作为参数
	fmt.Println(apply(func(i int, i2 int) int {
		return int(math.Pow(float64(i), float64(i2)))
	}, 2, 3))

	//fmt.Println(div(10,12))
	fmt.Println(sum(1, 2, 3, 4, 5))*/

	//指针
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a)
}
