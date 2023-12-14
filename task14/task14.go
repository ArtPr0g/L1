package main

import (
	"fmt"
	"reflect"
)

func main() {
	var variable interface{}

	variable = 42
	CheckType(variable)

	variable = "Hello, World!"
	CheckType(variable)

	variable = true
	CheckType(variable)

	variable = make(chan int)
	CheckType(variable)
}

func CheckType(variable interface{}) {
	fmt.Printf("Значение: %v\nТип: %s\n\n", variable, reflect.TypeOf(variable))
}
