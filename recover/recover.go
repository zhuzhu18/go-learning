package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic(fmt.Sprintf("I don't known what to do: %v", r))
		}
	}()
	// panic(errors.New("this is an error"))
	b := 0
	a := 5 / b
	fmt.Println(a)
	// panic(456)
}
func main() {
	tryRecover()
}
