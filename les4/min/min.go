package main

import (
	"fmt"
)

func main() {

	fmt.Println("Данная программа вычисляет минимальное число")
	fmt.Println("Введите число")
	var firstNumber int
	fmt.Scan(&firstNumber)
	fmt.Println("Введите второе число")
	var secondNumber int
	fmt.Scan(&secondNumber)

	if firstNumber < secondNumber {
		fmt.Println(firstNumber, " меньше", secondNumber, " на ", secondNumber-firstNumber)
	} else if secondNumber < firstNumber {
		fmt.Println(secondNumber, " меньше", firstNumber, " на ", firstNumber-secondNumber)
	} else {
		fmt.Println("Оба числа равны!")
	}

}
