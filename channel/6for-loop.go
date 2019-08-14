package main

import "fmt"

func squares(c chan int)  {
	for i := 0; i < 10; i++ {
		c <- i * i
	}
	//
	close(c)
}

func main()  {
	fmt.Println("<<-- main started -->")
	c := make(chan int)

	go squares(c)
	//
	for  {
		val, ok := <-c
		if ok {
			fmt.Printf("Val: %v,ok: %v\n", val, ok)
		} else {
			fmt.Printf("Val: %v,ok: %v, loop broke-->\n", val, ok)
			break
		}
	}

	fmt.Println("<<-- main stopped -->")
}
