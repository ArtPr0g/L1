package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}
	wg := &sync.WaitGroup{}

	for _, number := range numbers {
		wg.Add(1)
		go Squaring(wg, number)
	}
	wg.Wait()
}

func Squaring(wg *sync.WaitGroup, number int) {
	defer wg.Done()
	square := number * number
	fmt.Println(square)
}
