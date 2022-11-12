package main

import (
	"fmt"
)

func main() {

	fmt.Println("Данная программа проверяет на четность введеное число")
	fmt.Println("Введите целое число")
	var number int
	fmt.Scan(&number)

	remainder := number % 2

	if remainder == 0 {
		fmt.Println("Число", number, "является четным числом")

	} else {
		fmt.Println("Число", number, "является не четным числом")
	}
}
