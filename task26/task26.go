package main

import (
	"fmt"
	"strings"
)

func IsUnique(str string) bool {
	seen := make(map[rune]bool)

	str = strings.ToLower(str)

	for _, char := range str {
		if seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}

func main() {
	s1 := "abcd"
	fmt.Println(s1, "-", IsUnique(s1)) // Результат: true

	s2 := "abCdefAaf"
	fmt.Println(s2, "-", IsUnique(s2)) // Результат: false

	s3 := "aabcd"
	fmt.Println(s3, "-", IsUnique(s3)) // Результат: false
}
