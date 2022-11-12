package main

import (
	"fmt"
)

func main() {

	fmt.Println("Три числа")
	fmt.Println("Введите первое число")
	var firstNumber int
	fmt.Scan(&firstNumber)
	fmt.Println("Введите второе число")
	var secondNumber int
	fmt.Scan(&secondNumber)
	fmt.Println("Введите третье число")
	var thirdNumber int
	fmt.Scan(&thirdNumber)

	if firstNumber > 5 || secondNumber > 5 || thirdNumber > 5 {
		fmt.Println("Среди введённых чисел есть число больше 5")
	} else if firstNumber < 5 && secondNumber < 5 && thirdNumber < 5 {
		fmt.Println("Среди введёных чисел нет числа больше 5")
	} else {
		fmt.Println("Ошибка !")
		fmt.Println("Перезапустите программу!")
	}
}
