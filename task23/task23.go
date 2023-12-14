package main

import "log"

// порядок меняется
func Remove1(data []int, index int) []int {
	data[index] = data[len(data)-1]
	data = data[:len(data)-1]
	return data
}

// порядок остается
func Remove2(data []int, index int) []int {
	return append(data[:index], data[index+1:]...)
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7}
	index := 2
	data = Remove1(data, index)
	log.Println(data)
	data = Remove2(data, index)
	log.Println(data)
}
