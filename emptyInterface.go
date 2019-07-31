package main

import "fmt"

type EmptyInterface interface {}

func main()  {
	var a  interface{}
	i := 5
	s := "Sekiro: Shadows Die Twice"

	a = i
	fmt.Println(a)

	a = s
	fmt.Println(a)
}