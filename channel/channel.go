package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for {
		// fmt.Printf("worker %d received %c\n", id, <-c) // 如果发送方发送完后close，读取完前四个字符后还继续读(0值)直到一毫秒结束
		// n, ok := <-c
		// if !ok { // 如果发送方发送完后close，ok将为false，表明channel中没有数据了
		// 	break
		// }
		// fmt.Printf("worker %d received %c\n", id, n)
		for n := range c {
			fmt.Printf("worker %d received %c\n", id, n)
		}
	}
}
func createWorker(id int) chan<- int { // 该channel只能发送数据
	c := make(chan int)
	go worker(id, c)
	return c // channel可以作为返回的参数，并且用<-表明返回的c只能用于接收数据
}
func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i) // 创建10个channel放进一个channels数组中
	}
	for j := 0; j < 10; j++ {
		channels[j] <- 'a' + j // 给数组中的每个channel发送数据
	}
	for j := 0; j < 10; j++ {
		channels[j] <- 'A' + j
		// close(channels[j])
	}

	time.Sleep(time.Millisecond)
}
func bufferedChannel() {
	c := make(chan int, 3) // 创建一个带buffer的channel，缓冲区大小为3
	go worker(0, c)        // 读取channel中的数据
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}
func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c) // 明确告知接收方数据发送完了
	time.Sleep(time.Millisecond)
}
func main() {
	// chanDemo()
	// bufferedChannel()
	channelClose()
}
