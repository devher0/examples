package main

import (
	"fmt"
)

func main() {

	fmt.Println("Привет! Я автоупаковщик яблок.")
	fmt.Println("Укажите сколько поместится яблок в каждую корзину")
	fmt.Println("-------------------")

	fmt.Println("Карзина №1")
	var LimBasketOne int
	fmt.Scan(&LimBasketOne)

	fmt.Println("Карзина №2")
	var LimBasketTwo int
	fmt.Scan(&LimBasketTwo)

	fmt.Println("Карзина №3")
	var LimBasketThree int
	fmt.Scan(&LimBasketThree)

	basketOne := 0
	basketTwo := 0
	basketThree := 0

	for basketOne <= LimBasketOne && basketTwo <= LimBasketTwo && basketThree <= LimBasketThree {

		if basketOne != LimBasketOne {
			basketOne = basketOne + 1
			fmt.Println("Корзина №1", basketOne)
		} else if basketOne == LimBasketOne {
			fmt.Println("Корзина №1", basketOne)
		} else if basketOne == LimBasketOne && basketTwo == LimBasketTwo && basketThree == LimBasketThree {
			break
		}

		if basketTwo != LimBasketTwo {
			basketTwo = basketTwo + 1
			fmt.Println("Корзина №2", basketTwo)
		} else if basketTwo == LimBasketTwo {
			fmt.Println("Корзина №2", basketTwo)
		} else if basketOne == LimBasketOne && basketTwo == LimBasketTwo && basketThree == LimBasketThree {
			break
		}

		if basketThree != LimBasketThree {
			basketThree = basketThree + 1
			fmt.Println("Корзина №3", basketThree)
		} else if basketThree == LimBasketThree {
			fmt.Println("Корзина №3", basketThree)
		} else if basketOne == LimBasketOne && basketTwo == LimBasketTwo && basketThree == LimBasketThree {
			break
		}

		if basketOne != LimBasketOne || basketTwo != LimBasketTwo || basketThree != LimBasketThree {
			continue
		} else if basketOne == LimBasketOne && basketTwo == LimBasketTwo && basketThree == LimBasketThree {
			break
		}
	}

}
