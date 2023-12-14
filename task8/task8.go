package main

import (
	"fmt"
	"log"
)

var (
	errIndex = fmt.Errorf("the bit number cannot be negative")
)

func main() {
	var num int64 = 21 // Заданное число
	index := 3         // Индекс бита

	fmt.Printf("Исходное число: %d\n", num)

	numRes, err := ChangeBit(num, index, 0) // Установка i-го бита в 0
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("Число после установки %d-го бита в 0: %d\n", index, numRes)

	numRes, err = ChangeBit(num, index, 1) // Установка i-го бита в 1
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("Число после установки %d-го бита в 1: %d\n", index, numRes)
}

// Функция для установки i-го бита в числе n в значение bitValue (0 или 1)
func ChangeBit(n int64, index int, bitValue int) (int64, error) {
	if index < 0 {
		return 0, errIndex
	}
	var result int64
	mask := int64(1) << index
	if bitValue == 1 {
		result = mask | n
	} else {
		result = n & ^mask
	}
	return result, nil
}
