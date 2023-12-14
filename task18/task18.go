package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	value int32
	wg    *sync.WaitGroup
}

func (c *Counter) Increment() {
	atomic.AddInt32(&c.value, 1) // атомарно
	c.wg.Done()
}

func main() {
	counter := Counter{
		wg: &sync.WaitGroup{},
	}
	for i := 0; i < 10; i++ {
		counter.wg.Add(1)
		go counter.Increment()
	}
	counter.wg.Wait()
	fmt.Println("Final counter value:", counter.value)
}
