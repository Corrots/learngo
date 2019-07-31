package main

import "fmt"

func printArray( arr []int)  {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main()  {
	//定义数组
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}		//定义数组并赋初值
	//arr3 := [...]int{9, 4, 3, 5}
	var grid [2][3]int	//定义2行3列的二维数组

	//for i := 0 ; i < len(arr3) ; i++ {
	//	fmt.Println(arr3[i])
	//}

	fmt.Println("\n")

	arr3 := [...]int{9, 4, 3, 5, 6}
	//使用range循环
	for _, v := range arr3 {
		fmt.Println(v)
	}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(grid)

	fmt.Println("\n")

	printArray(arr1[:])
	printArray(arr3[:])

}
