package main

import "fmt"

func main() {
	ages := make(map[string]int)

	age := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	for key, val := range age {
		ages[key] = val
	}

	fmt.Println(ages)
}
