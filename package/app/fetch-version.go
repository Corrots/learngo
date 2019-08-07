package main

import (
	"fmt"
	version1 "learngo/package/version"
)

func init() {
	fmt.Println("app/fetch-version.go ==> init()")
}

func FetchVersion() string {
	fmt.Println("app/fetch-version.go ==> fetchVersion()")
	return version1.Version
}