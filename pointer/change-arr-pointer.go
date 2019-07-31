package main

import "fmt"

func changeArr(p *[3]int)  {
	(*p)[0] = 2
	(*p)[1] = 3
	(*p)[2] = 4
}


func main() {
	a := [3]int{1, 2, 3}

	changeArr(&a)

	fmt.Println(a)

}
