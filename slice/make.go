package main

import "fmt"

func main() {
	s1 := make([]int, 2, 4)
	s2 := []int{1, 2, 3}

	n1 := copy(s1, s2)
	fmt.Println("\n",n1)	//2
	fmt.Println(s1)	//[1 2]
	fmt.Println(s2)	//[1 2 3]

}
