package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"sync"
)

func CloseChannel(wg *sync.WaitGroup, ch chan string) {
	for v := range ch {
		log.Println(v)
	}
	log.Println("goroutine is ending")
	wg.Done()
}

func ControlChannel(wg *sync.WaitGroup, ch chan string, done chan bool) {
	for {
		select {
		case <-done:
			log.Println("goroutine is ending")
			wg.Done()
			return
		default:
			//делаем что-то с каналом данных ch
		}
	}
}

func UseContext(ctx context.Context, wg *sync.WaitGroup, ch chan string) {
	for {
		select {
		case <-ctx.Done():
			log.Println("goroutine is ending")
			wg.Done()
			return
		default:
			//делаем что-то с каналом данных ch и context
		}
	}
}

func UseGoExit(wg *sync.WaitGroup, ch chan string) {
	for {
		//что-то делаю
		wg.Done()
		runtime.Goexit()
	}
}

func main() {
	var numWay int
	fmt.Print("Введите номер способа: ")
	_, err := fmt.Scan(&numWay)
	if err != nil {
		log.Println("couldn't count the number of way: ", err)
	}
	wg := &sync.WaitGroup{}
	ch := make(chan string)
	switch numWay {
	case 1:
		wg.Add(1)
		go CloseChannel(wg, ch)
		close(ch)
	case 2:
		wg.Add(1)
		done := make(chan bool)
		go ControlChannel(wg, ch, done)
		close(done)
	case 3:
		ctx, cancel := context.WithCancel(context.Background())
		wg.Add(1)
		go UseContext(ctx, wg, ch)
		cancel()
	case 4:
		wg.Add(1)
		go UseGoExit(wg, ch)
	default:
		fmt.Println("Такой способ не найден")
	}
	wg.Wait()

}
