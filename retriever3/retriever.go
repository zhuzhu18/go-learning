package main

import (
	"fmt"
	"retriever2/mock"
	"retriever2/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}
type Poster interface {
	Post(url string, form map[string]string) string
}

func post(poster Poster) {
	poster.Post("http://www.baidu.com",
		map[string]string{
			"name":   "zhuzhu",
			"course": "golang",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

const url = "http://www.baidu.com"

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked www.baidu.com",
	})
	return s.Get(url)
}
func download(r Retriever) string {
	return r.Get(url)
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
func main() {
	var r Retriever
	retriever := mock.Retriever{"this is a fake www.baidu.com"}
	r = &retriever
	inspect(r)
	r = &real.Retriever{UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute}
	inspect(r)
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}
	fmt.Println("try a session")
	fmt.Println(session(&retriever))

}
