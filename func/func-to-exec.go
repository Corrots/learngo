package main

import "fmt"

func main() {
	//立即调用函数
	//声明并立即执行
	var Plus = func(a, b int) int {
		return a + b
	}(5, 3)

	fmt.Println("5 + 3 = ", Plus)

}
