package main

import "fmt"

func main() {

	pa := new(int)

	sa := new(string)

	fmt.Printf("data at %v is %v\n", pa, *pa)
	fmt.Printf("data at %v is '%v'\n", sa, *sa)

}
