package main

import "fmt"

func main() {
	a := 10
	b := 34

	a = a + b
	b = a - b
	a = a - b

	fmt.Println("После замены:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)
}
