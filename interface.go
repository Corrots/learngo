package main
//
//import "fmt"
//
//type Human struct {
//	name string
//	age int
//	phone string
//}
//
//type Student struct {
//	Human
//	school string
//	loan float32
//}
//
//type Employee struct {
//	Human
//	company string
//	money float32
//}
//
////
//func (h Human) SayHi()  {
//	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
//}
//
//func (h Human) Sing(lyrics string)  {
//	fmt.Println("La la, la la la, la la la la la...", lyrics)
//}
//
//func (h Human) Guzzle(beerStein string)  {
//	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
//}
////Employee重载Human SayHi方法
//func (e Employee) SayHi()  {
//	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
//}
//
//func (s *Student) BorrowMoney(amount float32)  {
//	s.loan += amount
//}
//
//func (e *Employee) SpendSalary(amount float32)  {
//	e.money -= amount
//}
//
////定义interface
//type Men interface {
//	SayHi()
//	Sing(lyrics string)
//	Guzzle(beerStein string)
//}
//
//type YoungChap interface {
//	SayHi()
//	Sing(song string)
//	BorrowMoney(amount float32)
//}
//
//type ElderlyGent interface {
//	SayHi()
//	Sing(song string)
//	SpendSalary(amount float32)
//}
//
//func main()  {
//	Alice := Student{Human{
//		"Alice",
//		22,
//		"13007615310",
//	},
//	"Oxford",
//	20000,
//	}
//
//	Alice.SayHi()
//	lyrics := "Coming out of my cage"
//	Alice.Sing(lyrics)
//	Alice.loan = 0
//	Alice.BorrowMoney(1000)
//	Alice.BorrowMoney(2000)
//	Alice.BorrowMoney(3000)
//	fmt.Printf("%s's loan is %f total now\n", Alice.name, Alice.loan)
//
//}