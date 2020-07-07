package main

import (
	"fmt"
	"time"
)

func main() {
	var a [30]int
	for i := 0; i < 30; i++ {
		go func(i int) {
			for {
				a[i]++
			}
		}(i)
	}
	time.Sleep(time.Second * 2)       // 让for里面的10个协程在两秒内有机会执行完
	fmt.Println(a)
}
