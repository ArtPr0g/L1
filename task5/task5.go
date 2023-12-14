package main

import (
	"log"
	"sync"
	"time"
)

// отправляет чето-там
func Sender(wg *sync.WaitGroup, ch chan<- int, chClose <-chan bool) {
	defer wg.Done()
	for i := 1; ; i++ {
		select {
		case <-chClose:
			close(ch)
			return
		default:
			ch <- i
			time.Sleep(1 * time.Second)
		}
	}
}

func Receiver(wg *sync.WaitGroup, ch <-chan int) {
	for value := range ch { // Читаем значение из канала
		log.Println("Received:", value)
	}
	wg.Done()
}

func main() {
	N := 10 // Количество секунд работы программы

	ch := make(chan int)       // Создаем канал для передачи значений
	chClose := make(chan bool) // Создаем канал для передачи значений
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go Sender(wg, ch, chClose) // Запускаем отправку значений в канал в отдельной горутине
	wg.Add(1)
	go Receiver(wg, ch) // Запускаем чтение значений из канала в отдельной горутине

	time.Sleep(time.Duration(N) * time.Second) // Ждем N секунд
	chClose <- true
	wg.Wait()
}
