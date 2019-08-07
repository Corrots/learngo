package main

import "fmt"

func main() {
	ages := make(map[string]int)

	age := map[string]int{
		"mina": 26,
		"mike": 28,
		"john": 32,
	}
	//尝试复制map,这样是不行的
	//因为map是引用数据类型
	ages = age

	fmt.Println("ages => ", ages)

	delete(age, "mina")
	fmt.Println("ages => ", ages)
}
