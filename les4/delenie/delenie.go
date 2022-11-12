package main

import (
	"fmt"
)

func main() {

	fmt.Println("Это программа, которая проверяет, делится ли одно число на другое без остатка.")

	fmt.Println("Введите любое целое число")
	var firstNumber int
	fmt.Scan(&firstNumber)

	fmt.Println("Введите второе любое целое число")
	var secondNumber int
	fmt.Scan(&secondNumber)

	if firstNumber%secondNumber == 0 {
		fmt.Println(firstNumber, " делится без остатка на", secondNumber)
	} else {
		fmt.Println(firstNumber, " не делится без остатка на", secondNumber)
	}

	fmt.Println(firstNumber % secondNumber)
}
