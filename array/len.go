package main

import "fmt"

func main() {
	// define array
	greetings := [...]string{
		"Good morning!",
		"Good afternoon!",
		"Good evening!",
		"Good night!",
	}

	fmt.Println(greetings)
	fmt.Println(len(greetings))
}
