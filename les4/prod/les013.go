package main

import (
	"fmt"
)

func main() {

	fmt.Println("Введите стоимость первого товара")
	var p1 int
	fmt.Scan(&p1)

	fmt.Println("Введите стоимость второго товара")
	var p2 int
	fmt.Scan(&p2)

	fmt.Println("Введите стоимость третьего товара")
	var p3 int
	fmt.Scan(&p3)

	fmt.Println("Введите стоимость доставки")
	var d int
	fmt.Scan(&d)

	fmt.Println("Введите скидку")
	var s int
	fmt.Scan(&s)

	basket := p1 + p2 + p3
	dilivery := d
	discount := s

	if basket >= 10000 {
		fmt.Println("--------")
		fmt.Println("Ваша общая стоимость товара состовляет больше 10 000 руб.")
		fmt.Println("По этому вы получаете скидку на доставку")
		superDiscount := dilivery * 10 / 100
		price := basket + dilivery - discount - superDiscount

		fmt.Println("--------")
		fmt.Println("Общая стоимость товара :", basket)
		fmt.Println("Стоимость доставки :", dilivery)
		fmt.Println("Скидка :", discount)
		fmt.Println("--------")
		fmt.Println("К оплате :", price)
	} else {
		price := basket + dilivery - discount

		fmt.Println("--------")
		fmt.Println("Общая стоимость товара :", basket)
		fmt.Println("Стоимость доставки :", dilivery)
		fmt.Println("Скидка :", discount)
		fmt.Println("--------")
		fmt.Println("К оплате :", price)
	}

}
