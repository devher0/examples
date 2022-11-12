package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Введите число основании степени")
	var numberOne float64
	fmt.Scan(&numberOne)

	fmt.Println("Введите число показателя степени")
	var numberTwo float64
	fmt.Scan(&numberTwo)

	result := math.Pow(numberOne, numberTwo)

	fmt.Println(numberOne, "в степени", numberTwo, "равно", result)
}
