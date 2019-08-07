package main

import "fmt"

func GetCountries() (c []string) {
	countries := []string{"United states", "United kingdom", "Austrilia", "India", "China", "Russia", "France", "Germany", "Spain"} // can be much more

	//c = append(c, countries[:3]...)
	c = make([]string, 3)
	copy(c, countries[:3])

	return
}

func main() {
	countries := GetCountries()

	fmt.Println(cap(countries)) // 9
}