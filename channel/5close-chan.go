package main

import "fmt"

func sayHello(c chan string)  {
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func main()  {
	fmt.Println("<<-- Main started -->>")
	c := make(chan string)
	
	go sayHello(c)
	c <- "Sekiro"
	//close(c)
	c <- "Mario Cart"

	fmt.Println("<<-- Main stopped -->>")
}