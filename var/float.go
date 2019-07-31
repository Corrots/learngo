package main

import "fmt"

type abc float64

func main() {
	var f abc = 52.2
	var f1 float64 = 52.2

	fmt.Printf("f has value %v and type %T", f, f)
	fmt.Println("f == f1", f == f1)	//Invalid operation: f == f1 (mismatched types abc and float64)
}
