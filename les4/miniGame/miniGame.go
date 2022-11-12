package main

import (
	"fmt"
)

func main() {

	fmt.Println("****Проверка вашей арифметики****")
	fmt.Println("Ваша задача посчитать сумму двух чисел в уме!")
	fmt.Println("А после чего я проверю вашь ответ!")
	fmt.Println("Удачи !!!")
	fmt.Println("------------------------")

	fmt.Println("Введите число")
	var firstNumber int
	fmt.Scan(&firstNumber)

	fmt.Println("Введиет второе число")
	var secondNumber int
	fmt.Scan(&secondNumber)

	fmt.Println("Теперь посчитайте в уме.")
	fmt.Println("Сколько будет", firstNumber, "+", secondNumber, "?")

	fmt.Println("Введите верный ответ.")
	var answer int
	fmt.Scan(&answer)

	if firstNumber+secondNumber == answer {
		fmt.Println("Верно !")
		fmt.Println(firstNumber, "+", secondNumber, "будет ровно", answer)
	} else {
		fmt.Println("Неверно... ")
		fmt.Println(firstNumber, "+", secondNumber, "будет ровно", firstNumber+secondNumber)
	}
}
