package main

import "fmt"

func init() {
	fmt.Println("app/entry.go ==> init()")
}

var myVersion = FetchVersion()

func main() {
	fmt.Println("app/fetch-version.go ==> fetchVersion()")
	fmt.Println("version ===> ", myVersion)
}