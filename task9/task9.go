package main

import (
	"log"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	// Создаем каналы для передачи данных
	input := make(chan int)
	output := make(chan int)

	wg.Add(1)
	// Горутина для чтения чисел из массива и отправки в input канал
	go Generate(wg, input)

	wg.Add(1)
	// Горутина для умножения чисел и отправки результата в output канал
	go MultiplicationTwo(wg, input, output)

	wg.Add(1)
	// Горутина для вывода результатов в stdout
	go Display(wg, output)

	wg.Wait()
}

func Generate(wg *sync.WaitGroup, input chan int) {
	defer wg.Done()
	numbers := [5]int{1, 2, 3, 4, 5} // Пример массива чисел
	for _, num := range numbers {
		input <- num
	}
	close(input)
}

func MultiplicationTwo(wg *sync.WaitGroup, input, output chan int) {
	defer wg.Done()
	for num := range input {
		result := num * 2
		output <- result
	}
	close(output)
}

func Display(wg *sync.WaitGroup, input chan int) {
	defer wg.Done()
	for res := range input {
		log.Println(res)
	}
}
