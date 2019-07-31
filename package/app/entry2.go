package main

import "fmt"

var arr [10]int

func init()  {
	for i := 0; i < 10; i++ {
		arr[i] = i
	}
}

func main()  {
	fmt.Println(arr)
}