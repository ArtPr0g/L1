package main

import "fmt"

type Human struct {
	// поля структуры Human
}

func (h *Human) Method1() {
	fmt.Println("Method1 called")
}

func (h *Human) Method2() {
	fmt.Println("Method2 called")
}

type Action struct {
	Human // Встраиваем структуру Human в структуру Action
	// Дополнительные поля структуры Action
}

func main() {
	action := Action{}
	action.Method1() // Вызов метода Method1 из структуры Action
	action.Method2() // Вызов метода Method2 из структуры Action
}
