package main

import "fmt"

type Shape1 interface {
	Area() float64
}

type Object1 interface {
	Volume() float64
}

type Metrial interface {
	Shape1
	Object1
}

type Cube1 struct {
	side float64
}

func (c Cube1) Area() float64 {
	return 6 * (c.side * c.side)
}

func (c Cube1) Volume() float64 {
	return c.side * c.side * c.side
}

//嵌入接口
func main() {
	c := Cube1{3}
	var m Metrial = c
	var s Shape1 = c
	var o Object1 = c

	fmt.Printf("dynamic type and value of interface m of static type Material is'%T' and '%v'\n", m, m)
	fmt.Printf("dynamic type and value of interface s of static type Shape is'%T' and '%v'\n", s, s)
	fmt.Printf("dynamic type and value of interface o of static type Object is'%T' and '%v'\n", o, o)
}
