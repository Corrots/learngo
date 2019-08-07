package main

import "fmt"

type Employee2 struct {
	firstName, lastName string
	salary int
	fullTime bool

}

func main() {
	ross := &Employee2{
		firstName: "ross",
		lastName:  "Bing",
		salary:    1200,
		fullTime:  true,
	}

	fmt.Println((*ross).firstName)
	fmt.Println((*ross).lastName)
	fmt.Println((*ross).salary)
	fmt.Println((*ross).fullTime)

	fmt.Println("但是 go 提供了另一种语法来访问字段，例如 array 指针。我们不必使用 *，而是直接使用指针来访问字段。")

	fmt.Println(ross.firstName)
	fmt.Println(ross.lastName)
	fmt.Println(ross.salary)
	fmt.Println(ross.fullTime)

}
