package utils

import "fmt"

type A interface {
	f()
}
type B struct {
}

func (a *B) Foo() {
	fmt.Println("B实现了A接口")
}
