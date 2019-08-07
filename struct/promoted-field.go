package main

import "fmt"

//提升字段


type Salary struct {
	basic     int
	insurance int
	allowance int
}

type Employee3 struct {
	firstName, lastName string
	Salary
}

func main() {
	ross := Employee3{
		firstName: "Ross",
		lastName:  "Geller",
		Salary:    Salary{1100, 50, 50},
	}

	ross.basic = 1800
	ross.insurance = 80
	ross.allowance = 100

	fmt.Println(ross)

}
