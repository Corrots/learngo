package main

import "fmt"

type Employee1 struct {
	FirstName, LastName string
	Salary int
	FullTime bool

}

func main() {
	ross := Employee1{
		FirstName: "ross",
		LastName:  "Bing",
		FullTime:  true,
		Salary:    1200,
	}

	fmt.Println(ross)
}
