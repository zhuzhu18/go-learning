package main

import (
	"fmt"
	"retriever/mock"
	"retriever/real"
	// "time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("https://www.baidu.com")
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v:=r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
func main() {
	var r Retriever
	r = mock.Retriever{"this is a fake www.baidu.com"}
	inspect(r)
	// r = &real.Retriever{UserAgent: "Mozilla/5.0",
	// 	TimeOut: time.Minute}
	// inspect(r)
	if mockRetriever, ok := r.(mock.Retriever);ok{
		fmt.Println(mockRetriever.Contents)
	}else{
		fmt.Println("not a mock retriever")
	}
	// fmt.Println(download(r))

}
