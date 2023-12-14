package main

import (
	"fmt"
	"log"
)

var (
	errElem = fmt.Errorf("item not found in the list")
)

func BinarySearch(items []int, elem int) (int, error) {
	if len(items) == 1 {
		if items[0] != elem {
			return 0, errElem
		}
		return 0, nil
	}
	if elem >= items[len(items)/2] {
		position, err := BinarySearch(items[len(items)/2:], elem)
		if err != nil {
			return 0, err
		}
		return len(items)/2 + position, nil
	} else {
		position, err := BinarySearch(items[:len(items)/2], elem)
		if err != nil {
			return 0, err
		}
		return position, nil
	}
}

func main() {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	elem := 5
	position, err := BinarySearch(items, elem)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(position)
}
