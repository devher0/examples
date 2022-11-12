package main

import (
	"fmt"
)

func main() {

	fmt.Println("Данная программа высчитывает сможете ли вы доехать до Рязани ")
	fmt.Println("при условной средней скорости за два часа .")
	fmt.Println("Для этого введите среднюю скорость движения вашего автомобиля.")

	var speed int
	fmt.Scan(&speed)

	road := speed * 2

	if road >= 200 {
		fmt.Println("Вы успеете доехать !")
	} else {
		fmt.Println("Вы не успеете")
		fmt.Println("Вы проедите", road, "километров")
	}

}
