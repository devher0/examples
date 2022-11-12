package main

import (
	"fmt"
)

func main() {

	fmt.Println("Давайте проверим ваш билетик !")
	fmt.Println("---------------------------------")

	fmt.Println("Введите четырёхзначный номер вашего билета")
	var ticket int
	fmt.Scan(&ticket)

	firstNumber := ticket / 1000
	secondNumber := (ticket / 100) % 10
	thirdNumber := (ticket % 100) / 10
	fourthNumber := ticket % 10

	partOne := firstNumber + secondNumber
	PartTwo := thirdNumber + fourthNumber

	if firstNumber == fourthNumber && secondNumber == thirdNumber {
		fmt.Println("Данный билетик зеркальный")
	} else if partOne == PartTwo {
		fmt.Println("Данный билетик счастливый")
	} else {
		fmt.Println("Данный билетик обычный")
	}

}
