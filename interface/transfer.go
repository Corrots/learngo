package main

import (
	"fmt"
	"strings"
)

func explain(i interface{})  {
	switch i.(type) {
	case string:
		fmt.Println("i stored string ", strings.ToUpper(i.(string)))
	case int:
		fmt.Println("i stored int ", i)
	default:
		fmt.Println("i stored sth else", i)
	}
}

func main() {
	explain("hello golang!")	// i stored string  HELLO GOLANG!
	explain(20)	// i stored int  20
	explain(true)	// i stored sth else true
}
