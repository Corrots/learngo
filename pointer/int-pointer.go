package main

import "fmt"

//让我们来创建一个类型为 int(整形) 的变量，并创建一个名为 pa 的指针来指向它。

func main() {
	a := 1

	pa := &a	//将变量a的内存地址分配给一个指针变量pa,此时pa虽然没有明确表示它指向哪个变量，但是可以通过内存地址来找到所保存的值

	c := *pa	//* 解引用指针，


	*pa = 2

	fmt.Println(a)

	fmt.Println(pa)	//存储的是内存地址
	fmt.Println(c)

}
