package main

import "fmt"

func main() {
	var s string

	s = "Hellõ Golang"
	fmt.Println(s)

	fmt.Println(len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}

	fmt.Println("")

	for i := 0; i < len(s); i++ {
		fmt.Printf("%v ", s[i])
	}

	fmt.Println("")

	for i := 0; i< len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	fmt.Println("")

	for i := 0; i< len(s); i++ {
		fmt.Printf("%T ", s[i])
	}
}
