package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, syscall.SIGINT)
	// Создаем канал для передачи данных
	dataChannel := make(chan string)
	// Запускаем горутину для записи данных в канал
	go SendData(dataChannel)

	// Получаем количество воркеров от пользователя
	var numWorkers int
	fmt.Print("Введите количество воркеров: ")
	_, err := fmt.Scan(&numWorkers)
	if err != nil {
		log.Println("couldn't count the number of workers: ", err)
	}

	// Создаем WaitGroup для ожидания завершения всех воркеров
	wg := &sync.WaitGroup{}

	// Запускаем указанное количество воркеров
	for id := 0; id < numWorkers; id++ {
		wg.Add(1)
		go Worker(wg, dataChannel, id)
	}
	// Ожидаем нажатие Ctrl+C для завершения программы
	<-gracefulShutdown
	// Закрываем канал и ждем завершения всех воркеров
	close(dataChannel)
	wg.Wait()
}

// Функция для записи данных в канал
func SendData(ch chan<- string) {
	for i := 1; ; i++ {
		time.Sleep(1 * time.Second)
		// Записываем данные в канал
		ch <- fmt.Sprintf("Данные %d", i)
	}
}

// Функция воркера, которая читает данные из канала и выводит их в stdout
func Worker(wg *sync.WaitGroup, ch <-chan string, id int) {
	defer wg.Done()
	for data := range ch {
		fmt.Printf("Воркер %d: %s\n", id, data)
	}
}
