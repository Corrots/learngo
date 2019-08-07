package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64  {
	return c.radius * c.radius * math.Pi
}

func main()  {
	r1 := Rectangle{10, 20}
	r2 := Rectangle{12, 20}

	c1 := Circle{10}
	c2 := Circle{20}

	fmt.Printf("Area of r1 is: %f \n", r1.area())
	fmt.Printf("Area of r2 is: %f \n", r2.area())
	fmt.Printf("Area of c1 is: %f \n", c1.area())
	fmt.Printf("Area of c2 is: %f \n", c2.area())
}