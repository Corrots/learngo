package main

import "fmt"

func ChangeValue(p *int)  {
	*p = 3
}

func main() {
	a := 1

	ChangeValue(&a)

	fmt.Println(a)
}
