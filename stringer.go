package main

import (
	"fmt"
	"strconv"
)

type human struct {
	name string
	age int
	phone string
}

// 通过这个方法 Human 实现了 fmt.Stringer
func (h human) String() string {
	return "❰"+h.name+" - "+strconv.Itoa(h.age)+" years -  ✆ " +h.phone+"❱"
}

func main() {
	Bob := human{"Bob", 39, "000-7777-XXX"}
	fmt.Println("This Human is : ", Bob)
}
