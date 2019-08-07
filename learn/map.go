package main

import "fmt"

func main()  {
	newMap := make(map[string]int)

	newMap["aaa"] = 1

	newMap["bbb"] = 2

	if value, ok := newMap["aaa"]; ok {
		fmt.Println(value)
		return
	}

	fmt.Println("invalid map key")

}