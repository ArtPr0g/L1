package main

import (
	"fmt"
	"strings"
)

func reverseWords(sentence string) string {
	words := strings.Split(sentence, " ")

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	sentence := "snow dog sun"
	reversedSentence := reverseWords(sentence)
	fmt.Println(reversedSentence)
}
