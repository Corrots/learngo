package main

import "fmt"

var complete = make(chan int)

func loop()  {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}

	fmt.Println("len(complete):", len(complete))

	complete <- 0	//执行完毕
}

func main()  {
	go loop()
	<- complete
}
