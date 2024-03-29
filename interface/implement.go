package main

import "fmt"

type Shape interface {
	Area() float64
}

type Object interface {
	Volume() float64
}

type Skin interface {
	Color() float64
} 

type Cube struct {
	side float64
}

func (c Cube) Area() float64 {
	return 6 * (c.side * c.side)
}

func (c Cube) Volume() float64 {
	return c.side * c.side * c.side
}

func main() {
	var s Shape = Cube{3}
	//c := s.(Cube)
	if value, ok := s.(Cube); ok {
		fmt.Println("area of c of type Cube is", value.Area())
		fmt.Println("volume of c of type Cube is", value.Volume())
	}


	value1, ok1 := s.(Object)
	fmt.Printf("dynamic value of Shape 's' with value %v implements interface Object? %v\n", value1, ok1)

	value2, ok2 := s.(Skin)
	fmt.Printf("dynamic value of Shape 's' with value %v implements interface Skin? %v\n", value2, ok2)
}
