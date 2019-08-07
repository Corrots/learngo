package main

import "fmt"

func main() {
	str := "Hellõ World"

	for index, val := range str {
		fmt.Printf("character at index %d is %c\n", index, val)
	}

	fmt.Println("")

	for _, val := range str{
		fmt.Printf("Character is %c\n", val)
	}
}
