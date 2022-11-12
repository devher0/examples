package main

import (
	"fmt"
)

func main() {
	fmt.Println("высший показатель IQ")
	fmt.Println("Введите IQ каждого экипажа команды последовательно.")

	fmt.Println("Введите IQ и имя членов команды.")
	fmt.Println("Введите имя первого члена экипажа.")
	var firstMemberName string
	fmt.Scan(&firstMemberName)
	fmt.Println("Введите IQ первого члена экипажа.")
	var firstMemberIQ int
	fmt.Scan(&firstMemberIQ)

	fmt.Println("Введите имя второго члена экипажа.")
	var secondMemberName string
	fmt.Scan(&secondMemberName)
	fmt.Println("Введите IQ второго члена экипажа.")
	var secondMemberIQ int
	fmt.Scan(&secondMemberIQ)

	fmt.Println("Введите имя третьего члена экипажа.")
	var thridMemberName string
	fmt.Scan(&thridMemberName)
	fmt.Println("Введите IQ третьего члена экипажа.")
	var thridMemberIQ int
	fmt.Scan(&thridMemberIQ)

	if secondMemberIQ < firstMemberIQ && thridMemberIQ < firstMemberIQ {
		fmt.Println("Поздравляем вас", firstMemberName)
		fmt.Println("Ваш IQ", firstMemberIQ, " самый высокий!!!")
		fmt.Println("Вы становитесь лидером группы!")
	} else if firstMemberIQ < secondMemberIQ && thridMemberIQ < secondMemberIQ {
		fmt.Println("Поздравляем вас", secondMemberName)
		fmt.Println("Ваш IQ", secondMemberIQ, " самый высокий!!!")
		fmt.Println("Вы становитесь лидером группы!")
	} else if firstMemberIQ < thridMemberIQ && secondMemberIQ < thridMemberIQ {
		fmt.Println("Поздравляем вас", thridMemberName)
		fmt.Println("Ваш IQ", thridMemberIQ, " самый высокий!!!")
		fmt.Println("Вы становитесь лидером группы!")
	} else {
		fmt.Println("Ошибка !")
		fmt.Println("Вы ввели не верные данные.")
		fmt.Println("Перезапустите программу!")
	}

}
