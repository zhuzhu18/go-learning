package main

import (
	"fmt"
	"sync"
)

func doWork(id int, c chan int, wg *sync.WaitGroup) {
	for n := range c {
		fmt.Printf("worker %d received %c\n", id, n)
		wg.Done()
	}
}

type worker struct {
	in   chan int
	wg *sync.WaitGroup
}

func createWorker(id int, wg *sync.WaitGroup) worker { // 该channel只能发送数据
	w := worker{
		in:   make(chan int),
		wg: wg,
	}
	go doWork(id, w.in, wg)
	return w
}
func chanDemo() {
	var wg sync.WaitGroup
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg) // 创建10个worker放进一个workers数组中
	}
	wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i // 给数组中的每个worker发送数据
	}

	for j, worker := range workers {
		worker.in <- 'A' + j
	}
	wg.Wait()
}

func main() {
	chanDemo()
}
