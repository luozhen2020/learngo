package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"xyy/learngo/errhandling/filelistingserver/filelisting"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

type UserError interface {
	error
	message() string
}

func errWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		//panic error
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)
		if err != nil {
			log.Printf("Error handling request: %s\n", err.Error())
			//user error
			if userErr, ok := err.(UserError); ok {
				http.Error(writer, userErr.message(), http.StatusBadRequest)
				return
			}

			//system error
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
