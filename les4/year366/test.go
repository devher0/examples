package main

import (
	"fmt"
)

func main() {

	fmt.Println("Данная программа вычесляет високосный год")
	fmt.Println("Введите год")
	var year int
	fmt.Scan(&year)

	if year%4 != 0 || year%100 == 0 && year%400 != 0 {
		fmt.Println("Год ", year, "является не високосным!")
	} else if year%4 == 0 || year%100 == 0 && year%400 == 0 {
		fmt.Println("Год ", year, "является  високосным!")
	}
}
