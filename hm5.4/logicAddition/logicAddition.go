package main

import (
	"fmt"
)

func main() {

	fmt.Println("Проверка на положительное число")
	fmt.Println("Введите первое число")
	var firstNumber int
	fmt.Scan(&firstNumber)
	fmt.Println("Введите второе число")
	var secondNumber int
	fmt.Scan(&secondNumber)
	fmt.Println("Введите третье число")
	var thirdNumber int
	fmt.Scan(&thirdNumber)

	if firstNumber > 0 || secondNumber > 0 || thirdNumber > 0 {
		fmt.Println("Среди введённых чисел есть положительное число")
	} else if firstNumber < 0 && secondNumber < 0 && thirdNumber < 0 {
		fmt.Println("Среди введёных чисел нет положительного числа ")
	} else {
		fmt.Println("Ошибка !")
		fmt.Println("Перезапустите программу!")
	}

	/*  	Второй вариант решения задачи
	var a, b, c int

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	positiveNumber := true
	if a < 0 && b < 0 && c < 0 {
		positiveNumber = false
	}
	if positiveNumber {
		fmt.Println("Среди введенных чисел есть положительное число")
	} else {
		fmt.Println("Среди введенных числе нет положительного числа")
	}
	*/
}
