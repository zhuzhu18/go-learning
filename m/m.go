package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator()chan int{
	c:=make(chan int)
	go func(){
		i := 0
		for{
			time.Sleep(time.Duration(rand.Intn(1500))*time.Millisecond)
			c<-i
			i++
		}
	}()
	return c
}
func worker(id int, c chan int){
	for n:= range c{
		time.Sleep(2*time.Second)
		fmt.Printf("worker %d received %d\n", id, n)
	}
}
func createWorker(id int)chan int{
	c:=make(chan int)
	go worker(id, c)
	return c
}
func main() {
	var c1, c2 = generator(), generator()
	hasValue := false
	n := 0
	var worker = createWorker(0)
	for{
		var w chan int
		if hasValue{
			w = worker
		}
		select{
		case n=<-c1:         // 两个channel发送数据给n，如果发送过快，则后一次发送的可能会覆盖前一次，需用一个队列存起来
			hasValue = true
		case n=<-c2:
			hasValue = true
		case w<-n:
			hasValue = false
		}
	}
}
