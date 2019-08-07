package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5}

	for index, val := range a {
		fmt.Printf("a[%d] = %d\n", index, val)
	}
}
