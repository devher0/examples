package main

import (
	"fmt"
)

func main() {

	fmt.Println("Проверте чек на скидку !")
	fmt.Println("Введите день недели выдачи чека")
	var day int
	fmt.Scan(&day)
	fmt.Println("Введите сумму оплаты")
	var check int
	fmt.Scan(&check)

	if day == 1 || day == 2 || day == 3 || day == 4 || day == 5 {
		discount := (check * 10) / 100
		totalToPay := check - discount
		fmt.Println("Скидка -", discount, " лей.")
		fmt.Println("К оплате -", totalToPay, " лей.")
	} else if day == 6 && check >= 10000 || day == 7 && check >= 10000 {
		discount := (check * 10) / 100
		totalToPay := check - discount
		fmt.Println("Скидка -", discount, " лей.")
		fmt.Println("К оплате -", totalToPay, " лей.")
	} else {
		fmt.Println("Скидки нет!")
		fmt.Println("К оплате -", check, " лей.")
	}

}
