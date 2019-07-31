package main

import (
	"fmt"
	"reflect"
)

func main()  {
	x := 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type: ", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value: ", v.Float())
	//fmt.Println("elem: ", v.Elem().Field(0).String())


	p := reflect.ValueOf(&x)
	t := p.Elem()
	t.SetFloat(7.1)
	fmt.Println("now t's value is ", t)
}
