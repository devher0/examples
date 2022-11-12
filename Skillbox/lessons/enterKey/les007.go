package main

import (
	"fmt"
)

func main() {

	fmt.Println("Введите число которое хотите возвести в квадрат")

	var i int
	fmt.Scan(&i)

	square := i * i

	fmt.Println("Вы ввели", i)
	fmt.Println("Квадрат введённого числа равен:", square)

}
