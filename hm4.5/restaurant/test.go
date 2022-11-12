package main

import (
	"fmt"
)

func main() {

	fmt.Println("Проверте чек на скидку !")
	fmt.Println("Введите день недели")
	var day int
	fmt.Scan(&day)
	fmt.Println("Введите количество людей")
	var man int
	fmt.Scan(&man)
	fmt.Println("Введите сумму оплаты")
	var check int
	fmt.Scan(&check)

	if day == 1 && man <= 5 {
		discount := (check * 10) / 100
		fmt.Println("Скидка ", discount, " лей, потому-что понедельник день тяжелый.")
		fmt.Println("Надбавки на обслуживание нет")
		totalToPay := check - discount
		fmt.Println("Сумма к оплате: ", totalToPay, "лей.")
	} else if day == 1 && man > 5 {
		discount := (check * 10) / 100
		fmt.Println("Скидка ", discount, " лей, потому-что понедельник день тяжелый.")
		increase := (check * 10) / 100
		fmt.Println("Надбавка на обслуживание ", increase, " лей.")
		totalToPay := check - discount + increase
		fmt.Println("Сумма к оплате: ", totalToPay, "лей.")
	} else if day == 5 && check > 10000 && man <= 5 {
		discount := (check * 5) / 100
		fmt.Println("'Скидка по пятницам'", discount, " лей.")
		fmt.Println("Надбавки на обслуживание нет")
		totalToPay := check - discount
		fmt.Println("Сумма к оплате: ", totalToPay, "лей.")
	} else if day == 5 && check > 10000 && man > 5 {
		discount := (check * 5) / 100
		fmt.Println("'Скидка по пятницам'", discount, " лей.")
		increase := (check * 10) / 100
		fmt.Println("Надбавка на обслуживание ", increase, " лей.")
		totalToPay := check - discount + increase
		fmt.Println("Сумма к оплате: ", totalToPay, "лей.")
	} else if day == 2 && man > 5 || day == 3 && man > 5 || day == 4 && man > 5 || day == 6 && man > 5 || day == 7 && man > 5 {
		fmt.Println("Скидки нет!")
		increase := (check * 10) / 100
		fmt.Println("Надбавка на обслуживание ", increase, " лей.")
		totalToPay := check + increase
		fmt.Println("Сумма к оплате: ", totalToPay, "лей.")
	} else {
		fmt.Println("Скидки нет!")
		fmt.Println("Сумма к оплате: ", check, "лей.")
	}

}
