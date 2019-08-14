package main

import (
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"io/ioutil"
	"net/http"
	"sync"
)

func main()  {
	var wg sync.WaitGroup
	urls := []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}

	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			resp, err := http.Get(url)
			if err != nil {
				log.Fatal(url, "--failed--", err.Error())
			}
			//写入文件
			body, err := ioutil.ReadAll(resp.Body)
			error := ioutil.WriteFile("gin.log", body, 0644)
			if err != nil {
				log.Fatal("write file failed--", error)
			}
			//
			//println(body)
			defer resp.Body.Close()
		}(url)
	}
	//
	wg.Wait()
	fmt.Println("done")
}
