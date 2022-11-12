package main

import (
	"fmt"
)

func main() {

	var circl int = 4
	var characterName string = "Шумахер"
	var speed int = 358
	var engine int = +254
	var wheels int = +93
	var steeringWheel int = +49
	var wind int = -21
	var rain int = -17

	fmt.Println("======================")
	fmt.Println("Супер гонки. Круг", circl)
	fmt.Println("======================")
	fmt.Print(characterName, " (", speed, ")\n")
	fmt.Println("======================")
	fmt.Println("Водитель:", characterName)
	fmt.Println("Скорость", speed)
	fmt.Println("-----------------")
	fmt.Println("Оснащение")
	fmt.Print("Двигатель: +", engine, "\n")
	fmt.Print("Колёса: +", wheels, "\n")
	fmt.Print("Руль: +", steeringWheel, "\n")
	fmt.Println("-----------------")
	fmt.Println("Действия плохой дороги")
	fmt.Println("Ветер:", wind)
	fmt.Println("Дождь:", rain)
}
