package main

import "fmt"

func main()  {
	var c chan int
	fmt.Println(c)	// <nil>
	// 一个nil通道是不可用的，不能向它传递数据或读取数据
	// 必须使用make来创建可用的通道
}