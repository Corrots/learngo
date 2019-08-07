package greet

import "fmt"

var Message = "Hey there. I am parent."

func init()  {
	fmt.Println("greet/parent.go ==> init()")
}