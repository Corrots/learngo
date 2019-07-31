package main

import "fmt"

func main() {
	s := "Hellõ World"
	r := []rune(s)

	fmt.Println(s)
	fmt.Println(r)

	fmt.Println(len(s))
	fmt.Println(len(r))

	//遍历字符串
	fmt.Println("遍历字符串")
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c ", r[i])
	}

	fmt.Println("\n遍历输出每个字符的十进制值")

	for i := 0; i < len(r); i++ {
		fmt.Printf("%v ", r[i])
	}

	fmt.Println("\n编码单元")

	for i := 0; i < len(r); i++ {
		fmt.Printf("%x ", r[i])
	}
	fmt.Println("\n证明rune是int32类型")

	for i := 0; i < len(r); i++ {
		fmt.Printf("%T ", r[i])
	}
}
