package main

import (
	"fmt"
	"learngo/queue"
)

func main() {
	q := queue.Queue{1}

	q.Push(2)
	q.Push(3)

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	/**/

	//q.Push("abc")
	//q.Push("Death_Stranding")
	q.Push("Bloodbrone")
	q.Push(123)
	q.Push(456)
	fmt.Println(q)
	q.Pop()
	fmt.Println(q)
}
