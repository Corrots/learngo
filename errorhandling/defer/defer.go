package main

import (
	"bufio"
	"fmt"
	"learngo/errorhandling/fib"
	"os"
)

func tryDefer()  {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("print too many")
		}
	}
}

func writeFile(filename string)  {
	//file, err := os.Create(filename)
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	//err = errors.New("This is an error")
	if err != nil {
		//panic(err)
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathError.Op, pathError.Path, pathError.Err)
		}

		//fmt.Println(err.Error())
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}


func main() {
	//tryDefer()
	writeFile("fib.txt")
}
