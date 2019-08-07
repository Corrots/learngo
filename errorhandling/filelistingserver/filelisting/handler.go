package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
)

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	//获取url path
	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		//http.Error(writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	defer file.Close()

	stream, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(stream)
	return nil
}
