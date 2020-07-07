package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func apply(op func(int, int)int, a, b int) int{
	p:=reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()          // 获得函数的名字
	fmt.Printf("calling function %s with arguments "+
				"(%d, %d)\n", opName, a, b)
	return op(a, b)
}
func pow(a, b int)int{
	return int(math.Pow(float64(a), float64(b)))
}
func main(){
	fmt.Println(apply(pow, 3, 4))
	fmt.Println(apply(
		func(a, b int)int{
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4,
	))
}