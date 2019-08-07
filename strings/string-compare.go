package main

import (
	"fmt"
)

func main() {
	fmt.Printf("value of character a is %v of type %T\n", 'a', 'a')
	fmt.Printf("value of character b is %v of type %T\n", 'b', 'b')
	fmt.Println("hence 'b' > 'a' is", 'b' > 'a')
}