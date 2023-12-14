package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}
	wg := &sync.WaitGroup{}
	mu := &sync.RWMutex{}
	sum := new(int)
	for _, number := range numbers {
		wg.Add(1)
		go Squaring(wg, mu, number, sum)
	}
	wg.Wait()
	log.Println(*sum)
}

func Squaring(wg *sync.WaitGroup, mu *sync.RWMutex, number int, sum *int) {
	defer wg.Done()
	square := number * number
	mu.RLock()
	newSum := *sum + square
	mu.RUnlock()
	mu.Lock()
	*sum = newSum
	fmt.Println(*sum)
	mu.Unlock()
}
