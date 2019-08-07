package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// 删除slice中index=2的元素

	s = append(s[:2], s[3:]...)

	fmt.Println(s)

}
