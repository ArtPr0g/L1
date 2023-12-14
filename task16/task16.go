package main

import "log"

func Quicksort(source []int) []int {
	if len(source) < 2 {
		return source
	}
	referenceNumber := source[0]
	result := make([]int, 0, len(source))
	smaller := make([]int, 0)
	larger := make([]int, 0)
	equal := make([]int, 0)
	for _, elem := range source {
		switch {
		case elem < referenceNumber:
			smaller = append(smaller, elem)
		case elem == referenceNumber:
			equal = append(equal, elem)
		case elem > referenceNumber:
			larger = append(larger, elem)
		}
	}
	result = append(result, Quicksort(smaller)...)
	result = append(result, equal...)
	result = append(result, Quicksort(larger)...)
	return result
}

func main() {
	testSlice := []int{100, 5, 2, 56, 123123, 4, 8, 6, 3, 12, 12, 12, 5, 6, 7, 8, 12, 4, 11}
	result := Quicksort(testSlice)
	log.Println(result)
}
