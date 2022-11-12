package main

import (
	"fmt"
)

func main() {

	fmt.Println("Привет ,вас приветствует умное меню.")
	fmt.Println("Введите день недели.")
	fmt.Println("(понедельник) - 1")
	fmt.Println("(вторник) - 2")
	fmt.Println("(среда) - 3")
	fmt.Println("(четверг) - 4")
	fmt.Println("(пятница) - 5")
	fmt.Println("(суббота) - 6")
	fmt.Println("(воскрессенье) - 7")

	fmt.Println("-------------------------")

	var day int
	fmt.Scan(&day)

	if day == 1 {
		fmt.Println("В этот понедельник в нашем Меню:")
	} else if day == 2 {
		fmt.Println("В этот вторник в нашем Меню:")
	} else if day == 3 {
		fmt.Println("В эту среду в нашем Меню:")
	} else if day == 4 {
		fmt.Println("В этот четверг в нашем Меню:")
	} else if day == 5 {
		fmt.Println("В эту пятницу в нашем Меню:")
	} else if day == 6 {
		fmt.Println("В эту субботу в нашем Меню:")
	} else if day == 7 {
		fmt.Println("В это восскрессенье в нашем Меню:")
	} else {
		fmt.Println("К сожелению вы ввели не коректный день недели")
		fmt.Println("Вам дотупно стандартное меню")
	}

	fmt.Println("горячее...")
	fmt.Println("закуски...")
	fmt.Println("напитки...")

	fmt.Println("Меню дня")

	if day == 1 {
		fmt.Println("Фаржированные криветки криветками на гриле")
	} else if day == 2 {
		fmt.Println("Хачапури по аджарски")
	} else if day == 3 {
		fmt.Println("Мамалыга со шкварками")
	} else if day == 4 {
		fmt.Println("Паста с песто от шеф повора Татьяны")
	} else if day == 5 {
		fmt.Println("Жаренные пельмени с чесночным соусом и соевым соусом")
	} else if day == 6 {
		fmt.Println("Шашлык по грузински")
	} else if day == 7 {
		fmt.Println("Зама")
	} else {
		fmt.Println("Чтоб узнать меню дня перезапустите программу ")
		fmt.Println("И введите корректный день недели")
	}

}
