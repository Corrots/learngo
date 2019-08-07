package main

import (
	"fmt"
	_ "learngo/package/greet"
	child "learngo/package/greet/greet"
)

func init()  {
	fmt.Println("app/packagealisa.go ==> init()")
}

func main()  {
	fmt.Println(child.Message)
}
