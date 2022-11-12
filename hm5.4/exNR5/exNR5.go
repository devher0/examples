package main

import (
	"fmt"
)

func main() {

	fmt.Println("Давайте определим выгодные процентные ставки")

	fmt.Println("Введите три процентные ставки")

	var firstRate float64
	fmt.Scan(&firstRate)
	var secondRate float64
	fmt.Scan(&secondRate)
	var thirdRate float64
	fmt.Scan(&thirdRate)

	if firstRate < secondRate && firstRate < thirdRate {
		fmt.Println("Выгодная ставка", secondRate, "%", "и", thirdRate, "%")
	} else if secondRate < thirdRate {
		fmt.Println("Выгодная ставка", firstRate, "%", "и", thirdRate, "%")
	} else {
		fmt.Println("Выгодная ставка", firstRate, "%", "и", secondRate, "%")
	}

}
