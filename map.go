package main

import "fmt"

func main() {
	m := map[string]string{
		"name":     "abc",
		"language": "Golang",
		"prefer":   "learn",
	}

	m2 := make(map[string]int) // m2 = empty map

	var m3 map[string]int // m3 = nil

	fmt.Println(m)
	fmt.Println(m2)
	fmt.Println(m3)

	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	for k := range m {
		fmt.Println(k)
	}

	fmt.Println("Get values")
	mapLanguage, ok := m["language"]
	fmt.Println(mapLanguage, ok)

	mapMiss, ok := m["miss"]
	if ok {
		fmt.Println(mapMiss, ok)
	} else {
		fmt.Println("invalid key")
	}

	fmt.Println("Delete values")
	prefer, ok := m["prefer"]
	fmt.Println(prefer)
	delete(m, "prefer")
	fmt.Println(m)

}