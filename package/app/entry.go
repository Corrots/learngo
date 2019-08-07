package main

import (
	"fmt"
	"learngo/package/greet/de"
)


func main() {
	fmt.Println(de.Morning)

	fmt.Println("version ==> ", version)

	var a int = b
	var b int = c
	var c int = 4

	fmt.Println(a, b, c )	//undefined: b	undefined: c
}
