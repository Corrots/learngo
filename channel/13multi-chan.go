package main

import "fmt"

func pingfang(c chan int)  {
	fmt.Println("[平方] reading")
	num := <-c
	c <- num * num
}

func lifang(c chan int)  {
	fmt.Println("[立方] reading")
	num := <-c
	c <- num * num * num
}

func main() {
	fmt.Println("<<- main started ->>")
	pfChan := make(chan int)
	lfChan := make(chan int)

	go pingfang(pfChan)
	go lifang(lfChan)

	testNum := 3
	pfChan <- testNum
	lfChan <- testNum

	pfVal, lfVal := <-pfChan, <-lfChan
	fmt.Printf("pfVal = %v, lfVal = %v\n", pfVal, lfVal)
	fmt.Println("pfVal + lfVal = ", pfVal + lfVal)

	fmt.Println("<<- main stopped ->>")
}