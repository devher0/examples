package main

import (
	"fmt"
)

func main() {

	i := 0
	a := 0

	for i < 4 {
		fmt.Println("Введите число")
		i = i + 1

		var number int
		fmt.Scan(&number)
		sum := a + number
		a = sum
	}
	fmt.Println(a)
}
