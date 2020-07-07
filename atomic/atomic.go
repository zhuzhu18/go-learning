package main

import (
	"fmt"
	"sync"
	"time"
)
type atomicInt struct{
	value int
	lock sync.Mutex
}
func (a *atomicInt) increment(){
	a.lock.Lock()         // 传统同步机制采用共享内存来通信，所以要加锁来保护，
	                      // 防止数据在一边读取时一边被修改
	defer a.lock.Unlock()
	a.value++
}
func (a *atomicInt) get()int{
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}
func main(){
	var a atomicInt
	a.increment()
	go func(){
		a.increment()
	}()
	time.Sleep(time.Microsecond)
	fmt.Println(a.get())
}