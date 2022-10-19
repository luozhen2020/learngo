package main

import (
	"fmt"
	"retriever/mock"
	"retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.zaobao.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"This is a fake Zaobao.com"}
	r = real.Retriever{}
	fmt.Println(download(r))
}
