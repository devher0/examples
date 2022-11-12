package main

import (
	"fmt"
)

func main() {

	fmt.Println("Проверка на совпадение")
	fmt.Println("Введите три числа")

	var a, b, c int

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	if a == b || a == c || b == c {
		fmt.Println("Есть совпадение среди введенных чисел !")
	} else {
		fmt.Println("Совпадений нет среди введенных чисел")
	}

}
