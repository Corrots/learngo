package main

import "fmt"

func updateSlice(s []int)  {
	s[0] = 100
}

func main()  {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	//fmt.Println(arr[2:6])	// [2 3 4 5]
	//fmt.Println(arr[:6])
	//fmt.Println(arr[2:])
	//fmt.Println(arr[:])

	s1 := arr[2:]
	s2 := arr[:]
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After updeteSlice(s1)")

	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After updeteSlice(s2)")

	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

}
