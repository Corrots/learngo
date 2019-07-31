package main

import (
	"fmt"
	"learngo/retriever/mock"
	real2 "learngo/retriever/real"
	"time"
)

const url  = "http://www.imooc.com"

// Poster interface
type Poster interface {
	Post(url string, form map[string]string) string
}

type Retriver interface {
	Get(url string) string
}

//定义组合接口
type RetrieverPoster interface {
	Retriver
	Poster
	//Connect(host string)
}

//接口的组合
func session(s RetrieverPoster) string {
	//
	param := map[string]string {
		"name": "Corrots",
		"contents": "接口的组合",
	}
	s.Post(url, param)
	return s.Get(url)
}

//
func post(poster Poster)  {
	poster.Post(url, map[string]string{
		"name": "Corrots",
		"course": "Golang",
	})
}

func download(r Retriver) string {
	return r.Get("http://www.imooc.com")
}

func main() {

	var r Retriver
	retriever := mock.Retriver{Contents:"This is imooc.com"}
	r = &retriever
	inspect(r)

	r = &real2.Retriver{
		UserAgent: "Mozilla/5.0",
		TimeOut: time.Second,
	}
	fmt.Println("\t")
	//fmt.Println(download(r))
	inspect(r)
	fmt.Println("\t")
	//Type assertion
	//realRetriver := r.(*real2.Retriver)
	if mockRetriver, ok := r.(*mock.Retriver); ok {
		fmt.Println(mockRetriver.Contents)
	} else {
		fmt.Println("not a mock Retriever")
	}
	//接口的组合
	fmt.Println("try a session")
	fmt.Println(session(&retriever))

	//String func

}

func inspect(r Retriver)  {
	fmt.Println("Inspecting Retriever: ")
	fmt.Printf("> %T  %v\n", r, r)
	fmt.Print("> Type switch:")
	switch v := r.(type) {
	case *mock.Retriver:
		fmt.Println("Contents", v.Contents)

	case *real2.Retriver:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}
