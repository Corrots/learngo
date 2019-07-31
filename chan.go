package main

import "fmt"

func main()  {
	messages := make(chan string)
	go func(message string) {
		messages <- message	//存
	}("Ping")

	fmt.Println(<-messages)	//取
}
