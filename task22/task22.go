package main

import (
	"fmt"
	"math/big"
)

func main() {
	var resDiv, x, y float64
	x = 12345678
	y = 23456789
	firstNum := big.NewFloat(x)
	secondNum := big.NewFloat(y)
	resultNum := new(big.Float).SetPrec(100)

	resultNum.Sub(firstNum, secondNum)
	fmt.Println("Вычитание:")
	fmt.Println(resultNum)

	resultNum.Add(firstNum, secondNum)
	fmt.Println("Сложение:")
	fmt.Println(resultNum)

	resultNum.Mul(firstNum, secondNum)
	fmt.Println("Умножение:")
	fmt.Println(resultNum)

	resultNum.Quo(firstNum, secondNum)
	fmt.Println("Деление, используя big:")
	fmt.Println(resultNum)
	fmt.Printf("%.50f\n", resultNum)

	resDiv = x / y
	fmt.Println("Встроенное деление:")
	fmt.Printf("%.50f\n", resDiv)
}
