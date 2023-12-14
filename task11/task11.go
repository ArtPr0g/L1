package main

import (
	"log"
)

func Intersection(set1, set2 []int) []int {
	setMap := make(map[int]bool, len(set1))
	intersection := make([]int, 0)

	for _, num := range set1 {
		setMap[num] = true
	}

	for _, num := range set2 {
		if setMap[num] {
			intersection = append(intersection, num)
		}
	}
	return intersection
}

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{4, 5, 6, 7, 8}

	result := Intersection(set1, set2)
	log.Println(result)
}
