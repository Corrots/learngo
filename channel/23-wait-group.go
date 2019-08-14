package main

import (
	"fmt"
	"sync"
	"time"
)

func service(wg *sync.WaitGroup, instance int)  {
	time.Sleep(2 * time.Second)
	fmt.Println("service called on instance", instance)
	wg.Done()
}

func main() {
	fmt.Println("main stopped")
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go service(&wg, i)
	}

	wg.Wait()
	fmt.Println("main stopped")
}
