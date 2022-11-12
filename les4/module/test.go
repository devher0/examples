package main

import (
	"fmt"
)

func main() {

	fmt.Println("Данная программа выводит модуль любого целого числа.")
	fmt.Println("Введите любое целое симло.")

	var c int
	fmt.Scan(&c)
	module := c
	module = -module

	fmt.Println("Модуль числа", c, " является", module)

}
