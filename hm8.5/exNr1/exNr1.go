package main

import (
	"fmt"
)

func main() {

	fmt.Println("Времена года")
	fmt.Println("Введите месяц")

	var month string
	fmt.Scan(&month)

	switch month {
	case "декабрь", "январь", "февраль":
		fmt.Println("Время года - Зима")
	case "март", "апрель", "май":
		fmt.Println("Время года - Весна")
	case "июнь", "июль", "август":
		fmt.Println("Время года - Лето")
	case "сентябрь", "октябрь", "ноябрь":
		fmt.Println("Время года - Осень")
	}
}
