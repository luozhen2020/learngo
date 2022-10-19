package main

import (
	"errhandling/filelistingserver/filelisting"
	"github.com/gpmgo/gopm/modules/log"
	log2 "log"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

type UserError interface {
	error
	Message() string
}

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		//Panic error
		defer func() {
			if r := recover(); r != nil {
				log2.Printf("Panic: %v", r)
			}
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}()

		err := handler(writer, request)

		if err != nil {
			log.Warn("Error occurred while handling request: %s\n", err.Error())

			//User error
			if userErr, ok := err.(UserError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
			}

			//System error
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
	http.HandleFunc("/list/", errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
