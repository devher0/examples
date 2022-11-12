package main

import (
	"fmt"
)

func main() {

	fmt.Println("Введите стоимость товара")
	var g int
	fmt.Scan(&g)

	fmt.Println("Введите стоимость доставки")
	var d int
	fmt.Scan(&d)

	fmt.Println("Введите скидку")
	var s int
	fmt.Scan(&s)

	goods := g
	dilivery := d
	discount := s

	price := goods + dilivery - discount

	fmt.Println("Стоимость товара :", goods)
	fmt.Println("Стоимость доставки :", dilivery)
	fmt.Println("Скидка :", discount)
	fmt.Println("--------")
	fmt.Println("К оплате :", price)
}
