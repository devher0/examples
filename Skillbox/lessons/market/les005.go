package main

import (
	"fmt"
)

func main() {

	goods := 6400
	dilivery := 350
	discount := 700

	price := goods + dilivery - discount

	fmt.Println("Стоимость товара :", goods)
	fmt.Println("Стоимость доставки :", dilivery)
	fmt.Println("Скидка :", discount)
	fmt.Println("--------")
	fmt.Println("К оплате :", price)
}
