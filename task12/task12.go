package main

import "log"

func GetPlenty(sequence []string) []string {
	set := make(map[string]bool)
	for _, str := range sequence {
		set[str] = true
	}
	var setResult []string
	for key := range set {
		setResult = append(setResult, key)
	}
	return setResult
}

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}
	setResult := GetPlenty(sequence)
	log.Println(setResult)
}
