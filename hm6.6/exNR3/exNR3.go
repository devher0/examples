package main

import (
	"fmt"
)

func main() {

	fmt.Println("Введитн цену товара")
	var product float32
	fmt.Scan(&product)

	fmt.Println("Введите процент скидки, процент не должен привышать 30!")
	var discount float32
	fmt.Scan(&discount)

	for discount > 30 {
		fmt.Println("Введите процент скидки, процент не должен привышать 30!")
		var discount float32
		fmt.Scan(&discount)

		if discount <= 30 {

			discSum := product * (discount / 100.0)

			if discSum > 2000 {
				fmt.Println("Скидка превышает норму в 2000 руб.")
				break
			} else if discSum <= 2000 {
				fmt.Println("Сумма скидки :", discSum, "руб.")
				break
			}

		}
	}

	if discount <= 30 {

		discSum := product * (discount / 100.0)

		if discSum > 2000 {
			fmt.Println("Скидка превышает норму в 2000 руб.")
		} else if discSum <= 2000 {
			fmt.Println("Сумма скидки :", discSum, "руб.")
		}

	}

}
