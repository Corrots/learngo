package main

import "fmt"

func main()  {

	//slice扩展
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	s3 := s2[1:3]

	fmt.Println(s1)		//[2 3 4 5]
	fmt.Println(s2)		//===>[5 6]
	fmt.Println(s3)		//[6 7]

	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))
}
