package main

import "fmt"

func main() {
	age := map[string]int{
		"mike": 28,
		"john": 33,
		"mina": 35,
		"alex": 27,
	}

	delete(age, "mina")

	for key, val := range age {
		fmt.Println(key, " => ", val)
	}
}
