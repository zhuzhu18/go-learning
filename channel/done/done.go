package main

import (
	"fmt"
)

func doWork(id int, c chan int, done chan bool) {
	// fmt.Printf("worker %d received %c\n", id, <-c) // 如果发送方发送完后close，读取完前四个字符后还继续读(0值)直到一毫秒结束
	// for{
	// 	n, ok := <-c
	// 	if !ok { // 如果发送方发送完后close，ok将为false，表明channel中没有数据了
	// 		break
	// 	}
	// 	fmt.Printf("worker %d received %c\n", id, n)
	// 	done<-true
	// }
	for n := range c {
		fmt.Printf("worker %d received %c\n", id, n)
		done<-true
	}
}

type worker struct {
	in   chan int
	done chan bool
}

func createWorker(id int) worker { // 该channel只能发送数据
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWork(id, w.in, w.done)
	return w
}
func chanDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i) // 创建10个worker放进一个workers数组中
	}
	for i, worker := range workers {
		worker.in <- 'a' + i // 给数组中的每个worker发送数据
	}
	for _, worker := range workers{
		<-worker.done
	}
	for j, worker := range workers {
		worker.in <- 'A' + j
	}
	for _, worker := range workers{
		<-worker.done
	}
}

func main() {
	chanDemo()
}
