package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func get(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	} else {
		defer func() { resp.Body.Close() }()
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("%s", contents)
		}
	}
}

func put(url string) {
	request, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		panic(err)
	} else {
		defer func() { request.Body.Close() }()
		r, err := http.DefaultClient.Do(request)
		if err != nil{panic(err)}
		contents, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s", contents)
	}
}
func main() {
	// get("http://httpbin.org")
	put("http://httpbin.org/put")
}
