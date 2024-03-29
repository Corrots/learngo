package main

import (
	"github.com/gpmgo/gopm/modules/log"
	"learngo/errorhandling/filelistingserver/filelisting"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request)  error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request)  {
	return func(writer http.ResponseWriter, r *http.Request) {

		err := handler(writer, r)
		if err != nil {
			log.Warn("Error handling request: %s", err.Error())
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
				http.Error(writer, http.StatusText(code), code)
			case os.IsPermission(err):
				code = http.StatusForbidden
				http.Error(writer, http.StatusText(code), code)

			default:
				code = http.StatusInternalServerError
				http.Error(writer, http.StatusText(code), code)
			}
		}
	}
}

func main() {
	http.HandleFunc("/list/", errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
