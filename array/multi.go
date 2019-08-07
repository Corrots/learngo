package main

import "fmt"

//多维数组
func main() {
	a := [...][2]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(a)	// [[1 2] [3 4] [5 6]]

	//遍历二维数组

	for index, val := range a {
		for key, elem := range val {
			fmt.Printf("a[%d][%d] = %d\n", index, key, elem)
		}
	}
}
