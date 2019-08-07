package main

import "fmt"

func changeValue(p *int)  {
	*p = 2
}

func main() {
	a := 1
	pa := &a

	changeValue(pa)

	fmt.Println(a)
}
