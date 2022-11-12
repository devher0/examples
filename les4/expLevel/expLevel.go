package main

import (
	"fmt"
)

func main() {

	fmt.Println("****Калькулятор уровня персонажа****")
	fmt.Println("Введите количество полученных очков опыта")
	var exp int
	fmt.Scan(&exp)

	if exp >= 0 && exp < 1000 {
		fmt.Println("Вашь персонаж первого уровня")
	} else if exp >= 1000 && exp < 2500 {
		fmt.Println("Вашь персонаж второго уровня")
	} else if exp >= 2500 && exp < 5000 {
		fmt.Println("Вашь персонаж третьего уровня")
	} else if exp >= 5000 {
		fmt.Println("Вашь персонаж четвертого уровня")
	} else {
		fmt.Println("Введите положительное число")
	}

}
