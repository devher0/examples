package main

import (
	"fmt"
)

func main() {

	fmt.Println("Давайте узнаем координатную четверть точки")

	fmt.Println("Введите координату X")
	var X int
	fmt.Scan(&X)

	fmt.Println("Введите координату Y")
	var Y int
	fmt.Scan(&Y)

	if X > 0 && Y > 0 {
		fmt.Println("Данная точка находится в первой четверти координат")
	} else if X < 0 && Y > 0 {
		fmt.Println("Данная точка находится во второй четверти координат")
	} else if X < 0 && Y < 0 {
		fmt.Println("Данная точка находится в третьей четверти координат")
	} else {
		fmt.Println("Данная точка находится в четвертой четверти координат")
	}

}
