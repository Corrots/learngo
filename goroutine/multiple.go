package main

import (
	"fmt"
	"time"
)

var start time.Time

func init()  {
	start = time.Now()
}

func getChars(s string)  {
	for _, c := range s {
		fmt.Printf("%c at time %v\n", c, time.Since(start))
		time.Sleep(10 * time.Millisecond)
	}
}

func getDigits(s []int)  {
	for _, d := range s {
		fmt.Printf("%d at time %v\n", d, time.Since(start))
		time.Sleep(30 * time.Millisecond)
	}
}

func main() {
	fmt.Println("<<--main function started-->")
	// call function
	go getChars("Hello")
	//
	go getDigits([]int{1, 2, 3, 4, 5})

	time.Sleep(200 * time.Millisecond)
	fmt.Println("\n<<--main function stopped-->")
}
