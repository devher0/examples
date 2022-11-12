package main

import (
	"fmt"
)

func main() {

	fmt.Println("Банкомат")
	fmt.Println("----------------------------")
	fmt.Println("Банкомат выдаёт купюры с номеналом в 100 руб.")
	fmt.Println("Максимальная сумма вывода 100 тыс.руб.")
	fmt.Println("Введите сумму которую желаете снять")
	var money int
	fmt.Scan(&money)

	if money > 0 && money%100 == 0 && money <= 100000 {
		fmt.Println("Операция выполнено успешно!")
		fmt.Println("Вы сняли", money, " руб.")
	} else if money%100 != 0 {
		fmt.Println("Ошибка!")
		fmt.Println("Вы ввели", money, "руб.")
		fmt.Println("Данный банкомат выдаёт купюры с номеналом в 100 руб.!")
	} else if money > 100000 {
		fmt.Println("Ошибка!")
		fmt.Println("Вы ввели", money, "руб.")
		fmt.Println("Данный банкомат выдаёт максимальную сумму в 100000 руб")
	} else {
		fmt.Println("Error not found")
	} /*else if money < 0 {
		fmt.Println("Ошибка!")
		fmt.Println("Вы ввели", money)
		fmt.Println("Вы не можете ввести отрицательное число!")
	}*/

}
