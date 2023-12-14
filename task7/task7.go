package main

import (
	"log"
	"math/rand"
	"sync"
)

func main() {
	data := make(map[int]int)
	wg := &sync.WaitGroup{}
	mu := &sync.RWMutex{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Record(wg, mu, data, i)
	}
	wg.Wait()
	log.Println(data)
}

func Record(wg *sync.WaitGroup, mu *sync.RWMutex, data map[int]int, index int) {
	defer wg.Done()
	mu.Lock()
	data[rand.Int()] = index
	mu.Unlock()
}
