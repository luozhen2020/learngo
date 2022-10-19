package main

import (
	"fmt"
	"time"
	"xyy/learngo/retriever/mock"
	"xyy/learngo/retriever/real"
)

type Retriever interface {
	Get(ur string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get("https://www.zaobao.com")
}

func post(poster Poster) {
	poster.Post("https://www.zaobao.com", map[string]string{
		"name":   "ccmouse",
		"course": "golang",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

const url = "https://www.zaobao.com"

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"Contents": "another faked Zaobao.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{"This is a fake ZaoBao.com"}
	r = &retriever
	inspect(r)
	r = &real.Retriever{
		"Mozilla/5.0",
		time.Minute,
	}
	inspect(r)

	//Type assertion
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	//fmt.Println(download(r))

	fmt.Println("Try a session")
	fmt.Println(session(&retriever))
}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > %T %v\n", r, r)
	fmt.Print(" > Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}
