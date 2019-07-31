package main

import "fmt"

type Employee struct {
	FirstName, LastName string
	Salary int
	FullTime bool

}


func main() {

	var ross Employee
	ross.FirstName = "ross"
	ross.LastName = "Bing"
	ross.Salary = 1200
	ross.FullTime = true

	fmt.Println(ross)

}
