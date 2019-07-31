package main

import "fmt"

func main() {
	var s1 []int
	s2 := []int{1, 2, 3}
	s3 := []int{4, 5, 6, 7}
	s4 := []int{1, 2, 3}

	n1 := copy(s1, s2)

	fmt.Printf("n1=%d, s1=%v, s2=%v\n", n1, s1, s2)
	fmt.Println("copy(s1, s2) ->",n1)

	n2 := copy(s2, s3)
	fmt.Printf("n2=%d, s2=%v, s3=%v\n", n2, s2, s3)

	n3 := copy(s3, s4)
	fmt.Printf("n3=%d, s3=%v, s4=%v\n", n3, s3, s4)


}
