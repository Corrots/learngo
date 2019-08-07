package main

import "fmt"

func main() {
	//a := 1
	a := 0x00
	b := 0x0A
	c := 0xFF

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(&a)
	fmt.Println(&b)
	fmt.Println(&c)

	a1 := &a
	fmt.Println(*a1)

}
