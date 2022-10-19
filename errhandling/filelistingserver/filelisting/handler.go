package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type userError string

func (u userError) Error() string {
	return u.message()
}

func (u userError) message() string {
	return string(u)
}

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	if index := strings.Index(request.URL.Path, prefix); index != 0 {
		return userError("Path must start with " + prefix)
	}
	path := request.URL.Path[len(prefix):] // /list/fib.txt
	file, err := os.Open(path)
	if err != nil {
		//panic(err)
		/*http.Error(writer, err.Error(), http.StatusInternalServerError)
		return*/
		return err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		/*panic(err)*/
		return err
	}

	writer.Write(bytes)
	return nil
}
