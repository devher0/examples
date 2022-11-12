package main

import (
	"fmt"
)

func main() {
	fmt.Println("Введите два произвольных числа")

	var numOne int
	var numTwo int

	fmt.Scan(&numOne)
	fmt.Scan(&numTwo)

	//sum := numOne + numTwo

	for i := 0; i == numTwo; i++ {

		len := numOne + i
		fmt.Println(len)

	}
}
