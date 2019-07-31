package main

import "fmt"

func main() {
	age := make(map[string]int)

	age["mina"] = 29
	age["mike"] = 30
	age["john"] = 35

	fmt.Println(age)
	fmt.Println(age["mike"])
	fmt.Println(age["alex"])

	if value, ok := age["alex"]; ok {
		fmt.Println("Alex's age is ", value)
	} else {
		fmt.Println("doesn't exist")
	}
}
