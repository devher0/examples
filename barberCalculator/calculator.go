package main

import (
	"fmt"
)

func main() {

	fmt.Println("******Барбершоп-Калькулятор******")

	fmt.Println("Введите число мужчин в городе.")
	var mansCount int
	fmt.Scan(&mansCount)

	fmt.Println("Сколько барберов уже удалось нанять ?")
	var barbersCount int
	fmt.Scan(&barbersCount)

	// сколько человек может постричь барбер за одну смену ?
	mansPerBarber := 8 // один человек в час , смена 8 часов

	// сколько человек может постричь барбер за месяц ?
	mansPerBarberPerMonth := mansPerBarber * 30

	// сколько нужно барберов чтоб постричь mansPerDay человек
	requairedBarbesCount := mansCount / mansPerBarberPerMonth
	if requairedBarbesCount*mansPerBarberPerMonth < mansCount {
		requairedBarbesCount++
	}

	fmt.Println("Необходимое число барберов", requairedBarbesCount)
	fmt.Println(requairedBarbesCount, " барбера могут постричь", mansPerBarberPerMonth*requairedBarbesCount, " мужчин за месяц.")

	//сравниваем с имеющимся количеством барберов
	if requairedBarbesCount > barbersCount {
		fmt.Println("Нужно больше барберов !")
	}
	if requairedBarbesCount <= barbersCount {
		fmt.Println("Барберов достатачно !")
	}

}
