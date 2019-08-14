package main

import (
	"fmt"
	"sync"
	"time"
)

const num  = 10000000

func TestFunc(name string, f func())  {
	//当前时间的纳秒Unix时间戳
	start := time.Now().UnixNano()
	f()
	fmt.Printf("task %s cost %d \r\n", name, (time.Now().UnixNano()-start)/int64(time.Millisecond))
}

func TestChan()  {
	var wg sync.WaitGroup
	c := make(chan string)
	wg.Add(1)

	go func() {
		for _ = range c {

		}
		wg.Done()
	}()

	for i := 0; i < num; i++ {
		c <- "123"
	}
	close(c)
	wg.Wait()
}

func main() {
	TestFunc("testChan", TestChan)
}
