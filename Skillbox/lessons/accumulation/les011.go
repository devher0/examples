package main

import (
	"fmt"
)

func main() {

	fmt.Println("Давай копить!")
	fmt.Println("А на что ? Введите цель!")

	var target string
	fmt.Scan(&target)

	total := 0
	var money int

	fmt.Println("Сумма 1")
	fmt.Scan(&money)
	total += money

	fmt.Println("Сумма 2")
	fmt.Scan(&money)
	total += money

	fmt.Println("Сумма 3")
	fmt.Scan(&money)
	total += money

	fmt.Println("Сумма 4")
	fmt.Scan(&money)
	total += money

	fmt.Println("Сумма 5")
	fmt.Scan(&money)
	total += money

	fmt.Println("Ты накопил на", target, ". Получилась накопить", total, "руб.")

}
